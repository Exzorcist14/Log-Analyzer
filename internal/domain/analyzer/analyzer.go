package analyzer

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/log"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/report"
	"github.com/montanaflynn/stats"
)

type parser interface {
	Parse(lg string) (*log.Record, error)
}

type statistics struct {
	from              string
	to                string
	field             string
	value             string
	requestsCount     int
	totalResponseSize int
	files             []string
	responseSizes     []float64
	resources         map[string]int
	codes             map[int]int
	clients           map[string]int
	agents            map[string]int
}

type Analyzer struct {
	parser            parser
	stats             statistics
	from              time.Time
	to                time.Time
	field             string
	value             string
	read              int
	isFromSpecified   bool
	isToSpecified     bool
	isFilterSpecified bool
}

func New(ps parser) *Analyzer {
	return &Analyzer{
		parser: ps,
		stats: statistics{
			resources: make(map[string]int),
			codes:     make(map[int]int),
			clients:   make(map[string]int),
			agents:    make(map[string]int),
		},
	}
}

func (a *Analyzer) Analyze(
	from, to time.Time,
	field, value string,
	read int,
	isFromSpecified, isToSpecified, isFilterSpecified bool,
	paths []string, isLocal bool,
) (rep report.Report, err error) {
	a.assignInitialData(from, to, field, value, paths, read, isFromSpecified, isToSpecified, isFilterSpecified)

	if isLocal {
		for _, path := range paths {
			err = a.ProcessLocalLogFile(path)
			if err != nil {
				return rep, fmt.Errorf("can`t process log file: %w", err)
			}
		}
	} else {
		for _, path := range paths {
			err = a.ProcessRemoteLogFile(path)
			if err != nil {
				return rep, fmt.Errorf("can`t process log file: %w", err)
			}
		}
	}

	rep, err = generateReport(&a.stats)
	if err != nil {
		return rep, fmt.Errorf("can`t generate report: %w", err)
	}

	return rep, nil
}

func (a *Analyzer) assignInitialData(
	from, to time.Time,
	field, value string,
	paths []string,
	read int,
	isFromSpecified, isToSpecified, isFilterSpecified bool,
) {
	if isFromSpecified {
		a.stats.from = from.String()
	} else {
		a.stats.from = "-"
	}

	if isToSpecified {
		a.stats.to = to.String()
	} else {
		a.stats.to = "-"
	}

	a.stats.files = append(a.stats.files, paths...)
	a.stats.field = field
	a.stats.value = value

	a.from = from
	a.to = to
	a.field = field
	a.value = value
	a.read = read
	a.isFromSpecified = isFromSpecified
	a.isToSpecified = isToSpecified
	a.isFilterSpecified = isFilterSpecified
}

func (a *Analyzer) ProcessLocalLogFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("can`t open local log file: %w", err)
	}
	defer file.Close()

	err = a.addToStatisticsFromLog(file)
	if err != nil {
		return fmt.Errorf("can`t add log to interim statistics: %w", err)
	}

	return nil
}

func (a *Analyzer) ProcessRemoteLogFile(u string) error {
	parsedURL, err := url.Parse(u)
	if err != nil {
		return fmt.Errorf("can`t parse url: %w", err)
	}

	resp, err := http.Get(parsedURL.String())
	if err != nil {
		return fmt.Errorf("can`t make GET request: %w", err)
	} else if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("can`t make GET request: %w", ErrWrongResponseCode{resp.StatusCode})
	}
	defer resp.Body.Close()

	err = a.addToStatisticsFromLog(resp.Body)
	if err != nil {
		return fmt.Errorf("can`t add to interim statistics: %w", err)
	}

	return nil
}

func (a *Analyzer) addToStatisticsFromLog(lg io.Reader) error {
	scn := bufio.NewScanner(lg)
	linesRead := 0

	for scn.Scan() && linesRead < a.read {
		logRecord, err := a.parser.Parse(scn.Text())
		if err != nil {
			return fmt.Errorf("can`t parse scan result: %w", err)
		}

		isCheckSuccessful, err := a.check(logRecord)
		if err != nil {
			return fmt.Errorf("can't check the lg to satisfy the conditions: %w", err)
		}

		if isCheckSuccessful {
			a.addToStatisticsFromLogRecord(logRecord)

			linesRead++
		}
	}

	if err := scn.Err(); err != nil {
		return fmt.Errorf("can`t scan: %w", err)
	}

	return nil
}

func (a *Analyzer) addToStatisticsFromLogRecord(logRecord *log.Record) {
	a.stats.requestsCount++
	a.stats.resources[logRecord.Request.Resource]++
	a.stats.codes[logRecord.Status]++
	a.stats.clients[logRecord.RemoteAddr]++
	a.stats.agents[logRecord.HTTPUserAgent]++
	a.stats.responseSizes = append(a.stats.responseSizes, float64(logRecord.BodyBytesSent))
	a.stats.totalResponseSize += logRecord.BodyBytesSent
}

func (a *Analyzer) check(record *log.Record) (bool, error) {
	var err error

	isTimeSuccessful := true
	isFilterSuccessful := true

	if a.isFromSpecified || a.isToSpecified {
		isTimeSuccessful = CheckTime(record.TimeLocal, a.from, a.to, a.isFromSpecified, a.isToSpecified)
	}

	if a.isFilterSpecified {
		isFilterSuccessful, err = checkFilter(record, a.field, a.value)
		if err != nil {
			return false, fmt.Errorf("can`t check filter: %w", err)
		}
	}

	return isTimeSuccessful && isFilterSuccessful, nil
}

func CheckTime(current, from, to time.Time, isFromSpecified, isToSpecified bool) bool {
	switch {
	case isFromSpecified && isToSpecified:
		return (current.After(from) || current.Equal(from)) && (current.Before(to) || current.Equal(to))
	case isFromSpecified:
		return current.After(from) || current.Equal(from)
	case isToSpecified:
		return current.Before(to) || current.Equal(to)
	default:
		return true
	}
}

func checkFilter(record *log.Record, field, value string) (bool, error) {
	var current string

	switch field {
	case "remote_add":
		current = record.RemoteAddr
	case "remote_user":
		current = record.RemoteUser
	case "time_local":
		current = record.TimeLocal.String()
	case "method":
		current = record.Request.Method
	case "resource":
		current = record.Request.Resource
	case "protocol":
		current = record.Request.Protocol
	case "status":
		current = strconv.Itoa(record.Status)
	case "body_bytes_sent":
		current = strconv.Itoa(record.BodyBytesSent)
	case "http_referer":
		current = record.HTTPRefer
	case "http_user_agent":
		current = record.HTTPUserAgent
	default:
		return false, ErrUnknownField{field}
	}

	return regexp.MustCompile(value).MatchString(current), nil
}

func generateReport(st *statistics) (report.Report, error) {
	var (
		percentile float64
		err        error
	)

	if len(st.responseSizes) != 0 {
		percentile, err = stats.Percentile(st.responseSizes, 95)
		if err != nil {
			return report.Report{}, fmt.Errorf("can`t calculate 95th percentile of the server response size: %w", err)
		}
	}

	return report.New(
		st.files,
		st.from,
		st.to,
		st.field,
		st.value,
		st.requestsCount,
		st.resources,
		st.codes,
		st.clients,
		st.agents,
		float64(st.totalResponseSize)/float64(st.requestsCount),
		percentile,
	), nil
}
