package services

import (
	"time"
	"fmt"
	"strings"
)

type Transformer struct{}

var layoutInput = "02/01/2006"
var layoutOutput = "2006-01-02"

//Todo: This logic is coupled to csv file. Make it generic
func (t *Transformer) Transform(data [][]string) ([][]string, error) {
	//header := data[0]
	data = append(data[:0], data[0+1:]...)

	for i := range data {

		transactionDate, err := t.transformDate(data[i][0])
		if err != nil {
			return nil, err
		}

		data[i][0] = transactionDate

		transactionDescription := t.cleanString(data[i][4])
		data[i][4] = transactionDescription
	}

	return data, nil
}

func (t *Transformer) transformDate(date string) (string, error) {
	stringOutput, err := time.Parse(layoutInput, date)

	if err != nil {
		fmt.Printf("Error parsing date: %s", err.Error())
		return "", err
	}
	return stringOutput.Format(layoutOutput), nil
}

func (t *Transformer) cleanString(value string) string {
	valueLower := strings.ToLower(value)
	return t.standardizeSpaces(valueLower)
}

func (t *Transformer) standardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
