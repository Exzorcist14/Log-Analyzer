package loader

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

// Loader умеет загружать данные для чтения.
type Loader struct{}

// Load загружает данные для чтения.
func (l *Loader) Load(path string, isLocal bool) (source io.ReadCloser, err error) {
	if isLocal {
		source, err = loadLocal(path)
		if err != nil {
			return nil, fmt.Errorf("can`t loadl local file: %v", err)
		}
	} else {
		source, err = loadRemote(path)
		if err != nil {
			return nil, fmt.Errorf("can`t load remote file: %v", err)
		}
	}

	return source, nil
}

// loadLocal загружает локальные данные.
func loadLocal(path string) (*os.File, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("can`t open the file: %w", err)
	}

	return file, nil
}

// loadRemote загружает удалённые данные.
func loadRemote(path string) (io.ReadCloser, error) {
	parsedURL, err := url.Parse(path)
	if err != nil {
		return nil, fmt.Errorf("can`t parse url: %w", err)
	}

	resp, err := http.Get(parsedURL.String())
	if err != nil {
		return nil, fmt.Errorf("can`t make GET request: %w", err)
	} else if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("can`t make GET request: %w", ErrWrongResponseCode{resp.StatusCode})
	}

	return resp.Body, nil
}
