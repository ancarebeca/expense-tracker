package etl

import (
	"fmt"
	"github.com/ancarebeca/expense-tracker/model"
	"github.com/ancarebeca/expense-tracker/repository"
	"log"
)

type Loader interface {
	Load(statements []model.Statement)
}

type LoadStatements struct {
	Loader repository.StatementRepository
}

func (l *LoadStatements) Load(statements []model.Statement) {
	for _, s := range statements {
		err := l.Loader.Create(&s)
		if err != nil {
			log.Fatal(fmt.Sprintf("Statement %v cannot be loaded:", s), err.Error())
			panic(err)
		}
	}
}
