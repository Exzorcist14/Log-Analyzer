package analyzer

import "fmt"

// ErrWrongResponseCode - ошибка кода ответа, отличного от кода 200.
type ErrWrongResponseCode struct {
	code int
}

func (e ErrWrongResponseCode) Error() string {
	return fmt.Sprintf("response code %v is not equal to 200", e.code)
}

// ErrUnknownField - ошибка неизвестного фильтруемого поля строки лога.
type ErrUnknownField struct {
	field string
}

func (e ErrUnknownField) Error() string {
	return fmt.Sprintf("%s is not a known field", e.field)
}
