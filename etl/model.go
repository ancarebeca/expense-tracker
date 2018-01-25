package etl

type Statement struct {
	TransactionDate        string
	TransactionType        string
	TransactionDescription string
	Category               string
	DebitAmount            float64
	CreditAmount           float64
	Balance                float64
}
