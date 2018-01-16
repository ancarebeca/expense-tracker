package etl

import (
	"fmt"
	"github.com/ancarebeca/expense-tracker/model"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strings"
)

const general = "general"

type CategoriesLoader interface {
	Categorise(statements []model.Statement) ([]model.Statement, error)
}

type Categorize struct {
	Categories   map[string]string
	CategoryFile string
}

func (c *Categorize) Categorise(statements []model.Statement) ([]model.Statement, error) {
	c.loadCategories()
	stms := []model.Statement{}
	for _, s := range statements {
		category, err := c.getCategory(s.TransactionDescription)
		if err != nil {
			fmt.Printf("Statement %v cannot be categorise: %s", s, err.Error())
			return nil, err
		}

		s.Category = category
		stms = append(stms, s)
	}

	return stms, nil
}

func (c *Categorize) loadCategories() error {
	yamlFile, err := ioutil.ReadFile(c.CategoryFile)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = yaml.Unmarshal(yamlFile, c.Categories)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}

func (c *Categorize) getCategory(keyword string) (string, error) {
	for k, v := range c.Categories {
		kLower := strings.ToLower(k)
		if strings.Contains(strings.ToLower(keyword), kLower) {
			return v, nil
		}
	}
	return general, nil
}
