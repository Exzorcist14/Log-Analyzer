package parser_test

import (
	"testing"
	"time"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/log"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/parser"
	"github.com/stretchr/testify/assert"
)

const (
	layout   = "02/Jan/2006:15:04:05 -0700"
	nginxLog = `244.103.237.229 - - [17/Nov/2024:16:07:52 +0000] "GET /reciprocal.hmtl HTTP/1.1" 200 2420 "-" ` +
		`"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/5331 (KHTML, like Gecko) Chrome/36.0.897.0 Mobile Safari/5331"`
	nonNginxLog = `145.127.201.137 - [17/Nov/2024:12:59:30 +0000] "POST ` +
		`/Synergistic/encompassing.gif HTTP/1.1" 200 806 "-" "Opera/9.52 (X11; Linux i686; en-US) Presto/2.12.195 Version/13.00"`
)

func TestParse(t *testing.T) {
	firstTime, err := time.Parse(layout, "17/Nov/2024:16:07:52 +0000")
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name    string
		lg      string
		want    *log.Record
		wantErr bool
	}{
		{
			name: "nginx log",
			lg:   nginxLog,
			want: &log.Record{
				RemoteAddr: "244.103.237.229",
				RemoteUser: "-",
				TimeLocal:  firstTime,
				Request: log.Request{
					Method:   "GET",
					Resource: "/reciprocal.hmtl",
					Protocol: "HTTP/1.1",
				},
				Status:        200,
				BodyBytesSent: 2420,
				HTTPRefer:     "-",
				HTTPUserAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/5331 (KHTML, like Gecko) Chrome/36.0.897.0 Mobile Safari/5331",
			},
			wantErr: false,
		},
		{
			name:    "non-nginx log",
			lg:      nonNginxLog,
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := (&parser.Parser{}).Parse(tt.lg)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}
