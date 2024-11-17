package analyzer_test

import (
	"testing"
	"time"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/analyzer"
	"github.com/stretchr/testify/assert"
)

func TestCheckTime(t *testing.T) {
	type args struct {
		current         time.Time
		from            time.Time
		to              time.Time
		isFromSpecified bool
		isToSpecified   bool
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "From and To are not specified",
			args: args{
				current:         time.Time{},
				from:            time.Time{},
				to:              time.Time{},
				isFromSpecified: false,
				isToSpecified:   false,
			},
			want: true,
		},
		{
			name: "From is specified, current is later then From",
			args: args{
				current:         time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				from:            time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
				to:              time.Time{},
				isFromSpecified: true,
				isToSpecified:   false,
			},
			want: true,
		},
		{
			name: "From is specified, current is equal to From",
			args: args{
				current:         time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				from:            time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				to:              time.Time{},
				isFromSpecified: true,
				isToSpecified:   false,
			},
			want: true,
		},
		{
			name: "From is specified, current is earlier then From",
			args: args{
				current:         time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
				from:            time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				to:              time.Time{},
				isFromSpecified: true,
				isToSpecified:   false,
			},
			want: false,
		},
		{
			name: "To is specified, current is earlier then To",
			args: args{
				current:         time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
				from:            time.Time{},
				to:              time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				isFromSpecified: false,
				isToSpecified:   true,
			},
			want: true,
		},
		{
			name: "To is specified, current is equal to To",
			args: args{
				current:         time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				from:            time.Time{},
				to:              time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				isFromSpecified: false,
				isToSpecified:   true,
			},
			want: true,
		},
		{
			name: "To is specified, current is later then To",
			args: args{
				current:         time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				from:            time.Time{},
				to:              time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				isFromSpecified: false,
				isToSpecified:   true,
			},
			want: false,
		},
		{
			name: "From and To is specified, current is included in [From; To]",
			args: args{
				current:         time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				from:            time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
				to:              time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				isFromSpecified: true,
				isToSpecified:   true,
			},
			want: true,
		},
		{
			name: "From and To is specified, current is earlier then From",
			args: args{
				current:         time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC),
				from:            time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
				to:              time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				isFromSpecified: true,
				isToSpecified:   true,
			},
			want: false,
		},
		{
			name: "From and To is specified, current is later then To",
			args: args{
				current:         time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
				from:            time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
				to:              time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				isFromSpecified: true,
				isToSpecified:   true,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want,
				analyzer.CheckTime(
					tt.args.current,
					tt.args.from,
					tt.args.to,
					tt.args.isFromSpecified,
					tt.args.isToSpecified,
				),
			)
		})
	}
}
