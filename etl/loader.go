package etl

import (
	"github.com/sirupsen/logrus"
)

type Loader interface {
	Load(statements []Statement)
}

type LoadStatements struct {
	Loader StatementRepository
}

func (l *LoadStatements) Load(statements []Statement) {
	for _, s := range statements {
		err := l.Loader.Create(&s)
		if err != nil {
			logrus.Errorf("Statement %v cannot be loaded: %s", s, err.Error())
			panic(err)
		}
	}
}
