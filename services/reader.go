package services

import (
	"encoding/csv"
	"fmt"
	"os"
)

type CsvReader struct{}

//go:generate counterfeiter . Reader
type Reader interface {
	ReadCsv(f string) ([][]string, error)
}

func (c *CsvReader) ReadCsv(file string) ([][]string, error) {
	f, err := os.Open(file)

	if err != nil {
		fmt.Printf("Error while opening CSV file: %s", err.Error())
		return nil, err
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		fmt.Printf("Error while reading the CSV: %s", err.Error())
		return nil, err
	}

	return lines, nil
}
