package marker

import "fmt"

// ErrUnknownFormat - ошибка строки, обозначающей неизвестный язык разметки.
type ErrUnknownFormat struct {
	format string
}

func (e ErrUnknownFormat) Error() string {
	return fmt.Sprintf("unknown format (%s) for markup", e.format)
}
