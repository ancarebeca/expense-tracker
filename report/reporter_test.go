package report_test

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/ancarebeca/expense-tracker/config"
	"github.com/ancarebeca/expense-tracker/report"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

func Test_YearlySpedingByCategoryReport(t *testing.T) {

	var path = "../config/config_test.yaml"
	conf := config.Conf{}
	conf.LoadConfig(path)

	dataSourceName := fmt.Sprintf("%s:%s@/%s?charset=utf8", conf.UserDb, conf.PassDb, conf.Database)
	db, err := sql.Open("mysql", dataSourceName)

	if err != nil {
		fmt.Println("Error (1) ", err.Error())
	}

	calculator := &report.CalculateYearlySpendingByCategory{
		conf,
		db,
	}

	adapter := report.DataAdapter{}
	rysc := report.YearlySpendingByCategoryReport{
		calculator,
		adapter,
	}

	report := report.Report{
		"output.csv",
		rysc,
	}
	report.Create()
}
