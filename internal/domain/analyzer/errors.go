package analyzer

import "fmt"

// ErrUnknownField - ошибка неизвестного фильтруемого поля строки лога.
type ErrUnknownField struct {
	field string
}

func (e ErrUnknownField) Error() string {
	return fmt.Sprintf("%s is not a known field", e.field)
}
