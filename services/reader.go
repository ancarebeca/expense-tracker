package services

import (
	"os"
	"fmt"
	"github.com/ancarebeca/expense-tracker/config"
	"encoding/csv"
)

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
