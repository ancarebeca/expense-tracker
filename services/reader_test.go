package services_test

import (
	"github.com/ancarebeca/expense-tracker/config"
	"github.com/ancarebeca/expense-tracker/services"
	_ "github.com/go-sql-driver/mysql"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Reading csv file", func() {

	FIt("reads a valid csv and returns a statements", func() {
		var conf config.Conf

		conf.FilePath = "../fixtures/valid_csv.csv"
		csvReader := services.CsvReader{
			Conf: conf,
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
				"29/07/2016",
				"debit_card",
				"'444-444-444",
				"11111",
				"Supermarket",
				"19.2",
				"",
				"925.12",
			},
		}

		lines, err := csvReader.ReadCsv()
		Expect(err).NotTo(HaveOccurred())
		Expect(len(lines)).To(Equal(18))
		Expect(lines[0]).To(Equal(expectedOutput[0]))
		Expect(lines[1]).To(Equal(expectedOutput[1]))

	})

	It("when the csv  is invalid it returns an error", func() {
		var conf config.Conf
		conf.FilePath = "../fixtures/invalid_format.csv"
		csvReader := services.CsvReader{
			Conf: conf,
		}
		_, err := csvReader.ReadCsv()
		Expect(err).To(HaveOccurred())
	})

	It("when the configuration is invalid it returns an error", func() {
		var conf config.Conf
		conf.FilePath = "wrong_path.csv"
		csvReader := services.CsvReader{
			Conf: conf,
		}
		_, err := csvReader.ReadCsv()
		Expect(err).To(HaveOccurred())
	})
})
