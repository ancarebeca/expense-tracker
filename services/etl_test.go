package services_test

import (
	"github.com/ancarebeca/expense-tracker/config"
	"github.com/ancarebeca/expense-tracker/model"
	"github.com/ancarebeca/expense-tracker/services"
	"github.com/ancarebeca/expense-tracker/services/servicesfakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Running extractions, transformation and loading process ", func() {

	It("runs etl process", func() {
		conf := config.Conf{
			FilePath: "../fixtures/valid_csv.csv",
		}
		readerFake := &servicesfakes.FakeReader{}
		transfomerFake := &servicesfakes.FakeTransformer{}
		parserFake := &servicesfakes.FakeParser{}
		categorizeFake := &servicesfakes.FakeCategoriesLoader{}
		loaderFake := &servicesfakes.FakeStatementsLoader{}

		etl := services.Etl{
			conf,
			readerFake,
			transfomerFake,
			parserFake,
			categorizeFake,
			loaderFake,
		}

		expectedOutput := [][]string{
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
				"06/07/2036",
				"debit_card",
				"'444-444-444",
				"11111",
				"SuPerMArket  ",
				"19.2",
				"",
				"925.12",
			},
		}

		transformedOutput := [][]string{
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
				"06-07-2036",
				"debit_card",
				"'444-444-444",
				"11111",
				"supermarket",
				"19.2",
				"",
				"925.12",
			},
		}

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

		readerFake.ReadCsvReturns(expectedOutput, nil)
		transfomerFake.TransformReturns(transformedOutput, nil)
		parserFake.ParseReturns(statements, nil)
		categorizeFake.CategoriseReturns(statements, nil)
		err := etl.Run()

		Expect(err).NotTo(HaveOccurred())
		Expect(readerFake.ReadCsvCallCount()).To(Equal(1))

		Expect(parserFake.ParseCallCount()).To(Equal(1))
		Expect(parserFake.ParseArgsForCall(0)).To(Equal(transformedOutput))

		Expect(categorizeFake.CategoriseCallCount()).To(Equal(1))
		Expect(categorizeFake.CategoriseArgsForCall(0)).To(Equal(statements))

		Expect(loaderFake.LoadCallCount()).To(Equal(1))
		Expect(loaderFake.LoadArgsForCall(0)).To(Equal(statements))
	})
})
