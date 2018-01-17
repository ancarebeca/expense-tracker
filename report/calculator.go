package report

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/ancarebeca/expense-tracker/config"
	"github.com/ancarebeca/expense-tracker/repository"
)

type CalculateYearlySpendingByCategory struct {
	Conf config.Conf
	DB   repository.Fetch
}

type Calculator interface {
	CalculateExpenses() ([]ReportModel, error)
}

func (r *CalculateYearlySpendingByCategory) CalculateExpenses() ([]ReportModel, error) {
	rows, err := r.DB.Query("SELECT YEAR(transaction_date) AS year, category AS category, sum(debit_amount) AS totalAmount FROM statements GROUP BY YEAR(transaction_date), category   ORDER BY category, YEAR(transaction_date)")
	if err != nil {
		fmt.Printf("Error while loading statement row for : ", err.Error())
		panic(err)
	}

	data, err := r.createReport(rows)
	if err != nil {
		fmt.Printf("Error while mapping row for : ", err.Error())
		panic(err)
	}

	return data, nil
}

func (r *CalculateYearlySpendingByCategory) createReport(rows *sql.Rows) ([]ReportModel, error) {
	var data []ReportModel

	for rows.Next() {
		var category, totalAmount, year string
		err := rows.Scan(&year, &category, &totalAmount)

		if err != nil {
			log.Fatal(err)
			panic(err)
		}

		var record []string
		record = append(record, category)
		record = append(record, year)
		record = append(record, totalAmount)

		m := ReportModel{
			Year:     year,
			Amount:   totalAmount,
			Category: category,
		}

		data = append(data, m)
	}
	return data, nil
}
