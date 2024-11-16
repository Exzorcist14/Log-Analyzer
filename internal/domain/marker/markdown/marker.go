package markdown

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/marker/mutils"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/report"
)

const separator = "<br>"

type Marker struct{}

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
	markUpTitle(builder, mutils.TitleGeneralInfo)
	markUpTableHeader(builder, mutils.Header1GeneralInfo, mutils.Header2GeneralInfo)
	markUpTableRow(builder, mutils.Row1GeneralInfo, mutils.GetTableCellWithMultipleValues(rep.Files, separator))
	markUpTableRow(builder, mutils.Row2GeneralInfo, rep.From)
	markUpTableRow(builder, mutils.Row3GeneralInfo, rep.To)
	markUpTableRow(builder, mutils.Row4GeneralInfo, strconv.Itoa(rep.RequestsCount))
	markUpTableRow(builder, mutils.Row5GeneralInfo, strconv.FormatFloat(rep.AverageResponseSize,
		mutils.FloatFormat, mutils.Prec, mutils.BitSize))
	markUpTableRow(builder, mutils.Row6GeneralInfo, strconv.FormatFloat(rep.Percentile95ResponseSize,
		mutils.FloatFormat, mutils.Prec, mutils.BitSize))
}

func markUpResources(builder *strings.Builder, rep *report.Report, highest int) {
	markUpTitle(builder, mutils.TitleResources)
	markUpTableHeader(builder, mutils.Header1Resources, mutils.Header2Resources)

	for i := 0; i < len(rep.MostFrequentResources) && i < highest; i++ {
		markUpTableRow(builder, rep.MostFrequentResources[i].Data, strconv.Itoa(rep.MostFrequentResources[i].Count))
	}
}

func markUpCodes(builder *strings.Builder, rep *report.Report, highest int) {
	markUpTitle(builder, mutils.TitleCodes)
	markUpTableHeader(builder, mutils.Header1Codes, mutils.Header2Codes, mutils.Header3Codes)

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
	markUpTitle(builder, mutils.TitleClients)
	markUpTableHeader(builder, mutils.Header1Clients, mutils.Header2Clients)

	for i := 0; i < len(rep.MostFrequentClients) && i < highest; i++ {
		markUpTableRow(builder, rep.MostFrequentClients[i].Data, strconv.Itoa(rep.MostFrequentClients[i].Count))
	}
}

func markUpAgents(builder *strings.Builder, rep *report.Report, highest int) {
	markUpTitle(builder, mutils.TitleAgents)
	markUpTableHeader(builder, mutils.Header1Agents, mutils.Header2Agents)

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
