package services

import (
	"os"
	"fmt"
	"github.com/ancarebeca/expense-tracker/config"
	"encoding/csv"
)

type Statement struct {
	TransactionDate        string  `csv:"Transaction Date"`
	TransactionType        string  `csv:"Transaction Type"`
	TransactionDescription string  `csv:"Transaction Description"`
	Category               string  `csv:"Category"`
	DebitAmount            float64 `csv:"Debit Amount"`
	CreditAmount           float64 `csv:"Credit Amount"`
	Balance                float64 `csv:"Balance"`
}

type CsvReader struct {
	Conf config.Conf
}

func (c *CsvReader) ReadCsv() ([][]string, error) {
	f, err := os.Open(c.Conf.FilePath)

	if err != nil {
		fmt.Printf("Error while opening configuration: %s", err.Error())
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		fmt.Printf("Error while reading the CSV: %s", err.Error())
		return nil, err
	}

	return lines, nil
}
