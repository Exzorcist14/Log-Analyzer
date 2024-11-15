package filer

import (
	"fmt"
	"os"
)

const relativePath = "/internal/infrastructure/reports/"

type Filer struct{}

func New() *Filer {
	return &Filer{}
}

func (w *Filer) File(markup, format string) error {
	var name string

	switch format {
	case "markdown":
		name = "report.md"
	case "adoc":
		name = "report.adoc"
	default:
		return ErrUnknownFormat{format}
	}

	path, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("can`t get current working directory: %w", err)
	}

	err = os.Chdir(path + relativePath)
	if err != nil {
		return fmt.Errorf("can`t change working directory: %w", err)
	}

	file, err := os.Create(name)
	if err != nil {
		return fmt.Errorf("can`t create file: %w", err)
	}
	defer file.Close()

	fmt.Fprint(file, markup)

	return nil
}
