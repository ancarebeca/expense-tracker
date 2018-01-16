package services_test

import (
	"github.com/ancarebeca/expense-tracker/model"
	"github.com/ancarebeca/expense-tracker/services"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_categorise_Statements(t *testing.T) {
	statements := createUncategoriseStatements()
	categorizer := services.Categorize{
		Categories:   make(map[string]string),
		CategoryFile: "../fixtures/categoriesTest.yaml",
	}
	stmsNormalized, _ := categorizer.Categorise(statements)
	assert.Equal(t, 1, len(stmsNormalized))
	assert.Equal(t, "bills", stmsNormalized[0].Category)
}

func createUncategoriseStatements() []model.Statement {
	return []model.Statement{
		{
			TransactionDate:        "2016-07-29",
			TransactionType:        "ddd",
			TransactionDescription: "thames water 5191374174",
			DebitAmount:            2,
			CreditAmount:           1,
			Balance:                4.6,
		},
	}
}
