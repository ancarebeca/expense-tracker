package repository

import (
	"database/sql"
	"log"

	"github.com/ancarebeca/expense-tracker/model"
)

type DatabaseQueryConnection interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
}

type Fetch interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
}

type StatementRepository interface {
	Create(s *model.Statement) error
}

type RepositoryDb struct {
	DB DatabaseQueryConnection
}

func (l *RepositoryDb) Create(s *model.Statement) error {
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
		log.Fatal("Error while inserting statements  into DB: ", err.Error())
		return err
	}

	return nil
}
