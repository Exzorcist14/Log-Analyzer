package finder_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/finder"
	"github.com/stretchr/testify/assert"
)

const (
	comparisonMessage = "checking for a match number of paths"
	matchMessage      = "checking for matching paths"
)

func TestFind(t *testing.T) {
	type args struct {
		path string
	}

	tests := []struct {
		name        string
		args        args
		wantMatches []string
		wantIsLocal bool
	}{
		{
			name: "the path to the file",
			args: args{
				path: `logs/2024-11-07/logs.txt`,
			},
			wantMatches: []string{
				`.*logs[\\/]+2024-11-07[\\/]+logs.txt`,
			},
			wantIsLocal: true,
		},
		{
			name: "the path to the directory",
			args: args{
				path: `logs/2024-11-07`,
			},
			wantMatches: []string{
				`.*logs[\\/]+2024-11-07[\\/]+logs.txt`,
			},
			wantIsLocal: true,
		},
		{
			name: "the path is a local template",
			args: args{
				path: `logs/*`,
			},
			wantMatches: []string{
				`.*logs[\\/]+2024-11-07[\\/]+logs.txt`,
				`.*logs[\\/]+2024-11-08[\\/]+logs.txt`,
			},
			wantIsLocal: true,
		},
		{
			name: "the path is the url",
			args: args{
				path: `https://raw.githubusercontent.com/elastic/examples/master/Common%20Data%20Formats/nginx_logs/nginx_logs`,
			},
			wantMatches: []string{
				`https://raw.githubusercontent.com/elastic/examples/master/Common%20Data%20Formats/nginx_logs/nginx_logs`,
			},
			wantIsLocal: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPaths, gotIsLocal, err := (&finder.Finder{}).Find(tt.args.path)

			assert.True(t, len(gotPaths) == len(tt.wantMatches), comparisonMessage)

			fmt.Println("Some debug information:", tt.name, gotPaths)

			for i, gotPath := range gotPaths {
				re := regexp.MustCompile(tt.wantMatches[i])

				assert.True(t, re.MatchString(gotPath), matchMessage)
			}

			assert.Equal(t, tt.wantIsLocal, gotIsLocal)
			assert.NoError(t, err)
		})
	}
}
