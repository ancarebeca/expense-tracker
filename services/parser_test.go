package services_test

import (
	"github.com/ancarebeca/expense-tracker/services"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Parser string into statement", func() {

	It("parser a string statement into a model statement", func() {

		input := [][]string{
			{
				"29/07/2016",
				"debit_card",
				"'444-444-444",
				"11111",
				"Description 1",
				"19.2",
				"",
				"3.12",
			},
			{
				"29/07/2016",
				"debit_card",
				"'444-444-444",
				"11111",
				"Description 2",
				"19.2",
				"",
				"3.12",
			},
		}

		p := services.Parse{}
		statements, err := p.Parser(input)

		Expect(err).ToNot(HaveOccurred())
		Expect(len(statements)).To(Equal(2))
		Expect(statements[0].TransactionDescription).To(Equal("Description 1"))
		Expect(statements[1].TransactionDescription).To(Equal("Description 2"))

	})

	It("returns an error when debit amount cannot be cast to float", func() {

		input := [][]string{
			{
				"29/07/2016",
				"debit_card",
				"'444-444-444",
				"11111",
				"Description 1",
				"error",
				"",
				"3.12",
			},
		}
		p := services.Parse{}
		_, err := p.Parser(input)

		Expect(err).To(HaveOccurred())
	})

	It("returns an error when debit amount cannot be cast to float", func() {

		input := [][]string{
			{
				"29/07/2016",
				"debit_card",
				"'444-444-444",
				"11111",
				"Description 1",
				"error",
				"",
				"3.12",
			},
		}
		p := services.Parse{}
		_, err := p.Parser(input)

		Expect(err).To(HaveOccurred())
	})

	It("returns an error when credit amount cannot be cast to float", func() {

		input := [][]string{
			{
				"29/07/2016",
				"debit_card",
				"'444-444-444",
				"11111",
				"Description 1",
				"23",
				"error",
				"925.12",
			},
		}
		p := services.Parse{}
		_, err := p.Parser(input)

		Expect(err).To(HaveOccurred())
	})
	It("returns an error when credit amount cannot be cast to float", func() {

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
				"Description 1",
				"23",
				"3",
				"error",
			},
		}
		p := services.Parse{}
		_, err := p.Parser(input)

		Expect(err).To(HaveOccurred())
	})
})
