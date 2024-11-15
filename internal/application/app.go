package application

import (
	"fmt"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/report"
)

type finder interface {
	Find(path string) (paths []string, isLocal bool, err error)
}

type analyzer interface {
	Analyze(
		from, to, field, value string,
		isFromSpecified, isToSpecified, isFilterSpecified bool,
		paths []string, isLocal bool,
	) (statistics report.Report, err error)
}

type marker interface {
	MarkUp(rep *report.Report, highest int) (markup string)
}

type filer interface {
	File(markup, format string) error
}

type Application struct {
	finder   finder
	analyzer analyzer
	marker   marker
	filer    filer
}

func New(finder finder, solver analyzer, packer marker, writer filer) *Application {
	return &Application{
		finder:   finder,
		analyzer: solver,
		marker:   packer,
		filer:    writer,
	}
}

func (a *Application) Run(
	path, from, to, format, field, value string, highest int,
	isFromSpecified, isToSpecified, isFilterSpecified bool,
) error {
	paths, isLocal, err := a.finder.Find(path)
	if err != nil {
		return fmt.Errorf("can`t find paths to files: %w", err)
	}

	rep, err := a.analyzer.Analyze(
		from, to, field, value,
		isFromSpecified, isToSpecified, isFilterSpecified,
		paths, isLocal,
	)
	if err != nil {
		return fmt.Errorf("can`t solve: %w", err)
	}

	markup := a.marker.MarkUp(&rep, highest)

	err = a.filer.File(markup, format)
	if err != nil {
		return fmt.Errorf("can`t write rep to file: %w", err)
	}

	return nil
}
