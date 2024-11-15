package main

import (
	"flag"
	"os"

	"github.com/es-debug/backend-academy-2024-go-template/internal/application"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/analyzer"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/finder"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/marker"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/parser"
	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure/filer"
)

const (
	defaultPath    = "-"
	defaultFrom    = "-"
	defaultTo      = "-"
	defaultFormat  = "markdown"
	defaultField   = "-"
	defaultValue   = "-"
	defaultHighest = 3
	pathUsage      = "path to the log files"
	fromUsage      = "the minimum time that must be exceeded by the time the log is recorded for analysis"
	toUsage        = "the maximum time that must exceed the time of recording the log in order for it to be analyzed"
	formatUsage    = "output format (available formats: markdown, adoc)"
	fieldUsage     = "Filter by nginx log field (available filters: remote_add, remote_user, time_local, " +
		"method, resource, protocol, status, body_bytes_sent, http_referer, http_user_agent)"
	valueUsage   = "The value of the filter field"
	highestUsage = "the number of the most common instances of characteristics that should be displayed on the screen" +
		" (if the available number of instances is exceeded, all are displayed)"
)

func main() {
	path := flag.String("path", defaultPath, pathUsage)
	from := flag.String("from", defaultFrom, fromUsage)
	to := flag.String("to", defaultTo, toUsage)
	format := flag.String("format", defaultFormat, formatUsage)
	field := flag.String("filter-field", defaultField, fieldUsage)
	value := flag.String("filter-value", defaultValue, valueUsage)
	highest := flag.Int("highest", defaultHighest, highestUsage)

	flag.Parse()

	if !areFlagValuesValid(*path, *format, *field, *value) {
		os.Exit(1)
	}

	anlz := application.New(finder.New(), analyzer.New(parser.New()), marker.New(*format), filer.New())

	err := anlz.Run(
		*path, *from, *to, *format, *field, *value, *highest,
		*from != defaultFrom, *to != defaultTo, *field != defaultPath,
	)

	if err != nil {
		os.Exit(1)
	}
}

func areFlagValuesValid(path, format, field, value string) bool {
	fields := map[string]bool{
		"remote_add":      true,
		"remote_user":     true,
		"time_local":      true,
		"method":          true,
		"resource":        true,
		"protocol":        true,
		"status":          true,
		"body_bytes_sent": true,
		"http_referer":    true,
		"http_user_agent": true,
	}

	formats := map[string]bool{
		"markdown": true,
		"adoc":     true,
	}

	if path == defaultPath {
		return false
	}

	if _, ok := formats[format]; !ok {
		return false
	}

	if _, ok := fields[field]; !ok && field != defaultField || ok && value == defaultValue {
		return false
	}

	return true
}
