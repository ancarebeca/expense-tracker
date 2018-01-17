package etl

import (
	"fmt"
	"strconv"

	"github.com/ancarebeca/expense-tracker/model"
	"github.com/sirupsen/logrus"
)

const (
	debitAmountIndex            = 5
	creditAmountIndex           = 6
	balanceAmountIndex          = 7
	transactionDateIndex        = 0
	transactionTypeIndex        = 1
	TransactionDescriptionIndex = 4
)

type SantanderParser struct{}

type Parser interface {
	Parse(data [][]string) []model.Statement
}

func (p *SantanderParser) Parse(data [][]string) []model.Statement {
	statements := []model.Statement{}
	data = removeHeader(data)

	for i := range data {
		debit, err := p.getDebitAmount(data[i])
		if err != nil {
			continue
		}
		credit, err := p.getCreditAmount(data[i])
		if err != nil {
			continue
		}
		balance, err := p.getBalanceAmount(data[i])
		if err != nil {
			continue
		}
		s := model.Statement{
			TransactionDate:        data[i][transactionDateIndex],
			TransactionType:        data[i][transactionTypeIndex],
			TransactionDescription: data[i][TransactionDescriptionIndex],
			DebitAmount:            debit,
			CreditAmount:           credit,
			Balance:                balance,
		}
		statements = append(statements, s)
	}

	return statements
}

func removeHeader(data [][]string) [][]string {
	data = append(data[:0], data[0+1:]...)
	return data
}

func (p *SantanderParser) getDebitAmount(data []string) (float64, error) {
	debit, err := p.stringToFloat(data[debitAmountIndex])
	if err != nil {
		logrus.Errorf("Error in SantanderParser while casting debit amount to string. Data = %s ", data)
		return 0, err
	}
	return debit, nil
}

func (p *SantanderParser) getCreditAmount(data []string) (float64, error) {
	credit, err := p.stringToFloat(data[creditAmountIndex])
	if err != nil {
		logrus.Errorf("Error in SantanderParser while casting credit amount to string. Data = %s ", data)
		return 0, err
	}
	return credit, nil
}

func (p *SantanderParser) getBalanceAmount(data []string) (float64, error) {
	balance, err := p.stringToFloat(data[balanceAmountIndex])
	if err != nil {
		logrus.Errorf("Error in SantanderParser while casting balance amount (%s) to string. Data = %s ", data[balanceAmountIndex], data)
		return 0, err
	}
	return balance, nil
}

func (p *SantanderParser) stringToFloat(value string) (float64, error) {
	if value == "" {
		value = "0"
	}

	valueFloat, err := strconv.ParseFloat(value, 64)
	if err != nil {
		logrus.Debugf(fmt.Sprintf("Error in SantanderParser while casting value = '%s' into float. %s .", value, err.Error))
		return 0, err
	}

	return valueFloat, nil
}
