package services

import (
	"strconv"
	"github.com/ancarebeca/expense-tracker/model"
	"fmt"
)

type DataParser struct {
}

//go:generate counterfeiter . Parser
type Parser interface {
	Parse(data [][]string) ([]*model.Statement, error)
}

//Todo: This logic is coupled to csv file. Make it generic
func (p *DataParser) Parse(data [][]string) ([]*model.Statement, error) {
	statements := []*model.Statement{}

	for i := range data {
		debitAmount, err := p.stringToFloat(data[i][5])
		if err != nil {
			return nil, err
		}

		creditAmount, err := p.stringToFloat(data[i][6])
		if err != nil {
			return nil, err
		}

		balanceAmount, err := p.stringToFloat(data[i][7])
		if err != nil {
			return nil, err
		}

		s := &model.Statement{
			TransactionDate:        data[i][0],
			TransactionType:        data[i][1],
			TransactionDescription: data[i][4],
			DebitAmount:            debitAmount,
			CreditAmount:           creditAmount,
			Balance:                balanceAmount,
		}
		statements = append(statements, s)
	}

	return statements, nil
}

func (p *DataParser) stringToFloat(value string) (float64, error) {
	if value == "" {
		value = "0"
	}

	valueFloat, err := strconv.ParseFloat(value, 64)
	if err != nil {
		fmt.Printf("Error while casting value %s into float: %s", value, err.Error())
		return 0, err
	}

	return valueFloat, nil
}
