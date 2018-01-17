package etl_test

import (
	"testing"

	"github.com/ancarebeca/expense-tracker/etl"
	"github.com/stretchr/testify/assert"
)

func Test_parser_dataIntoValidValues(t *testing.T) {
	p := etl.SantanderParser{}
	stms := p.Parse(getData())
	assert.Equal(t, 2, len(stms))

	assert.Equal(t, "29/07/2016", stms[0].TransactionDate)
	assert.Equal(t, "debit_card", stms[0].TransactionType)
	assert.Equal(t, "Description 1", stms[0].TransactionDescription)
	assert.Equal(t, 19.2, stms[0].DebitAmount)
	assert.Equal(t, 4.0, stms[0].CreditAmount)
	assert.Equal(t, 3.12, stms[0].Balance)
}

func getData() [][]string {
	return [][]string{
		{
			"Transaction Date",
			"Transaction Type",
			"Sort Code",
			"Account Number",
			"Transaction Description",
			"Debit Amount",
			"Credit Amount",
			"Balance",
		},
		{
			"29/07/2016",
			"debit_card",
			"'444-444-444",
			"11111",
			"Description 1",
			"19.2",
			"4.0",
			"3.12",
		},
		{
			"29/07/2017",
			"credit_card",
			"'444-4r4-444",
			"111e1",
			"Description 2",
			"1.2",
			"4.5",
			"3.22",
		},
		{
			"29/07/2016",
			"debit_card",
			"'444-444-444",
			"11111",
			"Description 2",
			"19.2",
			"4.22",
			"wrong",
		},
		{
			"29/07/2016",
			"debit_card",
			"'444-444-444",
			"11111",
			"Description 2",
			"19.2",
			"wrong",
			"3",
		},
		{
			"29/07/2016",
			"debit_card",
			"'444-444-444",
			"11111",
			"Description 2",
			"wrong",
			"4",
			"3",
		},
	}
}
