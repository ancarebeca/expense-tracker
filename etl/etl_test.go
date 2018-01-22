package etl

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/ancarebeca/expense-tracker/config"
	"github.com/ancarebeca/expense-tracker/model"
	"github.com/ancarebeca/expense-tracker/repository"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

var (
	db *sql.DB
)

func Test_etl(t *testing.T) {
	conf := config.Conf{}
	conf.LoadConfig("../config/config_test.yaml")

	dataSourceName := fmt.Sprintf("%s:%s@/%s?charset=utf8", conf.UserDb, conf.PassDb, conf.Database)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Println(err.Error())
	}
	r := CsvReader{}

	repository := repository.RepositoryDb{
		DB: db,
	}
	l := LoadStatements{
		&repository,
	}

	dt := DataTransformer{}
	p := SantanderParser{}

	c := Categorize{
		Categories:   make(map[string]string),
		CategoryFile: conf.CategoryPath,
	}

	etl := Etl{
		conf,
		&r,
		&dt,
		&p,
		&c,
		&l,
	}
	etl.Run()

	//query
	dbStatements := []*model.Statement{}

	rows, err := db.Query("SELECT * FROM statements")
	if err != nil {
		fmt.Println(err.Error())
	}

	for rows.Next() {
		var id int
		var transaction_date string
		var transaction_type string
		var sortCode sql.NullString
		var accountNumber sql.NullString
		var transaction_description string
		var debit_amount float64
		var credit_amount float64
		var balance float64
		var category string

		err = rows.Scan(
			&id,
			&transaction_date,
			&transaction_type,
			&sortCode,
			&accountNumber,
			&transaction_description,
			&debit_amount,
			&credit_amount,
			&balance,
			&category,
		)
		if err != nil {
			fmt.Println(err.Error())
		}

		s := model.Statement{
			TransactionDate:        transaction_date,
			TransactionType:        transaction_type,
			TransactionDescription: transaction_description,
			Category:               category,
			CreditAmount:           credit_amount,
			DebitAmount:            debit_amount,
			Balance:                balance,
		}
		dbStatements = append(dbStatements, &s)
	}
	assert.Equal(t, 36, len(dbStatements))
	assert.Equal(t, "supermarket", dbStatements[0].TransactionDescription)
	assert.Equal(t, 19.2, dbStatements[0].DebitAmount)
	assert.Equal(t, 925.12, dbStatements[0].Balance)
	assert.Equal(t, "2016-07-29", dbStatements[0].TransactionDate)
	assert.Equal(t, float64(0), dbStatements[0].CreditAmount)

	_, err = db.Exec(`TRUNCATE TABLE statements`)
	if err != nil {
		fmt.Println(err.Error())
	}

}
