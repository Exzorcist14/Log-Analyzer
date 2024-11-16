package finder

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

const relativePath = "/internal/infrastructure/"

type Finder struct{}

func (f *Finder) Find(path string) (paths []string, isLocal bool, err error) {
	urlRegExp := regexp.MustCompile(`^(https?://)?([a-zA-Z0-9-]+\.)+[a-zA-Z]{2,6}(:\d+)?(/[^\s]*)?$`)

	if !urlRegExp.MatchString(path) {
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

func getAbsolutePrefix(path string) (string, error) {
	prefix := ""

	if !filepath.IsAbs(path) {
		absolutePath, err := os.Getwd()
		if err != nil {
			return "", fmt.Errorf("can`t get current directory: %v", err)
		}

		prefix = absolutePath + relativePath
	}

	return prefix, nil
}

func getAbsolutePostfix(path string) (postfix string) {
	fileRegExp := regexp.MustCompile(`^(.*/)?([^/]+)\.txt$`)

	if !fileRegExp.MatchString(path) {
		postfix = "/*.txt"
	}

	return postfix
}
