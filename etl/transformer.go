package etl

import (
	"strings"
	"time"

	"github.com/ancarebeca/expense-tracker/model"
	"github.com/sirupsen/logrus"
)

type DataTransformer struct{}

var layoutInput = "02/01/2006"
var layoutOutput = "2006-01-02"

type Transformer interface {
	Transform(statements []model.Statement) []model.Statement
}

func (t *DataTransformer) Transform(statements []model.Statement) []model.Statement {
	data := []model.Statement{}

	for _, s := range statements {
		date, err := t.transformDate(s.TransactionDate)
		if err != nil {
			continue
		}
		s.TransactionDate = date
		s.TransactionDescription = t.cleanString(s.TransactionDescription)
		data = append(data, s)
	}

	return data
}

func (t *DataTransformer) transformDate(date string) (string, error) {
	stringOutput, err := time.Parse(layoutInput, date)
	if err != nil {
		logrus.Errorf("Error parsing transaction date: %s", err.Error())
		return "", err
	}
	return stringOutput.Format(layoutOutput), nil
}

func (t *DataTransformer) cleanString(value string) string {
	valueLower := strings.ToLower(value)
	return t.standardizeSpaces(valueLower)
}

func (t *DataTransformer) standardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
