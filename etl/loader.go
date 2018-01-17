package etl

import (
	"github.com/ancarebeca/expense-tracker/model"
	"github.com/ancarebeca/expense-tracker/repository"
	"github.com/sirupsen/logrus"
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
			logrus.Errorf("Statement %v cannot be loaded: %s", s, err.Error())
			panic(err)
		}
	}
}
