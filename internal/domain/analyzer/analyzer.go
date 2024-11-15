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

const layout = "2006-01-02T15:04:05Z07:00"

type parser interface {
	Parse(log string) (*log.Record, error)
}

type interimStatistics struct {
	from              string
	to                string
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
	stats             interimStatistics
	from              time.Time
	to                time.Time
	field             string
	value             string
	isFromSpecified   bool
	isToSpecified     bool
	isFilterSpecified bool
}

func New(ps parser) *Analyzer {
	return &Analyzer{
		parser: ps,
		stats: interimStatistics{
			resources: make(map[string]int),
			codes:     make(map[int]int),
			clients:   make(map[string]int),
			agents:    make(map[string]int),
		},
	}
}

func (s *Analyzer) Analyze(
	from, to, field, value string,
	isFromSpecified, isToSpecified, isFilterSpecified bool,
	paths []string, isLocal bool,
) (rep report.Report, err error) {
	s.field = field
	s.value = value
	s.isFromSpecified = isFromSpecified
	s.isToSpecified = isToSpecified
	s.isFilterSpecified = isFilterSpecified
	s.stats.from = from
	s.stats.to = to

	if isFromSpecified {
		s.from, err = time.Parse(layout, from)
		if err != nil {
			return rep, fmt.Errorf("can`t to parse -from: %v", err)
		}
	}

	if isToSpecified {
		s.to, err = time.Parse(layout, to)
		if err != nil {
			return rep, fmt.Errorf("can`t to parse -to: %v", err)
		}
	}

	if isLocal {
		for _, path := range paths {
			s.stats.files = append(s.stats.files, path)

			err = s.processLocalLogFile(path)
			if err != nil {
				return rep, fmt.Errorf("can`t process log file: %w", err)
			}
		}
	} else {
		for _, path := range paths {
			s.stats.files = append(s.stats.files, path)

			err = s.processRemoteLogFile(path)
			if err != nil {
				return rep, fmt.Errorf("can`t process log file: %w", err)
			}
		}
	}

	rep, err = generateReport(&s.stats)
	if err != nil {
		return rep, fmt.Errorf("can`t generate report: %w", err)
	}

	return rep, nil
}

func (s *Analyzer) processLocalLogFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("can`t open local log file: %w", err)
	}
	defer file.Close()

	err = s.addToStatisticsFromLog(file)
	if err != nil {
		return fmt.Errorf("can`t add log to interim statistics: %w", err)
	}

	return nil
}

func (s *Analyzer) processRemoteLogFile(path string) error {
	parsedURL, err := url.Parse(path)
	if err != nil {
		return fmt.Errorf("can`t parse urle: %w", err)
	}

	resp, err := http.Get(parsedURL.String())
	if err != nil {
		return fmt.Errorf("can`t make GET request: %w", err)
	} else if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("can`t make GET request: %w", ErrWrongResponseCode{resp.StatusCode})
	}
	defer resp.Body.Close()

	err = s.addToStatisticsFromLog(resp.Body)
	if err != nil {
		return fmt.Errorf("can`t add to interim statistics: %w", err)
	}

	return nil
}

func (s *Analyzer) addToStatisticsFromLog(lg io.Reader) error {
	scn := bufio.NewScanner(lg)

	for scn.Scan() {
		logRecord, err := s.parser.Parse(scn.Text())
		if err != nil {
			return fmt.Errorf("can`t parse scan result: %w", err)
		}

		isCheckSuccessful, err := s.check(logRecord)
		if err != nil {
			return fmt.Errorf("can't check the lg to satisfy the conditions: %w", err)
		}

		if isCheckSuccessful {
			s.addToStatisticsFromLogRecord(logRecord)
		}
	}

	if err := scn.Err(); err != nil {
		return fmt.Errorf("can`t scan: %w", err)
	}

	return nil
}

func (s *Analyzer) addToStatisticsFromLogRecord(logRecord *log.Record) {
	s.stats.requestsCount++
	s.stats.resources[logRecord.Request.Resource]++
	s.stats.codes[logRecord.Status]++
	s.stats.clients[logRecord.RemoteAddr]++
	s.stats.agents[logRecord.HTTPUserAgent]++
	s.stats.responseSizes = append(s.stats.responseSizes, float64(logRecord.BodyBytesSent))
	s.stats.totalResponseSize += logRecord.BodyBytesSent
}

func (s *Analyzer) check(record *log.Record) (bool, error) {
	var err error

	isTimeSuccessful := true
	isFilterSuccessful := true

	if s.isFromSpecified || s.isToSpecified {
		isTimeSuccessful = checkTime(record.TimeLocal, s.from, s.to, s.isFromSpecified, s.isToSpecified)
	}

	if s.isFilterSpecified {
		isFilterSuccessful, err = checkFilter(record, s.field, s.value)
		if err != nil {
			return false, fmt.Errorf("can`t check filter: %w", err)
		}
	}

	return isTimeSuccessful && isFilterSuccessful, nil
}

func checkTime(current, from, to time.Time, isFromSpecified, isToSpecified bool) bool {
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

func generateReport(is *interimStatistics) (report.Report, error) {
	var (
		percentile float64
		err        error
	)

	if len(is.responseSizes) != 0 {
		percentile, err = stats.Percentile(is.responseSizes, 95)
		if err != nil {
			return report.Report{}, fmt.Errorf("can`t calculate 95th percentile of the server response size: %w", err)
		}
	}

	return report.New(
		is.files,
		is.from,
		is.to,
		is.requestsCount,
		transformMapToSlice(is.resources),
		transformMapToSlice(is.codes),
		transformMapToSlice(is.clients),
		transformMapToSlice(is.agents),
		float64(is.totalResponseSize)/float64(is.requestsCount),
		percentile,
	), nil
}

func transformMapToSlice[T string | int](mp map[T]int) []report.DataWithCount[T] {
	slice := []report.DataWithCount[T]{}

	for data, count := range mp {
		slice = append(slice, report.DataWithCount[T]{
			Data:  data,
			Count: count,
		})
	}

	return slice
}
