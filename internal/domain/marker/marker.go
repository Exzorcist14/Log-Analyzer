package marker

import (
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/marker/adoc"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/marker/markdown"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/report"
)

// marker описывает интерфейса разметчика.
type marker interface {
	MarkUp(rep *report.Report, highest int) (markup string)
}

// New как фабрика возвращает конкретную реализацию marker в соответствии с markerType.
func New(markerType string) marker {
	switch markerType {
	case "markdown":
		return &markdown.Marker{}
	case "adoc":
		return &adoc.Marker{}
	default:
		return &markdown.Marker{}
	}
}
