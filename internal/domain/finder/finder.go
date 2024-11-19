package finder

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

const (
	projectDirectory = "backend_academy_2024_project_3-go-Exzorcist14" // Название директории проекта.
	relativePath     = "/internal/infrastructure/"                     // Относительный путь от проекта к директории с файлами.
)

// Finder умеет находить пути.
type Finder struct{}

// Find возвращает все пути, соответствующие path, который может быть представлен локальным шаблоном или url.
// Если path локальный, то в качестве второго значения возвращает true, иначе - false.
func (f *Finder) Find(path string) (paths []string, isLocal bool, err error) {
	urlRegExp := regexp.MustCompile(`^(https?://)?([a-zA-Z0-9-]+\.)+[a-zA-Z]{2,6}(:\d+)?(/[^\s]*)?$`)

	if !urlRegExp.MatchString(path) { // Если путь не содержит url.
		paths, err = findByLocalPath(path)
		if err != nil {
			return nil, false, fmt.Errorf("can`t find by local path: %v", err)
		}

		isLocal = true
	} else {
		paths = []string{path}
		isLocal = false
	}

	return paths, isLocal, nil
}

// findByLocalPath ищет все локальные пути, соответствующие шаблону path.
func findByLocalPath(path string) ([]string, error) {
	prefix, err := getAbsolutePrefix(path)
	if err != nil {
		return nil, fmt.Errorf("can`t get absolute prefix: %v", err)
	}

	postfix := getAbsolutePostfix(path)

	paths, err := filepath.Glob(prefix + path + postfix)
	if err != nil {
		return nil, fmt.Errorf("can`t glob path: %v", err)
	}

	return paths, nil
}

// getAbsolutePrefix возвращает абсолютный префикс для пути к файлу.
// Добавляет к названию файла абсолютный путь до проекта и относительный путь от проекта до директории с файлами.
func getAbsolutePrefix(path string) (string, error) {
	prefix := ""

	if !filepath.IsAbs(path) {
		absolutePath, err := os.Getwd()
		if err != nil {
			return "", fmt.Errorf("can`t get current directory: %v", err)
		}

		for filepath.Base(absolutePath) != projectDirectory {
			err = os.Chdir("..")
			if err != nil {
				return "", fmt.Errorf("can`t change directory: %v", err)
			}

			absolutePath, err = os.Getwd()
			if err != nil {
				return "", fmt.Errorf("can`t get current directory: %v", err)
			}
		}

		prefix = absolutePath + relativePath
	}

	return prefix, nil
}

// getAbsolutePostfix возвращает абсолютный постфикс для пути к файлу.
// Если path не является путём к файлу, формирует постфикс-шаблон для всех лежащих внутри txt-файлов.
func getAbsolutePostfix(path string) (postfix string) {
	fileRegExp := regexp.MustCompile(`^(.*/)?([^/]+)\.txt$`)

	if !fileRegExp.MatchString(path) {
		postfix = "/*.txt"
	}

	return postfix
}
