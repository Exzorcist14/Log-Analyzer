package marker

import "fmt"

type ErrUnknownFormat struct {
	format string
}

func (e ErrUnknownFormat) Error() string {
	return fmt.Sprintf("unknown format (%s) for markup", e.format)
}
