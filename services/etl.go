package services

import "github.com/ancarebeca/expense-tracker/config"

type Etl struct {
	Conf            config.Conf
	CsvReader       Reader
	DataTransformer Transformer
	DataParser      Parser
	Categories      CategoriesLoader
	Loader          StatementsLoader
}

func (e Etl) Run() error {

	data, err := e.CsvReader.ReadCsv()
	if err != nil {
		return err
	}

	dataNormalized, err := e.DataTransformer.Transform(data)
	if err != nil {
		return err
	}

	statements, err := e.DataParser.Parse(dataNormalized)
	if err != nil {
		panic(err.Error())
	}

	err = e.Categories.Load()
	if err != nil {
		panic(err.Error())
	}

	statementsCategorise, err := e.Categories.Categorise(statements)
	if err != nil {
		panic(err.Error())
	}

	e.Loader.Load(statementsCategorise)

	return nil
}
