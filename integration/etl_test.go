package integration_test

import (
	"database/sql"
	"fmt"
	"github.com/ancarebeca/expense-tracker/config"
	"github.com/ancarebeca/expense-tracker/services"
	_ "github.com/go-sql-driver/mysql"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ancarebeca/expense-tracker/model"
)

var _ = Describe("A csv file is processed, transformed and loaded in a database", func() {
	var (
		conf config.Conf
		db   *sql.DB
	)

	BeforeEach(func() {
		var err error
		conf.UserDb = "root"
		conf.PassDb = "root"
		conf.Database = "test_expenses"
		conf.FilePath = "../fixtures/valid_csv.csv"

		dataSourceName := fmt.Sprintf("%s:%s@/%s?charset=utf8", conf.UserDb, conf.PassDb, conf.Database)
		db, err = sql.Open("mysql", dataSourceName)
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		_, err := db.Exec(`TRUNCATE TABLE statements`)
		Expect(err).NotTo(HaveOccurred())
	})

	It("e2e", func() {
		r := services.CsvReader{
			Conf: conf,
		}

		l := services.LoadDb{
			DB: db,
		}

		t := services.DataTransformer{}
		p := services.DataParser{}

		c := services.Categorize{
			Categories:   make(map[string]string),
			CategoryFile: "../fixtures/categoriesTest.yaml",
		}

		etl := services.Etl{
			conf,
			&r,
			&t,
			&p,
			&c,
			&l,
		}
		etl.Run()

		// query
		dbStatements := []*model.Statement{}

		rows, err := db.Query("SELECT * FROM statements")

		Expect(err).NotTo(HaveOccurred())
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
			Expect(err).NotTo(HaveOccurred())

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

		Expect(len(dbStatements)).To(Equal(17))
		Expect(dbStatements[0].TransactionDescription).To(Equal("supermarket"))
		Expect(dbStatements[0].DebitAmount).To(Equal(19.2))
		Expect(dbStatements[0].Balance).To(Equal(925.12))
		Expect(dbStatements[0].TransactionDate).To(Equal("2016-07-29"))
		Expect(dbStatements[0].CreditAmount).To(Equal(float64(0)))
		Expect(dbStatements[0].TransactionDescription).To(Equal("supermarket"))
		Expect(dbStatements[1].CreditAmount).To(Equal(float64(90)))
		Expect(dbStatements[3].Category).To(Equal("bills"))
	})

})
