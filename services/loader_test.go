package services_test
//
//import (
//	"database/sql"
//	"github.com/ancarebeca/expense-tracker/services"
//	"github.com/ancarebeca/expense-tracker/services/servicesfakes"
//	_ "github.com/go-sql-driver/mysql"
//	. "github.com/onsi/ginkgo"
//	. "github.com/onsi/gomega"
//)
//
//var _ = Describe("Data is loaded into the final target database ", func() {
//
//	var (
//		result sql.Result
//	)
//
//	It("loads statements into database", func() {
//		statements := []*services.Statement{
//			{
//				TransactionDate:        "2016-07-29",
//				TransactionType:        "ddd",
//				TransactionDescription: "bla bla bla",
//				Category:               "category",
//				DebitAmount:            2,
//				CreditAmount:           1,
//				Balance:                4.6,
//			},
//		}
//		fakeDatabaseConnection := &servicesfakes.FakeDatabaseQueryConnection{}
//
//		l := services.Loader{
//			DB: fakeDatabaseConnection,
//		}
//
//		fakeDatabaseConnection.ExecReturns(result, nil)
//
//		l.Loader(statements)
//		Expect(fakeDatabaseConnection.ExecCallCount()).To(Equal(1))
//
//	})
//
//	It("returns an error when the date is invalid", func() {
//		statements := []*services.Statement{
//			{
//				TransactionDate: "29/2016/09",
//			},
//		}
//
//		db, err := sql.Open("mysql", "root:@/statements?charset=utf8")
//
//		l := services.Loader{
//			DB: db,
//		}
//
//		t := services.Transformer{}
//		s, err := t.Transform(statements)
//		l.Loader(s)
//		Expect(err).To(HaveOccurred())
//	})
//})
