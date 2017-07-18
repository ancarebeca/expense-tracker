package services_test

import (
	"database/sql"
	"fmt"
	"github.com/ancarebeca/expense-tracker/config"
	"github.com/ancarebeca/expense-tracker/services"
	"github.com/ancarebeca/expense-tracker/services/servicesfakes"
	_ "github.com/go-sql-driver/mysql"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Csv file is processed and return statements", func() {

	It("reads a valid csv and returns a statements", func() {
		var conf config.Conf
		statements := []*services.Statement{}

		conf.FilePath = "../fixtures/valid_csv.csv"
		csvReader := services.CsvReader{
			Conf: conf,
		}

		s, err := csvReader.ReadFromCsv(statements)
		Expect(err).NotTo(HaveOccurred())
		Expect(len(s)).To(Equal(17))
		Expect(s[0].TransactionDescription).To(Equal("Supermarket"))
		Expect(s[0].Balance).To(Equal(925.12))
		Expect(s[0].TransactionDate).To(Equal("29/07/2016"))
		Expect(s[0].CreditAmount).To(Equal(float64(0)))
		Expect(s[0].TransactionDescription).To(Equal("Supermarket"))

	})

	It("when the csv  is invalid it returns an error", func() {
		var conf config.Conf
		statements := []*services.Statement{}
		conf.FilePath = "../fixtures/invalid_format.csv"
		csvReader := services.CsvReader{
			Conf: conf,
		}
		_, err := csvReader.ReadFromCsv(statements)
		Expect(err).To(HaveOccurred())
	})

	It("when the configuration is invalid it returns an error", func() {
		var conf config.Conf
		statements := []*services.Statement{}
		conf.FilePath = "wrong_path.csv"
		csvReader := services.CsvReader{
			Conf: conf,
		}
		_, err := csvReader.ReadFromCsv(statements)
		Expect(err).To(HaveOccurred())
	})

	Describe("Statements are transformed into a proper format for the purposes of analysis", func() {

		It("transforms transaction data into a valid format", func() {
			statements := []*services.Statement{
				{
					TransactionDate:        "29/07/2016",
					TransactionDescription: "This is a tranSaction DeScription  ",
				},
			}

			t := services.Transformer{}
			statementsNormalized, err := t.Transform(statements)
			Expect(err).NotTo(HaveOccurred())
			Expect(statementsNormalized[0].TransactionDate).To(Equal("2016-07-29"), "Transforms TransactionDate")
			Expect(statementsNormalized[0].TransactionDescription).To(Equal("this is a transaction description"), "Transforms TransactionDate")
		})

		It("returns an error when the date is invalid", func() {
			statements := []*services.Statement{
				{
					TransactionDate: "29/2016/09",
				},
			}

			t := services.Transformer{}
			_, err := t.Transform(statements)
			Expect(err).To(HaveOccurred())
		})
	})

	Describe("Data is loaded into the final target database ", func() {

		var (
			result sql.Result
		)

		It("loads statements into database", func() {
			statements := []*services.Statement{
				{
					TransactionDate:        "2016-07-29",
					TransactionType:        "ddd",
					TransactionDescription: "bla bla bla",
					Category:               "category",
					DebitAmount:            2,
					CreditAmount:           1,
					Balance:                4.6,
				},
			}
			fakeDatabaseConnection := &servicesfakes.FakeDatabaseQueryConnection{}

			l := services.Loader{
				DB: fakeDatabaseConnection,
			}

			fakeDatabaseConnection.ExecReturns(result, nil)

			l.Loader(statements)
			Expect(fakeDatabaseConnection.ExecCallCount()).To(Equal(1))

		})

		It("returns an error when the date is invalid", func() {
			statements := []*services.Statement{
				{
					TransactionDate: "29/2016/09",
				},
			}

			db, err := sql.Open("mysql", "root:@/statements?charset=utf8")

			l := services.Loader{
				DB: db,
			}

			t := services.Transformer{}
			s, err := t.Transform(statements)
			l.Loader(s)
			Expect(err).To(HaveOccurred())
		})
	})

	Describe("Expenses are categorise by description ", func() {

		It("loads categories from yaml file into a map structure", func() {
			c := services.Categorize{
				Categories:   make(map[string]string),
				CategoryFile: "../fixtures/categoriesTest.yaml",
			}
			err := c.LoadCategories()
			Expect(err).NotTo(HaveOccurred())

			Expect(c.Categories["virgin"]).To(Equal("bills"))
			Expect(c.Categories["uber"]).To(Equal("transport"))
			Expect(c.Categories["sainsburys"]).To(Equal("groceries"))
		})

		It("returns an error is something goes wrong loading categories", func() {
			c := services.Categorize{
				Categories:   make(map[string]string),
				CategoryFile: "../fixtures/wrong-categoriesTest.yaml",
			}
			err := c.LoadCategories()
			Expect(err).To(HaveOccurred())
		})

		It("returns a category depending on key value introduced", func() {
			c := services.Categorize{
				CategoryFile: "../fixtures/categoriesTest.yaml",
				Categories:   make(map[string]string),
			}
			err := c.LoadCategories()
			Expect(err).NotTo(HaveOccurred())

			category, err := c.GetCategory("sainsburys supermarket")
			Expect(err).NotTo(HaveOccurred())
			Expect(category).To(Equal("groceries"))

			category, err = c.GetCategory("SAINSBURYS supermarket")
			Expect(err).NotTo(HaveOccurred())
			Expect(category).To(Equal("groceries"))

			category, err = c.GetCategory("water park")
			Expect(err).NotTo(HaveOccurred())
			Expect(category).To(Equal("entertainment"))

			category, err = c.GetCategory("cafe")
			Expect(err).NotTo(HaveOccurred())
			Expect(category).To(Equal("general"))

			category, err = c.GetCategory("THAMES WATER")
			Expect(err).NotTo(HaveOccurred())
			Expect(category).To(Equal("bills"))

		})

		It("categorise statements", func() {
			c := services.Categorize{
				Categories:   make(map[string]string),
				CategoryFile: "../fixtures/categoriesTest.yaml",
			}
			statements := []*services.Statement{
				{
					TransactionDate:        "2016-07-29",
					TransactionType:        "ddd",
					TransactionDescription: "thames water 5191374174",
					DebitAmount:            2,
					CreditAmount:           1,
					Balance:                4.6,
				},
			}

			err := c.LoadCategories()
			Expect(err).NotTo(HaveOccurred())

			s, err := c.Categorise(statements)
			Expect(err).NotTo(HaveOccurred())
			Expect(s[0].Category).To(Equal("bills"))
		})
	})

	Describe("A csv file is processed, transformed and loaded in a database", func() {
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
			statements := []*services.Statement{}

			csvReader := services.CsvReader{
				Conf: conf,
			}

			l := services.Loader{
				DB: db,
			}

			csvStatements, err := csvReader.ReadFromCsv(statements)
			Expect(err).NotTo(HaveOccurred())
			Expect(len(csvStatements)).To(Equal(17))

			t := services.Transformer{}
			statementsNormalized, err := t.Transform(csvStatements)
			Expect(err).NotTo(HaveOccurred())
			Expect(len(statementsNormalized)).To(Equal(17))
			Expect(statementsNormalized[0].TransactionDate).To(Equal("2016-07-29"), "Transforms TransactionDate")

			c := services.Categorize{
				Categories:   make(map[string]string),
				CategoryFile: "../fixtures/categoriesTest.yaml",
			}

			err = c.LoadCategories()
			Expect(err).NotTo(HaveOccurred())

			statementsCategorise, err := c.Categorise(statementsNormalized)
			Expect(err).NotTo(HaveOccurred())

			l.Loader(statementsCategorise)

			// query
			dbStatements := []*services.Statement{}

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

				s := services.Statement{
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
})
