package marker

import (
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/marker/adoc"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/marker/markdown"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/report"
)

type marker interface {
	MarkUp(rep *report.Report, highest int) (markup string)
}

func New(marker string) marker {
	switch marker {
	case "markdown":
		return &markdown.Marker{}
	case "adoc":
		return &adoc.Marker{}
	default:
		return &markdown.Marker{}
	}
}
