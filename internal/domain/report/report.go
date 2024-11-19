package report

import (
	"sort"
)

// DataWithCount хранит пару значений Data и Count.
// Data может быть типом string или int.
type DataWithCount[T string | int] struct {
	Data  T
	Count int
}

// Report - структура отчёта, содержащая результаты анализа и метаинформацию о нём.
type Report struct {
	Files                    []string
	From                     string
	To                       string
	Field                    string
	Value                    string
	RequestsCount            int
	MostFrequentResources    []DataWithCount[string]
	MostFrequentCodes        []DataWithCount[int]
	MostFrequentClients      []DataWithCount[string]
	MostFrequentAgents       []DataWithCount[string]
	AverageResponseSize      float64
	Percentile95ResponseSize float64
}

// New возвращает инициализированный Report.
func New(
	files []string,
	from, to, field, value string,
	requestCount int,
	resources map[string]int,
	codes map[int]int,
	clients map[string]int,
	agents map[string]int,
	averageServerResponseSize, serverResponseSize95Percentile float64,
) Report {
	rs := transformMapToSlice(resources)
	cd := transformMapToSlice(codes)
	cl := transformMapToSlice(clients)
	ag := transformMapToSlice(agents)

	sort.Slice(rs, func(i, j int) bool {
		if rs[i].Count == rs[j].Count {
			return rs[i].Data < rs[j].Data
		}

		return rs[i].Count > rs[j].Count
	})

	sort.Slice(cd, func(i, j int) bool {
		if cd[i].Count == cd[j].Count {
			return cd[i].Data < cd[j].Data
		}

		return cd[i].Count > cd[j].Count
	})

	sort.Slice(cl, func(i, j int) bool {
		if cl[i].Count == cl[j].Count {
			return cl[i].Data < cl[j].Data
		}

		return cl[i].Count > cl[j].Count
	})

	sort.Slice(ag, func(i, j int) bool {
		if ag[i].Count == ag[j].Count {
			return ag[i].Data < ag[j].Data
		}

		return ag[i].Count > ag[j].Count
	})

	return Report{
		Files:                    files,
		From:                     from,
		To:                       to,
		Field:                    field,
		Value:                    value,
		RequestsCount:            requestCount,
		MostFrequentResources:    rs,
		MostFrequentCodes:        cd,
		MostFrequentClients:      cl,
		MostFrequentAgents:       ag,
		AverageResponseSize:      averageServerResponseSize,
		Percentile95ResponseSize: serverResponseSize95Percentile,
	}
}

// transformMapToSlice преобразует словарь {данные - значение} в слайс DataWithCount.
// Данные могут быть представлены типами string или int.
func transformMapToSlice[T string | int](mp map[T]int) []DataWithCount[T] {
	slice := []DataWithCount[T]{}

	for data, count := range mp {
		slice = append(slice, DataWithCount[T]{
			Data:  data,
			Count: count,
		})
	}

	return slice
}
