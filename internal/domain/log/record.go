package log

import (
	"time"
)

// Request - представление http-запроса.
type Request struct {
	Method   string
	Resource string
	Protocol string
}

// Record - промежуточное представление строки nginx лога.
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
