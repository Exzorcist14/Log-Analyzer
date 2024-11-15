package log

import (
	"time"
)

type Request struct {
	Method   string
	Resource string
	Protocol string
}

type Record struct {
	RemoteAddr    string
	RemoteUser    string
	TimeLocal     time.Time
	Request       Request
	Status        int
	BodyBytesSent int
	HTTPRefer     string
	HTTPUserAgent string
}
