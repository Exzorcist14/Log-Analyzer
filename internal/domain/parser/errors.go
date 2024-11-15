package parser

import "fmt"

type ErrNonNginxLog struct {
	data string
}

func (e ErrNonNginxLog) Error() string {
	return fmt.Sprintf("%s is not an NGINX log", e.data)
}

type ErrNonRequest struct {
	data string
}

func (e ErrNonRequest) Error() string {
	return fmt.Sprintf("%s is not an http-request", e.data)
}
