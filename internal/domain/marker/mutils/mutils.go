package mutils

import "strings"

const (
	TitleGeneralInfo   = "Общая информация"
	TitleResources     = "Запрашиваемые ресурсы"
	TitleCodes         = "Коды ответа"
	TitleClients       = "IP-адреса клиентов"
	TitleAgents        = "HTTP-заголовки User-Agent"
	Header1GeneralInfo = "Метрика"
	Header2GeneralInfo = "Значение"
	Row1GeneralInfo    = "Файл(-ы)"
	Row2GeneralInfo    = "Начальная дата"
	Row3GeneralInfo    = "Конечная дата"
	Row4GeneralInfo    = "Количество запросов"
	Row5GeneralInfo    = "Средний размер ответа"
	Row6GeneralInfo    = "95p размера ответа"
	Header1Resources   = "Ресурс"
	Header2Resources   = "Количество"
	Header1Codes       = "Код"
	Header2Codes       = "Имя"
	Header3Codes       = "Количество"
	Header1Clients     = "Клиент"
	Header2Clients     = "Количество"
	Header1Agents      = "Агент"
	Header2Agents      = "Количество"
	FloatFormat        = 'f'
	Prec               = -1
	BitSize            = 64
)

func GetTableCellWithMultipleValues(cell []string, separator string) string {
	var builder strings.Builder

	for _, data := range cell {
		builder.WriteString(data)
		builder.WriteString(separator)
	}

	return builder.String()
}
