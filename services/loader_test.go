package services_test

import (
	"database/sql"
	"github.com/ancarebeca/expense-tracker/services"
	"github.com/ancarebeca/expense-tracker/services/servicesfakes"
	_ "github.com/go-sql-driver/mysql"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ancarebeca/expense-tracker/model"
)

var _ = Describe("Data is loaded into the final target database ", func() {

	var (
		result sql.Result
	)

	It("loads statements into database", func() {
		statements := []*model.Statement{
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

})
