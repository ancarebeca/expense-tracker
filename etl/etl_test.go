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
	db             *sql.DB
	conf           config.Conf
	confPath       = "../config/config-test.yaml"
	dataSourceName string
	err            error
)

func Test_etl_extractFromCsv_transformData_loadIntoDatabase(t *testing.T) {
	etl := createEtl()
	etl.Run()
	dbStatements := getStatementsFromDb()
	assert.Equal(t, 36, len(dbStatements))
	assert.Equal(t, "supermarket", dbStatements[0].TransactionDescription)
	assert.Equal(t, 19.2, dbStatements[0].DebitAmount)
	assert.Equal(t, 925.12, dbStatements[0].Balance)
	assert.Equal(t, "2016-07-29", dbStatements[0].TransactionDate)
	assert.Equal(t, float64(0), dbStatements[0].CreditAmount)
	cleanDb()
}

func createEtl() Etl {
	conf = config.Conf{}
	conf.LoadConfig(confPath)
	fmt.Println(dataSourceName)

	dataSourceName = fmt.Sprintf("%s:%s@/%s?charset=utf8", conf.UserDb, conf.PassDb, conf.Database)
	db, err = sql.Open("mysql", dataSourceName)
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

	return Etl{
		conf,
		&r,
		&dt,
		&p,
		&c,
		&l,
	}
}

func getStatementsFromDb() []*model.Statement {
	dbStatements := []*model.Statement{}
	rows, err := db.Query("SELECT * FROM statements")
	if err != nil {
		fmt.Println("Error when getting Statements: ", err.Error())
		cleanDb()
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
			cleanDb()
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
	return dbStatements
}

func cleanDb() {
	_, err = db.Exec(`TRUNCATE TABLE statements`)
	if err != nil {
		fmt.Println(err.Error())
	}
}
