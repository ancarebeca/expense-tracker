package services

import (
	"encoding/csv"
	"fmt"
	"github.com/ancarebeca/expense-tracker/config"
	"os"
)

type CsvReader struct {
	Conf config.Conf
}

//go:generate counterfeiter . Reader
type Reader interface {
	ReadCsv() ([][]string, error)
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
