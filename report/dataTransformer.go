package report

import (
	"sort"
)

type YearlySpendingByCategoryReport struct {
	Calculator Calculator
	Adapter    DataAdapter
}

type Reporter interface {
	Create() ([][]string, error)
}

type DataAdapter struct {
	Values map[string]map[string]string
	Years  []string
}

func (ysc YearlySpendingByCategoryReport) Create() ([][]string, error) {
	data, _ := ysc.Calculator.CalculateExpenses()
	reportData := ysc.Adapter.TransformData(data)

	var output [][]string
	header := ysc.createrHeader(reportData)
	output = append(output, header)
	output = ysc.createOutput(reportData, output)

	return output, nil
}

func (ysc YearlySpendingByCategoryReport) createOutput(data DataAdapter, output [][]string) [][]string {

	for key, expense := range data.Values {
		var record []string
		record = append(record, key)

		for _, v := range data.Years {
			if _, ok := expense[v]; !ok {
				record = append(record, "0")
			} else {
				record = append(record, expense[v])
			}
		}

		output = append(output, record)
	}
	return output
}

func (ysc YearlySpendingByCategoryReport) createrHeader(a DataAdapter) []string {
	var header []string

	header = append(header, "Category")

	for _, v := range a.Years {
		header = append(header, v)
	}
	return header
}

func (a DataAdapter) TransformData(report []ReportModel) DataAdapter {
	Values := make(map[string]map[string]string)
	Years := make(map[string]bool)

	v := make(map[string]string)

	for _, report := range report {

		if _, ok := Values[report.Category]; !ok {
			v = make(map[string]string)
		}
		v[report.Year] = report.Amount
		Values[report.Category] = v
		Years[report.Year] = true
	}
	a.Values = Values

	years := make([]string, 0)
	for k := range Years {
		years = append(years, k)
	}
	sort.Strings(years)
	a.Years = years
	return a
}
