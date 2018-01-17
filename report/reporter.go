package report

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type Report struct {
	OutputPath string
	Reporter
}

func (e *Report) Create() error {

	report, _ := e.Reporter.Create()

	file, err := os.Create(e.OutputPath)
	if err != nil {
		fmt.Printf("Cannot create file ", err.Error())
		return err
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for i := range report {
		err = writer.Write(report[i])
		if err != nil {
			log.Fatal(err)
			return err
		}
	}

	return nil
}
