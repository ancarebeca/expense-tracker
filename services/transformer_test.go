package services_test

import (

	"github.com/ancarebeca/expense-tracker/services"
	_ "github.com/go-sql-driver/mysql"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("csv data are transformed into a proper format for the purposes of analysis", func() {

	It("transforms csv data into a valid format", func() {
		input := [][]string{
			{
				"Transaction Date",
				"Transaction Type",
				"Sort Code",
				"Account Number",
				"Transaction Description",
				"Debit Amount",
				"Credit Amount",
				"Balance",
			},
			{
				"29/07/2016",
				"debit_card",
				"'444-444-444",
				"11111",
				"This is a tranSaction DeScription  ",
				"19.2",
				"",
				"925.12",
			},
			{
				"29/07/2016",
				"debit_card",
				"'444-444-444",
				"11111",
				"bla bla",
				"19.2",
				"",
				"925.12",
			},
		}

		t := services.Transformer{}
		inputNormalized, err := t.Transform(input)
		Expect(err).NotTo(HaveOccurred())

		Expect(len(inputNormalized)).To(Equal(2))
		Expect(inputNormalized[0][0]).To(Equal("2016-07-29"), "Transforms TransactionDate")
		Expect(inputNormalized[0][4]).To(Equal("this is a transaction description"), "transaction description")
	})

	It("returns an error when the date is invalid", func() {
		input := [][]string{
			{
				"Transaction Date",
				"Transaction Type",
				"Sort Code",
				"Account Number",
				"Transaction Description",
				"Debit Amount",
				"Credit Amount",
				"Balance",
			},
			{
				"29-07-2016",
				"debit_card",
				"'444-444-444",
				"11111",
				"This is a tranSaction DeScription  ",
				"19.2",
				"",
				"925.12",
			},
			{
				"29/07/2016",
				"debit_card",
				"'444-444-444",
				"11111",
				"bla bla",
				"19.2",
				"",
				"925.12",
			},
		}


		t := services.Transformer{}
		_, err := t.Transform(input)
		Expect(err).To(HaveOccurred())
	})
})

