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
	var db *sql.DB

	conf := config.Conf{}
	conf.LoadConfig(path)

	csvReader := services.CsvReader{
		Conf: conf,
	}

	transformer := services.DataTransformer{}

	parser := services.DataParser{}

	categorizer := services.Categorize{
		Categories:   make(map[string]string),
		CategoryFile: "config/categories.yaml",
	}

	dataSourceName := fmt.Sprintf("%s:%s@/%s?charset=utf8", conf.UserDb, conf.PassDb, conf.Database)
	db, err := sql.Open("mysql", dataSourceName)

	if err != nil {
		panic(err.Error())
	}

	loader := services.LoadDb{
		DB: db,
	}
	etl := services.Etl{
		conf,
		&csvReader,
		&transformer,
		&parser,
		&categorizer,
		&loader,
	}
	etl.Run()
}
