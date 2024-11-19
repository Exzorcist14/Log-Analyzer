package filer

import "fmt"

// ErrUnknownFormat - ошибка неизвестного формата, в котором нужно сохранить.
type ErrUnknownFormat struct {
	format string
}

func (e ErrUnknownFormat) Error() string {
	return fmt.Sprintf("unknown format (%s) for writing", e.format)
}
