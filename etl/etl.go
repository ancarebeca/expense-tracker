package etl

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/ancarebeca/expense-tracker/config"
)

type Etl struct {
	Conf            config.Conf
	CsvReader       Reader
	DataTransformer Transformer
	DataParser      Parser
	Categories      CategoriesLoader
	Loader          Loader
}

func (e Etl) Run() error {

	files, err := e.getFiles(e.Conf.DirPath)
	if err != nil {
		panic(err.Error())
	}

	for i := 0; i < len(files); i++ {
		data, err := e.CsvReader.ReadCsv(files[i])
		if err != nil {
			panic(err.Error())
		}

		statements := e.DataParser.Parse(data)
		dataNormalized := e.DataTransformer.Transform(statements)
		statementsCategorise, err := e.Categories.Categorise(dataNormalized)
		if err != nil {
			panic(err.Error())
		}
		e.Loader.Load(statementsCategorise)
	}
	return nil
}

func (e Etl) getFiles(dir string) ([]string, error) {
	names := []string{}

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return nil, err
	}

	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			n := fmt.Sprintf("%s/%s", e.Conf.DirPath, info.Name())
			names = append(names, n)

		}
		return nil
	})

	if len(names) == 0 {
		return nil, errors.New("Empty folder")
	}
	return names, nil
}
