package loader

import "fmt"

// ErrWrongResponseCode - ошибка кода ответа, отличного от кода 200.
type ErrWrongResponseCode struct {
	code int
}

func (e ErrWrongResponseCode) Error() string {
	return fmt.Sprintf("response code %v is not equal to 200", e.code)
}
