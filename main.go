package main

import (
	"database/sql"
	"fmt"

	"github.com/ancarebeca/expense-tracker/config"
	"github.com/ancarebeca/expense-tracker/etl"
	_ "github.com/go-sql-driver/mysql"
)

var path = "config/config.yaml"

func main() {
	var db *sql.DB
	conf := config.Conf{}
	conf.LoadConfig(path)
	callEtl(conf, db)
}

func callEtl(conf config.Conf, db *sql.DB) {

	dataSourceName := fmt.Sprintf("%s:%s@/%s?charset=utf8", conf.UserDb, conf.PassDb, conf.Database)
	db, _ = sql.Open("mysql", dataSourceName)
	r := etl.CsvReader{}

	repository := etl.RepositoryDb{
		DB: db,
	}
	l := etl.LoadStatements{
		&repository,
	}

	t := etl.DataTransformer{}
	p := etl.SantanderParser{}

	c := etl.Categorize{
		Categories:   make(map[string]string),
		CategoryFile: conf.CategoryPath,
	}

	etl := etl.Etl{
		conf,
		&r,
		&t,
		&p,
		&c,
		&l,
	}
	etl.Run()
}
