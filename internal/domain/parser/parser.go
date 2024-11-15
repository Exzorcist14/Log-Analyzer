package parser

import (
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/log"
)

const layout = "02/Jan/2006:15:04:05 -0700"

type Parser struct{}

func New() *Parser {
	return &Parser{}
}

func (p *Parser) Parse(lg string) (*log.Record, error) {
	logRegExp := regexp.MustCompile(
		`(?P<RemoteAddr>.*) - (?P<RemoteUser>.*) ` +
			`\[(?P<TimeLocal>.*)\] "(?P<Request>.*)" ` +
			`(?P<Status>.*) (?P<BodyBytesSent>.*) ` +
			`"(?P<HTTPRefer>.*)" "(?P<HTTPUserAgent>.*)"`,
	)

	match := logRegExp.FindStringSubmatch(lg)
	if !logRegExp.MatchString(lg) || match == nil {
		return nil, fmt.Errorf("can`t find string submatch for log: %w", ErrNonNginxLog{lg})
	}

	result := make(map[string]string)

	for i, name := range logRegExp.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}

	record := log.Record{
		RemoteAddr:    result["RemoteAddr"],
		RemoteUser:    result["RemoteUser"],
		HTTPRefer:     result["HTTPRefer"],
		HTTPUserAgent: result["HTTPUserAgent"],
	}

	request, err := parseRequest(result["Request"])
	if err != nil {
		return nil, fmt.Errorf("can`t parse request: %w", err)
	}

	record.Request = request

	timeLocal, err := time.Parse(layout, result["TimeLocal"])
	if err != nil {
		return nil, fmt.Errorf("can`t parse time: %w", err)
	}

	record.TimeLocal = timeLocal

	status, err := strconv.Atoi(result["Status"])
	if err != nil {
		return nil, fmt.Errorf("can`t parse status: %w", err)
	}

	record.Status = status

	bodyBytesSend, err := strconv.Atoi(result["BodyBytesSent"])
	if err != nil {
		return nil, fmt.Errorf("can`t parse body bytes sent: %w", err)
	}

	record.BodyBytesSent = bodyBytesSend

	return &record, nil
}

func parseRequest(request string) (log.Request, error) {
	reqRexEpx := regexp.MustCompile(`^(\w+)\s+(\S+)\s+(HTTP/\d\.\d)$`)

	parts := reqRexEpx.FindStringSubmatch(request)
	if len(parts) == 0 {
		return log.Request{}, ErrNonRequest{request}
	}

	return log.Request{
		Method:   parts[1],
		Resource: parts[2],
		Protocol: parts[3],
	}, nil
}
