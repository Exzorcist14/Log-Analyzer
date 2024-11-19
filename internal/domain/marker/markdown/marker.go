package markdown

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/marker/mutils"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/report"
)

const separator = "<br>" // Разделитель строк внутри ячейки таблицы.

// Marker умеет размечать отчёт в соответствии с markdown.
type Marker struct{}

// MarkUp размечает отчёт, используя markdown, записывая первые highest значений таблиц, не содержащих общую информацию.
func (p *Marker) MarkUp(rep *report.Report, highest int) string {
	var builder strings.Builder

	markUpGeneralInfo(&builder, rep)
	markUpResources(&builder, rep, highest)
	markUpCodes(&builder, rep, highest)
	markUpClients(&builder, rep, highest)
	markUpAgents(&builder, rep, highest)

	return builder.String()
}

// markUpGeneralInfo размечает заголовок и таблицу общей информации.
func markUpGeneralInfo(builder *strings.Builder, rep *report.Report) {
	markUpTitle(builder, mutils.TitleGeneralInfo)
	markUpTableHeader(builder, mutils.Header1GeneralInfo, mutils.Header2GeneralInfo)
	markUpTableRow(builder, mutils.Row1GeneralInfo, mutils.GetTableCellWithMultipleValues(rep.Files, separator))
	markUpTableRow(builder, mutils.Row2GeneralInfo, rep.From)
	markUpTableRow(builder, mutils.Row3GeneralInfo, rep.To)
	markUpTableRow(builder, mutils.Row4GeneralInfo, rep.Field)
	markUpTableRow(builder, mutils.Row5GeneralInfo, rep.Value)
	markUpTableRow(builder, mutils.Row6GeneralInfo, strconv.Itoa(rep.RequestsCount))
	markUpTableRow(builder, mutils.Row7GeneralInfo, strconv.FormatFloat(rep.AverageResponseSize,
		mutils.FloatFormat, mutils.Prec, mutils.BitSize))
	markUpTableRow(builder, mutils.Row8GeneralInfo, strconv.FormatFloat(rep.Percentile95ResponseSize,
		mutils.FloatFormat, mutils.Prec, mutils.BitSize))
}

// markUpResources размечает заголовок и таблицу заправшиваемых ресурсов.
func markUpResources(builder *strings.Builder, rep *report.Report, highest int) {
	markUpTitle(builder, mutils.TitleResources)
	markUpTableHeader(builder, mutils.Header1Resources, mutils.Header2Resources)

	// Размечаются первые highest значений, или все, если highest больше их количества.
	for i := 0; i < len(rep.MostFrequentResources) && i < highest; i++ {
		markUpTableRow(builder, rep.MostFrequentResources[i].Data, strconv.Itoa(rep.MostFrequentResources[i].Count))
	}
}

// markUpCodes размечает заголовок и таблицу кодов ответа.
func markUpCodes(builder *strings.Builder, rep *report.Report, highest int) {
	markUpTitle(builder, mutils.TitleCodes)
	markUpTableHeader(builder, mutils.Header1Codes, mutils.Header2Codes, mutils.Header3Codes)

	// Размечаются первые highest значений, или все, если highest больше их количества.
	for i := 0; i < len(rep.MostFrequentCodes) && i < highest; i++ {
		markUpTableRow(
			builder,
			strconv.Itoa(rep.MostFrequentCodes[i].Data),
			http.StatusText(rep.MostFrequentCodes[i].Data),
			strconv.Itoa(rep.MostFrequentCodes[i].Count),
		)
	}
}

// markUpClients размечает заголовок и таблицу ip-адресов клиентов.
func markUpClients(builder *strings.Builder, rep *report.Report, highest int) {
	markUpTitle(builder, mutils.TitleClients)
	markUpTableHeader(builder, mutils.Header1Clients, mutils.Header2Clients)

	// Размечаются первые highest значений, или все, если highest больше их количества.
	for i := 0; i < len(rep.MostFrequentClients) && i < highest; i++ {
		markUpTableRow(builder, rep.MostFrequentClients[i].Data, strconv.Itoa(rep.MostFrequentClients[i].Count))
	}
}

// markUpAgents размечает заголовок и таблицу HTTP-заголовков User-Agent.
func markUpAgents(builder *strings.Builder, rep *report.Report, highest int) {
	markUpTitle(builder, mutils.TitleAgents)
	markUpTableHeader(builder, mutils.Header1Agents, mutils.Header2Agents)

	// Размечаются первые highest значений, или все, если highest больше их количества.
	for i := 0; i < len(rep.MostFrequentAgents) && i < highest; i++ {
		markUpTableRow(builder, rep.MostFrequentAgents[i].Data, strconv.Itoa(rep.MostFrequentAgents[i].Count))
	}
}

// markUpTitle размечает заголовок второго уровня в markdown.
func markUpTitle(builder *strings.Builder, name string) {
	fmt.Fprintf(builder, "## %s\n", name)
}

// markUpTableHeader размечает начало таблицы в markdown.
func markUpTableHeader(builder *strings.Builder, headers ...string) {
	markUpTableRow(builder, headers...)

	for range headers {
		builder.WriteString("|:-:")
	}

	builder.WriteString("|\n")
}

// markUpTableHeader размечает строку таблицы в markdown.
func markUpTableRow(builder *strings.Builder, cells ...string) {
	for _, cell := range cells {
		builder.WriteString("|")
		builder.WriteString(cell)
	}

	builder.WriteString("|\n")
}
