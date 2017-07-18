package main

import (
	"database/sql"
	"fmt"
	"github.com/ancarebeca/expense-tracker/config"
	"github.com/ancarebeca/expense-tracker/services"
	_ "github.com/go-sql-driver/mysql"
)

var path = "config/config.yaml"

func main() {
	statements := []*services.Statement{}

	conf := config.Conf{}
	conf.LoadConfig(path)

	csvReader := services.CsvReader{
		Conf: conf,
	}

	csvStatements, err := csvReader.ReadFromCsv(statements)
	if err != nil {
		panic(err.Error())
	}

	t := services.Transformer{}
	statementsNormalized, err := t.Transform(csvStatements)
	if err != nil {
		panic(err.Error())
	}

	var db *sql.DB

	dataSourceName := fmt.Sprintf("%s:%s@/%s?charset=utf8", conf.UserDb, conf.PassDb, conf.Database)
	db, err = sql.Open("mysql", dataSourceName)

	if err != nil {
		panic(err.Error())
	}

	c := services.Categorize{
		Categories:   make(map[string]string),
		CategoryFile: "config/categories.yaml",
	}

	err = c.LoadCategories()
	if err != nil {
		panic(err.Error())
	}

	statementsCategorise, err := c.Categorise(statementsNormalized)
	if err != nil {
		panic(err.Error())
	}

	l := services.Loader{
		DB: db,
	}

	l.Loader(statementsCategorise)

}
