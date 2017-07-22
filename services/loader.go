package services

import (
	"fmt"
	"database/sql"
	"github.com/ancarebeca/expense-tracker/model"
)

//go:generate counterfeiter . DatabaseQueryConnection
type DatabaseQueryConnection interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
}

type Loader struct {
	DB DatabaseQueryConnection
}

func (l *Loader) Loader(statements []*model.Statement) {
	for _, s := range statements {
		err := l.create(s)
		if err != nil {
			fmt.Printf("Statement %v cannot be loaded: %s", s, err.Error())
		}
	}
}

func (l *Loader) create(s *model.Statement) error {
	_, err := l.DB.Exec("INSERT INTO `statements` (`transaction_date`, `transaction_type`, `transaction_description`, `debit_amount`, `credit_amount`, `balance`, `category`) VALUES (?, ?, ?, ?, ?, ?, ?)",
		s.TransactionDate,
		s.TransactionType,
		s.TransactionDescription,
		s.DebitAmount,
		s.CreditAmount,
		s.Balance,
		s.Category,
	)

	if err != nil {
		fmt.Printf("Error while loading statement row for : ", err.Error())
		return err
	}

	return nil
}
