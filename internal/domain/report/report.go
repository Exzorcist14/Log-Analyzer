package report

import (
	"sort"
)

type DataWithCount[T string | int] struct {
	Data  T
	Count int
}

type Report struct {
	Files                    []string
	From                     string
	To                       string
	RequestsCount            int
	MostFrequentResources    []DataWithCount[string]
	MostFrequentCodes        []DataWithCount[int]
	MostFrequentClients      []DataWithCount[string]
	MostFrequentAgents       []DataWithCount[string]
	AverageResponseSize      float64
	Percentile95ResponseSize float64
}

func New(
	files []string,
	from, to string,
	requestCount int,
	resources []DataWithCount[string],
	codes []DataWithCount[int],
	clients []DataWithCount[string],
	agents []DataWithCount[string],
	averageServerResponseSize, serverResponseSize95Percentile float64,
) Report {
	sort.Slice(resources, func(i, j int) bool {
		return resources[i].Count > resources[j].Count
	})

	sort.Slice(codes, func(i, j int) bool {
		return codes[i].Count > codes[j].Count
	})

	sort.Slice(clients, func(i, j int) bool {
		return clients[i].Count > clients[j].Count
	})

	sort.Slice(agents, func(i, j int) bool {
		return agents[i].Count > agents[j].Count
	})

	return Report{
		Files:                    files,
		From:                     from,
		To:                       to,
		RequestsCount:            requestCount,
		MostFrequentResources:    resources,
		MostFrequentCodes:        codes,
		MostFrequentClients:      clients,
		MostFrequentAgents:       agents,
		AverageResponseSize:      averageServerResponseSize,
		Percentile95ResponseSize: serverResponseSize95Percentile,
	}
}
