package services_test

import (
	"github.com/ancarebeca/expense-tracker/services"
	_ "github.com/go-sql-driver/mysql"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ancarebeca/expense-tracker/model"
)

var _ = Describe("Expenses are categorise by description ", func() {

	It("loads categories from yaml file into a map structure", func() {
		c := services.Categorize{
			Categories:   make(map[string]string),
			CategoryFile: "../fixtures/categoriesTest.yaml",
		}
		err := c.Load()
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
		err := c.Load()
		Expect(err).To(HaveOccurred())
	})

	It("categorise statements", func() {
		c := services.Categorize{
			Categories:   make(map[string]string),
			CategoryFile: "../fixtures/categoriesTest.yaml",
		}
		statements := []*model.Statement{
			{
				TransactionDate:        "2016-07-29",
				TransactionType:        "ddd",
				TransactionDescription: "thames water 5191374174",
				DebitAmount:            2,
				CreditAmount:           1,
				Balance:                4.6,
			},
		}

		err := c.Load()
		Expect(err).NotTo(HaveOccurred())

		s, err := c.Categorise(statements)
		Expect(err).NotTo(HaveOccurred())
		Expect(s[0].Category).To(Equal("bills"))
	})
})
