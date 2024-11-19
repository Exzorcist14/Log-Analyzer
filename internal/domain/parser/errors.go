package parser

import "fmt"

// ErrNonNginxLog - ошибка строки, не соотвествующий формату строки nginx лога.
type ErrNonNginxLog struct {
	data string
}

func (e ErrNonNginxLog) Error() string {
	return fmt.Sprintf("%s is not an NGINX log", e.data)
}

// ErrNonRequest - ошибка строки, не соотвествующий формату nginx http-запроса.
type ErrNonRequest struct {
	data string
}

func (e ErrNonRequest) Error() string {
	return fmt.Sprintf("%s is not an http-request", e.data)
}
