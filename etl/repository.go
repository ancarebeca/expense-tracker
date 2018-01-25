package etl

import (
	"database/sql"
	"log"
)

type DatabaseQueryConnection interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
}

type StatementRepository interface {
	Create(s *Statement) error
}

type RepositoryDb struct {
	DB DatabaseQueryConnection
}

func (l *RepositoryDb) Create(s *Statement) error {
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
