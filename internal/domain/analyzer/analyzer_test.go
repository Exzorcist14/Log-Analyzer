package analyzer_test

import (
	"testing"
	"time"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/analyzer"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/finder"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/parser"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/report"
	"github.com/stretchr/testify/assert"
)

func TestAnalyze(t *testing.T) {
	f := finder.Finder{}

	patternPaths, patternIsLocal, err := f.Find(`logs/*`)
	assert.NoError(t, err)

	urlPath, urlIsLocal, err := f.Find(`https://raw.githubusercontent.com/elastic/` +
		`examples/master/Common%20Data%20Formats/nginx_logs/nginx_logs`)
	assert.NoError(t, err)

	localPath1, isLocal1, err := f.Find(`logs/2024-11-07`)
	assert.NoError(t, err)

	localPath2, isLocal2, err := f.Find(`logs/2024-11-08`)
	assert.NoError(t, err)

	type args struct {
		from              time.Time
		to                time.Time
		field             string
		value             string
		read              int
		isFromSpecified   bool
		isToSpecified     bool
		isFilterSpecified bool
		paths             []string
		isLocal           bool
	}

	tests := []struct {
		name    string
		args    args
		wantRep report.Report
	}{
		{
			name: "local paths without flags",
			args: args{
				read:    4,
				paths:   patternPaths,
				field:   "-",
				value:   "-",
				isLocal: patternIsLocal,
			},
			wantRep: report.New(
				patternPaths,
				"-",
				"-",
				"-",
				"-",
				8,
				map[string]int{
					"/Organized-open%20system/intranet.jpg":           1,
					"/homogeneous/customer%20loyalty/bottom-line.css": 1,
					"/Phased.php": 1,
					"/Monitored-Streamlined%20national/logistical.svg": 1,
					"/extranet/Cross-platform.jpg":                     1,
					"/productivity/core.svg":                           1,
					"/composite.js":                                    1,
					"/interface/knowledge%20user/asynchronous/task-force-Graphical%20User%20Interface.htm": 1,
				},
				map[int]int{
					200: 6,
					301: 2,
				},
				map[string]int{
					"89.225.138.240":  1,
					"153.84.179.223":  1,
					"78.120.107.169":  1,
					"197.154.168.123": 1,
					"20.19.155.88":    1,
					"81.69.92.165":    1,
					"135.3.66.63":     1,
					"70.27.134.194":   1,
				},
				map[string]int{
					"Opera/8.62 (Windows 98; Win 9x 4.90; en-US) Presto/2.10.183 Version/12.00":                                 1,
					"Opera/9.55 (Macintosh; PPC Mac OS X 10_6_10; en-US) Presto/2.10.224 Version/12.00":                         1,
					"Mozilla/5.0 (X11; Linux x86_64; rv:6.0) Gecko/1982-11-05 Firefox/37.0":                                     1,
					"Mozilla/5.0 (Windows NT 5.0) AppleWebKit/5360 (KHTML, like Gecko) Chrome/36.0.853.0 Mobile Safari/5360":    1,
					"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/5351 (KHTML, like Gecko) Chrome/39.0.889.0 Mobile Safari/5351": 1,
					"Mozilla/5.0 (X11; Linux i686; rv:8.0) Gecko/1945-02-11 Firefox/37.0":                                       1,
					"Mozilla/5.0 (X11; Linux i686; rv:8.0) Gecko/1990-04-04 Firefox/37.0":                                       1,
					"Mozilla/5.0 (Windows; U; Windows 98) AppleWebKit/534.37.5 (KHTML, like Gecko) Version/5.0 Safari/534.37.5": 1,
				},
				1453.25,
				2917,
			),
		},
		{
			name: "url without flags",
			args: args{
				read:    10,
				paths:   urlPath,
				field:   "-",
				value:   "-",
				isLocal: urlIsLocal,
			},
			wantRep: report.New(
				urlPath,
				"-",
				"-",
				"-",
				"-",
				10,
				map[string]int{
					"/downloads/product_1": 8,
					"/downloads/product_2": 2,
				},
				map[int]int{
					304: 6,
					200: 2,
					404: 2,
				},
				map[string]int{
					"217.168.17.5": 4,
					"93.180.71.3":  4,
					"80.91.33.133": 2,
				},
				map[string]int{
					"Debian APT-HTTP/1.3 (0.8.10.3)":                4,
					"Debian APT-HTTP/1.3 (0.8.16~exp12ubuntu10.21)": 4,
					"Debian APT-HTTP/1.3 (0.8.16~exp12ubuntu10.17)": 2,
				},
				164.9,
				490,
			),
		},
		{
			name: "local path with -from",
			args: args{
				read:            10,
				paths:           localPath1,
				from:            time.Date(2024, 11, 7, 16, 7, 56, 0, time.FixedZone("+0000", 0)),
				field:           "-",
				value:           "-",
				isLocal:         isLocal2,
				isFromSpecified: true,
			},
			wantRep: report.New(
				localPath1,
				"2024-11-07 16:07:56 +0000 +0000",
				"-",
				"-",
				"-",
				10,
				map[string]int{
					"/real-time/Organized_tertiary/Vision-oriented.js":                      1,
					"/frame-service-desk/application.css":                                   1,
					"/regional/intermediate-leverage_contextually-based.png":                1,
					"/zero%20tolerance_emulation.svg":                                       1,
					"/Diverse/synergy/system-worthy%20full-range.hmtl":                      1,
					"/Multi-tiered-coherent-process%20improvement%20impactful/neutral.hmtl": 1,
					"/disintermediate-Innovative.js":                                        1,
					"/forecast/Persevering-hardware-algorithm.php":                          1,
					"/paradigm/Visionary/project/modular/Implemented.svg":                   1,
					"/system-worthy-motivating%20solution/static/Enterprise-wide.svg":       1,
				},
				map[int]int{
					200: 10,
				},
				map[string]int{
					"167.153.110.201": 1,
					"196.46.49.13":    1,
					"130.119.158.208": 1,
					"206.232.12.89":   1,
					"239.95.216.124":  1,
					"150.87.216.31":   1,
					"162.115.191.162": 1,
					"32.61.217.122":   1,
					"59.177.209.163":  1,
					"197.14.67.27":    1,
				},
				map[string]int{
					"Opera/9.55 (X11; Linux i686; en-US) Presto/2.12.198 Version/10.00":                                         1,
					"Opera/10.90 (Macintosh; PPC Mac OS X 10_9_2; en-US) Presto/2.10.247 Version/11.00":                         1,
					"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/5320 (KHTML, like Gecko) Chrome/37.0.838.0 Mobile Safari/5320": 1,
					"Mozilla/5.0 (X11; Linux i686) AppleWebKit/5350 (KHTML, like Gecko) Chrome/38.0.844.0 Mobile Safari/5350":   1,
					"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_9_7) AppleWebKit/5320 (KHTML, like Gecko)" +
						" Chrome/39.0.803.0 Mobile Safari/5320": 1,
					"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_5_10 rv:7.0; en-US) AppleWebKit/536.35.6 (KHTML, like Gecko)" +
						" Version/6.2 Safari/536.35.6": 1,
					"Mozilla/5.0 (X11; Linux i686; rv:6.0) Gecko/1928-14-10 Firefox/35.0":                                       1,
					"Mozilla/5.0 (X11; Linux i686; rv:7.0) Gecko/1968-27-11 Firefox/35.0":                                       1,
					"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/5330 (KHTML, like Gecko) Chrome/38.0.877.0 Mobile Safari/5330": 1,
					"Opera/8.24 (Macintosh; Intel Mac OS X 10_5_3; en-US) Presto/2.11.348 Version/12.00":                        1,
				},
				2001.6,
				2866,
			),
		},
		{
			name: "local path with -to",
			args: args{
				read:          10,
				paths:         localPath1,
				to:            time.Date(2024, 11, 7, 16, 7, 56, 0, time.FixedZone("+0000", 0)),
				field:         "-",
				value:         "-",
				isLocal:       isLocal1,
				isToSpecified: true,
			},
			wantRep: report.New(
				localPath1,
				"-",
				"2024-11-07 16:07:56 +0000 +0000",
				"-",
				"-",
				10,
				map[string]int{
					"/middleware-disintermediate%20intangible_Reduced.js":             1,
					"/database/De-engineered-intermediate/local%20area%20network.jpg": 1,
					"/architecture/focus%20group.js":                                  1,
					"/composite.js":                                                   1,
					"/Monitored-Streamlined%20national/logistical.svg":                1,
					"/model.png":                   1,
					"/extranet/Cross-platform.jpg": 1,
					"/needs-based_frame-leading%20edge-budgetary%20management-protocol.php": 1,
					"/open%20architecture.js": 1,
					"/productivity/core.svg":  1,
				},
				map[int]int{
					200: 8,
					301: 1,
					500: 1,
				},
				map[string]int{
					"128.220.222.217": 1,
					"153.84.179.223":  1,
					"67.6.173.187":    1,
					"70.27.134.194":   1,
					"78.120.107.169":  1,
					"89.225.138.240":  1,
					"237.217.84.116":  1,
					"57.136.240.250":  1,
					"156.242.41.143":  1,
					"173.19.234.43":   1,
				},
				map[string]int{
					"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_2 rv:5.0) Gecko/1960-16-09 Firefox/35.0": 1,
					"Opera/9.55 (Macintosh; PPC Mac OS X 10_6_10; en-US) Presto/2.10.224 Version/12.00":   1,
					"Mozilla/5.0 (Macintosh; U; PPC Mac OS X 10_5_6) AppleWebKit/5352 (KHTML, like Gecko)" +
						" Chrome/37.0.872.0 Mobile Safari/5352": 1,
					"Opera/8.62 (Windows 98; Win 9x 4.90; en-US) Presto/2.10.183 Version/12.00":                                 1,
					"Mozilla/5.0 (X11; Linux x86_64; rv:8.0) Gecko/1924-14-09 Firefox/36.0":                                     1,
					"Mozilla/5.0 (Windows 98; en-US; rv:1.9.0.20) Gecko/1994-19-02 Firefox/36.0":                                1,
					"Mozilla/5.0 (Windows NT 5.0) AppleWebKit/5360 (KHTML, like Gecko) Chrome/36.0.853.0 Mobile Safari/5360":    1,
					"Mozilla/5.0 (X11; Linux x86_64; rv:6.0) Gecko/1982-11-05 Firefox/37.0":                                     1,
					"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/5361 (KHTML, like Gecko) Chrome/40.0.808.0 Mobile Safari/5361": 1,
					"Mozilla/5.0 (X11; Linux i686) AppleWebKit/5332 (KHTML, like Gecko) Chrome/39.0.876.0 Mobile Safari/5332":   1,
				},
				1439.2,
				2720,
			),
		},
		{
			name: "local path with -from and -to",
			args: args{
				read:            10,
				paths:           localPath2,
				from:            time.Date(2024, 11, 8, 14, 39, 44, 0, time.FixedZone("+0000", 0)),
				to:              time.Date(2024, 11, 8, 14, 40, 3, 0, time.FixedZone("+0000", 0)),
				field:           "-",
				value:           "-",
				isLocal:         isLocal2,
				isFromSpecified: true,
				isToSpecified:   true,
			},
			wantRep: report.New(
				localPath2,
				"2024-11-08 14:39:44 +0000 +0000",
				"2024-11-08 14:40:03 +0000 +0000",
				"-",
				"-",
				10,
				map[string]int{
					"/6th%20generation.php":                                            1,
					"/approach/Adaptive%20infrastructure.gif":                          1,
					"/Synchronised-secured%20line.svg":                                 1,
					"/approach.hmtl":                                                   1,
					"/database/object-oriented%20neural-net%20Virtual%20Universal.htm": 1,
					"/encryption/Virtual.png":                                          1,
					"/interactive.css":                                                 1,
					"/reciprocal/Synergistic/transitional.hmtl":                        1,
					"/migration/human-resource%20Organized/Stand-alone.png":            1,
					"/dedicated/help-desk%20Exclusive/didactic%20structure.php":        1,
				},
				map[int]int{
					200: 7,
					404: 2,
					400: 1,
				},
				map[string]int{
					"215.215.249.196": 1,
					"232.22.135.239":  1,
					"76.219.118.160":  1,
					"161.123.208.218": 1,
					"93.46.164.250":   1,
					"147.145.53.17":   1,
					"107.1.82.37":     1,
					"169.145.174.150": 1,
					"196.101.108.15":  1,
					"21.45.234.173":   1,
				},
				map[string]int{
					"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_6_7) AppleWebKit/5322 (KHTML, like Gecko)" +
						" Chrome/38.0.821.0 Mobile Safari/5322": 1,
					"Opera/10.13 (Macintosh; PPC Mac OS X 10_9_3; en-US) Presto/2.10.206 Version/12.00":                         1,
					"Opera/10.78 (Macintosh; U; Intel Mac OS X 10_8_8; en-US) Presto/2.12.254 Version/12.00":                    1,
					"Opera/10.10 (X11; Linux x86_64; en-US) Presto/2.10.230 Version/13.00":                                      1,
					"Mozilla/5.0 (Windows NT 6.2; en-US; rv:1.9.0.20) Gecko/1976-07-11 Firefox/35.0":                            1,
					"Mozilla/5.0 (Windows; U; Windows CE) AppleWebKit/532.41.4 (KHTML, like Gecko) Version/5.1 Safari/532.41.4": 1,
					"Mozilla/5.0 (Macintosh; PPC Mac OS X 10_9_1 rv:7.0) Gecko/2024-10-07 Firefox/37.0":                         1,
					"Mozilla/5.0 (Macintosh; U; PPC Mac OS X 10_7_8) AppleWebKit/5351 (KHTML, like Gecko)" +
						" Chrome/36.0.834.0 Mobile Safari/5351": 1,
					"Mozilla/5.0 (Windows NT 4.0) AppleWebKit/5340 (KHTML, like Gecko) Chrome/39.0.841.0 Mobile Safari/5340": 1,
					"Mozilla/5.0 (Windows NT 6.2) AppleWebKit/5321 (KHTML, like Gecko) Chrome/40.0.855.0 Mobile Safari/5321": 1,
				},
				1474.9,
				2733.5,
			),
		},
		{
			name: "local path with filter",
			args: args{
				read:              10,
				paths:             localPath2,
				field:             "http_user_agent",
				value:             "Opera*",
				isLocal:           isLocal2,
				isFilterSpecified: true,
			},
			wantRep: report.New(
				localPath2,
				"-",
				"-",
				"http_user_agent",
				"Opera*",
				10,
				map[string]int{
					"/Multi-layered/responsive/disintermediate/task-force-Triple-buffered.gif":   1,
					"/database/Visionary/Universal/analyzing_regional.jpg":                       1,
					"/alliance-Vision-oriented_3rd%20generation%20hardware-Multi-channelled.css": 1,
					"/Face%20to%20face%20heuristic/monitoring/Versatile-well-modulated.htm":      1,
					"/tangible-data-warehouse%20Stand-alone.css":                                 1,
					"/static/focus%20group_open%20architecture-Switchable/implementation.gif":    1,
					"/human-resource_installation/matrix/high-level/Balanced.png":                1,
					"/exuding.css": 1,
					"/encompassing-Quality-focused-Switchable/archive%20Exclusive.htm":  1,
					"/demand-driven_multi-state%20User-friendly_workforce_Extended.svg": 1,
				},
				map[int]int{
					200: 10,
				},
				map[string]int{
					"54.183.46.90":    1,
					"33.12.134.41":    1,
					"199.211.174.40":  1,
					"89.196.204.84":   1,
					"19.226.223.159":  1,
					"172.241.115.5":   1,
					"167.135.102.87":  1,
					"147.230.144.173": 1,
					"138.118.139.56":  1,
					"111.85.221.83":   1,
				},
				map[string]int{
					"Opera/9.81 (X11; Linux i686; en-US) Presto/2.11.163 Version/13.00":                     1,
					"Opera/9.54 (Macintosh; Intel Mac OS X 10_7_0; en-US) Presto/2.9.247 Version/13.00":     1,
					"Opera/8.12 (X11; Linux x86_64; en-US) Presto/2.10.340 Version/13.00":                   1,
					"Opera/10.88 (Macintosh; PPC Mac OS X 10_6_9; en-US) Presto/2.10.337 Version/13.00":     1,
					"Opera/10.19 (Macintosh; Intel Mac OS X 10_7_8; en-US) Presto/2.9.282 Version/12.00":    1,
					"Opera/9.43 (Windows NT 5.2; en-US) Presto/2.9.260 Version/13.00":                       1,
					"Opera/9.38 (Windows NT 6.1; en-US) Presto/2.9.256 Version/12.00":                       1,
					"Opera/9.26 (Windows NT 4.0; en-US) Presto/2.8.253 Version/11.00":                       1,
					"Opera/9.11 (Macintosh; U; Intel Mac OS X 10_7_4; en-US) Presto/2.11.176 Version/10.00": 1,
					"Opera/8.74 (X11; Linux i686; en-US) Presto/2.8.226 Version/12.00":                      1,
				},
				1696.8,
				2448.5,
			),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := analyzer.New(&parser.Parser{})

			gotRep, err := a.Analyze(
				tt.args.from,
				tt.args.to,
				tt.args.field,
				tt.args.value,
				tt.args.read,
				tt.args.isFromSpecified,
				tt.args.isToSpecified,
				tt.args.isFilterSpecified,
				tt.args.paths,
				tt.args.isLocal)

			assert.Equal(t, tt.wantRep, gotRep)
			assert.NoError(t, err)
		})
	}
}

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
