package analyzer

import "fmt"

type ErrWrongResponseCode struct {
	code int
}

func (e ErrWrongResponseCode) Error() string {
	return fmt.Sprintf("response code %v is not equal to 200", e.code)
}

type ErrUnknownField struct {
	field string
}

func (e ErrUnknownField) Error() string {
	return fmt.Sprintf("%s is not a known field", e.field)
}
