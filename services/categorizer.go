package services

import (
	"gopkg.in/yaml.v2"
	"fmt"
	"strings"
	"io/ioutil"
	"github.com/ancarebeca/expense-tracker/model"
)

const general = "general"

type Categorize struct {
	Categories   map[string]string
	CategoryFile string
}

func (c *Categorize) LoadCategories() error {
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

func (c *Categorize) GetCategory(keyword string) (string, error) {

	for k, v := range c.Categories {
		kLower := strings.ToLower(k)
		if strings.Contains(strings.ToLower(keyword), kLower) {
			return v, nil
		}
	}
	return general, nil
}

func (c *Categorize) Categorise(statements []*model.Statement) ([]*model.Statement, error) {
	for _, s := range statements {
		category, err := c.GetCategory(s.TransactionDescription)
		if err != nil {
			fmt.Printf("Statement %v cannot be categorise: %s", s, err.Error())
			return nil, err
		}

		s.Category = category
	}

	return statements, nil
}
