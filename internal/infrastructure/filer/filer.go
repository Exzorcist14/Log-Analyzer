package filer

import (
	"fmt"
	"os"
)

const relativePath = "/internal/infrastructure/reports/" // Относительный путь от проекта к директории, куда необходимо сохранить файл.

// Filer умеет сохранять файл с размеченным отчётом.
type Filer struct{}

// File сохраняет файл соответствующего расширения с записанным в него размеченным отчётом, возвращая указатель на него.
func (w *Filer) File(markup, format string) (*os.File, error) {
	var name string

	switch format {
	case "markdown":
		name = "report.md"
	case "adoc":
		name = "report.adoc"
	default:
		return nil, ErrUnknownFormat{format}
	}

	path, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("can`t get current working directory: %w", err)
	}

	err = os.Chdir(path + relativePath)
	if err != nil {
		return nil, fmt.Errorf("can`t change working directory: %w", err)
	}

	file, err := os.Create(name)
	if err != nil {
		return nil, fmt.Errorf("can`t create file: %w", err)
	}
	defer file.Close()

	fmt.Fprint(file, markup)

	return file, nil
}
