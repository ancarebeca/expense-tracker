package services_test

import (
	"github.com/ancarebeca/expense-tracker/model"
	"github.com/ancarebeca/expense-tracker/services"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_transform_normalizeDataModel(t *testing.T) {
	statements := createStatements()
	transformer := services.DataTransformer{}
	stmsNormalized := transformer.Transform(statements)
	assert.Equal(t, 2, len(stmsNormalized))
	assert.Equal(t, "2016-01-02", stmsNormalized[1].TransactionDate)
	assert.Equal(t, "rrr", stmsNormalized[1].TransactionType)
	assert.Equal(t, "description 2", stmsNormalized[1].TransactionDescription)
	assert.Equal(t, 19.2, stmsNormalized[1].DebitAmount)
	assert.Equal(t, 4.0, stmsNormalized[1].CreditAmount)
	assert.Equal(t, 3.12, stmsNormalized[1].Balance)
}

func createStatements() []model.Statement {
	return []model.Statement{
		{
			TransactionDate:        "02/01/2006",
			TransactionType:        "ddd",
			TransactionDescription: "Description 1",
			Category:               "category",
			DebitAmount:            2,
			CreditAmount:           1,
			Balance:                4.6,
		},
		{
			TransactionDate:        "02/01/2016",
			TransactionType:        "rrr",
			TransactionDescription: "Description 2",
			Category:               "category 2",
			DebitAmount:            19.2,
			CreditAmount:           4.0,
			Balance:                3.12,
		},
		{
			TransactionDate:        "wrong date",
			TransactionType:        "rrr",
			TransactionDescription: "Description 2",
			Category:               "category 2",
			DebitAmount:            19.2,
			CreditAmount:           4.0,
			Balance:                3.12,
		},
	}
}
