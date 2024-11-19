package adoc_test

import (
	"testing"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/marker/adoc"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/report"
	"github.com/stretchr/testify/assert"
)

func TestMarkUp(t *testing.T) {
	type args struct {
		rep     report.Report
		highest int
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "checking the format and content of a marked-up report",
			args: args{
				rep: report.New(
					[]string{
						`C:\Users\vova_\GolandProjects\backend_academy_2024_project_3-go-Exzorcist14\internal\infrastructure\logs\2024-11-08\logs.txt`,
					},
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
						"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_6_7) AppleWebKit/5322 (KHTML, like Gecko) Chrome/38.0.821.0 Mobile Safari/5322":  1,
						"Opera/10.13 (Macintosh; PPC Mac OS X 10_9_3; en-US) Presto/2.10.206 Version/12.00":                                         1,
						"Opera/10.78 (Macintosh; U; Intel Mac OS X 10_8_8; en-US) Presto/2.12.254 Version/12.00":                                    1,
						"Opera/10.10 (X11; Linux x86_64; en-US) Presto/2.10.230 Version/13.00":                                                      1,
						"Mozilla/5.0 (Windows NT 6.2; en-US; rv:1.9.0.20) Gecko/1976-07-11 Firefox/35.0":                                            1,
						"Mozilla/5.0 (Windows; U; Windows CE) AppleWebKit/532.41.4 (KHTML, like Gecko) Version/5.1 Safari/532.41.4":                 1,
						"Mozilla/5.0 (Macintosh; PPC Mac OS X 10_9_1 rv:7.0) Gecko/2024-10-07 Firefox/37.0":                                         1,
						"Mozilla/5.0 (Macintosh; U; PPC Mac OS X 10_7_8) AppleWebKit/5351 (KHTML, like Gecko) Chrome/36.0.834.0 Mobile Safari/5351": 1,
						"Mozilla/5.0 (Windows NT 4.0) AppleWebKit/5340 (KHTML, like Gecko) Chrome/39.0.841.0 Mobile Safari/5340":                    1,
						"Mozilla/5.0 (Windows NT 6.2) AppleWebKit/5321 (KHTML, like Gecko) Chrome/40.0.855.0 Mobile Safari/5321":                    1,
					},
					1474.9,
					2733.5,
				),
				highest: 10,
			},
			want: "== Общая информация\n" +
				"[cols=\"^,^\", options=\"header\"]\n" +
				"|===\n" +
				"|Метрика|Значение\n" +
				"\n" +
				"|Файл(-ы)|C:\\Users\\vova_\\GolandProjects\\backend_academy_2024_project_3-go-Exzorcist14\\internal\\" +
				"infrastructure\\logs\\2024-11-08\\logs.txt +\n" +
				"\n" +
				"|Начальная дата|2024-11-08 14:39:44 +0000 +0000\n" +
				"|Конечная дата|2024-11-08 14:40:03 +0000 +0000\n" +
				"|Фильтр|-\n" +
				"|Значение фильтра|-\n" +
				"|Количество запросов|10\n" +
				"|Средний размер ответа|1474.9\n" +
				"|95p размера ответа|2733.5\n" +
				"|===\n" +
				"== Запрашиваемые ресурсы\n" +
				"[cols=\"^,^\", options=\"header\"]\n" +
				"|===\n" +
				"|Ресурс|Количество\n" +
				"\n" +
				"|/6th%20generation.php|1\n" +
				"|/Synchronised-secured%20line.svg|1\n" +
				"|/approach.hmtl|1\n" +
				"|/approach/Adaptive%20infrastructure.gif|1\n" +
				"|/database/object-oriented%20neural-net%20Virtual%20Universal.htm|1\n" +
				"|/dedicated/help-desk%20Exclusive/didactic%20structure.php|1\n" +
				"|/encryption/Virtual.png|1\n" +
				"|/interactive.css|1\n" +
				"|/migration/human-resource%20Organized/Stand-alone.png|1\n" +
				"|/reciprocal/Synergistic/transitional.hmtl|1\n" +
				"|===\n" +
				"== Коды ответа\n" +
				"[cols=\"^,^,^\", options=\"header\"]\n" +
				"|===\n" +
				"|Код|Имя|Количество\n" +
				"\n" +
				"|200|OK|7\n" +
				"|404|Not Found|2\n" +
				"|400|Bad Request|1\n" +
				"|===\n" +
				"== IP-адреса клиентов\n" +
				"[cols=\"^,^\", options=\"header\"]\n" +
				"|===\n" +
				"|Клиент|Количество\n" +
				"\n" +
				"|107.1.82.37|1\n" +
				"|147.145.53.17|1\n" +
				"|161.123.208.218|1\n" +
				"|169.145.174.150|1\n" +
				"|196.101.108.15|1\n" +
				"|21.45.234.173|1\n" +
				"|215.215.249.196|1\n" +
				"|232.22.135.239|1\n" +
				"|76.219.118.160|1\n" +
				"|93.46.164.250|1\n" +
				"|===\n" +
				"== HTTP-заголовки User-Agent\n" +
				"[cols=\"^,^\", options=\"header\"]\n" +
				"|===\n" +
				"|Агент|Количество\n" +
				"\n" +
				"|Mozilla/5.0 (Macintosh; Intel Mac OS X 10_6_7) AppleWebKit/5322 (KHTML, like Gecko) Chrome/38.0.821.0 Mobile Safari/5322|1\n" +
				"|Mozilla/5.0 (Macintosh; PPC Mac OS X 10_9_1 rv:7.0) Gecko/2024-10-07 Firefox/37.0|1\n" +
				"|Mozilla/5.0 (Macintosh; U; PPC Mac OS X 10_7_8) AppleWebKit/5351 (KHTML, like Gecko) Chrome/36.0.834.0 Mobile Safari/5351|1\n" +
				"|Mozilla/5.0 (Windows NT 4.0) AppleWebKit/5340 (KHTML, like Gecko) Chrome/39.0.841.0 Mobile Safari/5340|1\n" +
				"|Mozilla/5.0 (Windows NT 6.2) AppleWebKit/5321 (KHTML, like Gecko) Chrome/40.0.855.0 Mobile Safari/5321|1\n" +
				"|Mozilla/5.0 (Windows NT 6.2; en-US; rv:1.9.0.20) Gecko/1976-07-11 Firefox/35.0|1\n" +
				"|Mozilla/5.0 (Windows; U; Windows CE) AppleWebKit/532.41.4 (KHTML, like Gecko) Version/5.1 Safari/532.41.4|1\n" +
				"|Opera/10.10 (X11; Linux x86_64; en-US) Presto/2.10.230 Version/13.00|1\n" +
				"|Opera/10.13 (Macintosh; PPC Mac OS X 10_9_3; en-US) Presto/2.10.206 Version/12.00|1\n" +
				"|Opera/10.78 (Macintosh; U; Intel Mac OS X 10_8_8; en-US) Presto/2.12.254 Version/12.00|1\n" +
				"|===\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := (&adoc.Marker{}).MarkUp(&tt.args.rep, tt.args.highest)

			assert.Equal(t, tt.want, got)
		})
	}
}
