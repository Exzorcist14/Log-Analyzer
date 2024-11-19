package application

import (
	"fmt"
	"os"
	"time"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/report"
)

type finder interface {
	// Find ищет путь/пути по переданному path, соответствующему локальному шаблоном или url.
	Find(path string) (paths []string, isLocal bool, err error)
}

type analyzer interface {
	// Analyze анализирует логи по указанным путям в соответствии с флагами и возвращает готовый для разметки отчёт.
	Analyze(
		from, to time.Time,
		field, value string,
		read int,
		isFromSpecified, isToSpecified, isFilterSpecified bool,
		paths []string, isLocal bool,
	) (rep report.Report, err error)
}

type marker interface {
	// MarkUp размечает отчёт.
	MarkUp(rep *report.Report, highest int) (markup string)
}

type filer interface {
	// File создаёт файл с расширением соответствующего формата, содержащий размеченный отчёт.
	File(markup, format string) (file *os.File, err error)
}

// Application описывает приложение анализатора.
type Application struct {
	finder   finder
	analyzer analyzer
	marker   marker
	filer    filer
}

// New возвращает инициализированный Application.
func New(finder finder, solver analyzer, packer marker, writer filer) *Application {
	return &Application{
		finder:   finder,
		analyzer: solver,
		marker:   packer,
		filer:    writer,
	}
}

// Run запускает приложение.
func (a *Application) Run(
	path string, from, to time.Time, format, field, value string, highest, read int,
	isFromSpecified, isToSpecified, isFilterSpecified bool,
) error {
	paths, isLocal, err := a.finder.Find(path)
	if err != nil {
		return fmt.Errorf("can`t find paths to files: %w", err)
	}

	rep, err := a.analyzer.Analyze(
		from, to, field, value, read,
		isFromSpecified, isToSpecified, isFilterSpecified,
		paths, isLocal,
	)
	if err != nil {
		return fmt.Errorf("can`t solve: %w", err)
	}

	markup := a.marker.MarkUp(&rep, highest)

	_, err = a.filer.File(markup, format)
	if err != nil {
		return fmt.Errorf("can`t write rep to file: %w", err)
	}

	return nil
}
