package markdown

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/report"
)

const (
	titleGeneralInfo   = "Общая информация"
	titleResources     = "Запрашиваемые ресурсы"
	titleCodes         = "Коды ответа"
	titleClients       = "IP-адреса клиентов"
	titleAgents        = "HTTP-заголовки User-Agent"
	header1GeneralInfo = "Метрика"
	header2GeneralInfo = "Значение"
	row1GeneralInfo    = "Файл(-ы)"
	row2GeneralInfo    = "Начальная дата"
	row3GeneralInfo    = "Конечная дата"
	row4GeneralInfo    = "Количество запросов"
	row5GeneralInfo    = "Средний размер ответа"
	row6GeneralInfo    = "95p размера ответа"
	header1Resources   = "Ресурс"
	header2Resources   = "Количество"
	header1Codes       = "Код"
	header2Codes       = "Имя"
	header3Codes       = "Количество"
	header1Clients     = "Клиент"
	header2Clients     = "Количество"
	header1Agents      = "Агент"
	header2Agents      = "Количество"
)

type Marker struct{}

func New() *Marker {
	return &Marker{}
}

func (p *Marker) MarkUp(rep *report.Report, highest int) string {
	var builder strings.Builder

	markUpGeneralInfo(&builder, rep)
	markUpResources(&builder, rep, highest)
	markUpCodes(&builder, rep, highest)
	markUpClients(&builder, rep, highest)
	markUpAgents(&builder, rep, highest)

	return builder.String()
}

func markUpGeneralInfo(builder *strings.Builder, rep *report.Report) {
	markUpTitle(builder, titleGeneralInfo)
	markUpTableHeader(builder, header1GeneralInfo, header2GeneralInfo)
	markUpTableRow(builder, row1GeneralInfo, getCellWithMultipleValues(rep.Files))
	markUpTableRow(builder, row2GeneralInfo, rep.From)
	markUpTableRow(builder, row3GeneralInfo, rep.To)
	markUpTableRow(builder, row4GeneralInfo, strconv.Itoa(rep.RequestsCount))
	markUpTableRow(builder, row5GeneralInfo, strconv.FormatFloat(rep.AverageResponseSize, 'f', -1, 64))
	markUpTableRow(builder, row6GeneralInfo, strconv.FormatFloat(rep.Percentile95ResponseSize, 'f', -1, 64))
}

func markUpResources(builder *strings.Builder, rep *report.Report, highest int) {
	markUpTitle(builder, titleResources)
	markUpTableHeader(builder, header1Resources, header2Resources)

	for i := 0; i < len(rep.MostFrequentResources) && i < highest; i++ {
		markUpTableRow(builder, rep.MostFrequentResources[i].Data, strconv.Itoa(rep.MostFrequentResources[i].Count))
	}
}

func markUpCodes(builder *strings.Builder, rep *report.Report, highest int) {
	markUpTitle(builder, titleCodes)
	markUpTableHeader(builder, header1Codes, header2Codes, header3Codes)

	for i := 0; i < len(rep.MostFrequentCodes) && i < highest; i++ {
		markUpTableRow(
			builder,
			strconv.Itoa(rep.MostFrequentCodes[i].Data),
			http.StatusText(rep.MostFrequentCodes[i].Data),
			strconv.Itoa(rep.MostFrequentCodes[i].Count),
		)
	}
}

func markUpClients(builder *strings.Builder, rep *report.Report, highest int) {
	markUpTitle(builder, titleClients)
	markUpTableHeader(builder, header1Clients, header2Clients)

	for i := 0; i < len(rep.MostFrequentClients) && i < highest; i++ {
		markUpTableRow(builder, rep.MostFrequentClients[i].Data, strconv.Itoa(rep.MostFrequentClients[i].Count))
	}
}

func markUpAgents(builder *strings.Builder, rep *report.Report, highest int) {
	markUpTitle(builder, titleAgents)
	markUpTableHeader(builder, header1Agents, header2Agents)

	for i := 0; i < len(rep.MostFrequentAgents) && i < highest; i++ {
		markUpTableRow(builder, rep.MostFrequentAgents[i].Data, strconv.Itoa(rep.MostFrequentAgents[i].Count))
	}
}

func markUpTitle(builder *strings.Builder, name string) {
	fmt.Fprintf(builder, "## %s\n", name)
}

func markUpTableHeader(builder *strings.Builder, headers ...string) {
	markUpTableRow(builder, headers...)

	for range headers {
		builder.WriteString("|:-:")
	}

	builder.WriteString("|\n")
}

func markUpTableRow(builder *strings.Builder, cells ...string) {
	for _, cell := range cells {
		builder.WriteString("|")
		builder.WriteString(cell)
	}

	builder.WriteString("|\n")
}

func getCellWithMultipleValues(cell []string) string {
	var builder strings.Builder

	for _, data := range cell {
		builder.WriteString(data)
		builder.WriteString("<br>")
	}

	return builder.String()
}
