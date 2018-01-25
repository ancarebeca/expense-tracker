package etl

import (
	"io/ioutil"
	"strings"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

const general = "general"

type CategoriesLoader interface {
	Categorise(statements []Statement) ([]Statement, error)
}

type Categorize struct {
	Categories   map[string]string
	CategoryFile string
}

func (c *Categorize) Categorise(statements []Statement) ([]Statement, error) {
	err := c.loadCategories()
	if err != nil {
		return nil, err
	}

	stms := []Statement{}
	for _, s := range statements {
		s.Category = c.getCategory(s.TransactionDescription)
		stms = append(stms, s)
	}

	return stms, nil
}

func (c *Categorize) loadCategories() error {
	yamlFile, err := ioutil.ReadFile(c.CategoryFile)
	if err != nil {
		logrus.Errorf("Categorizer, cannot load category file: %s", err.Error())
		return err
	}

	err = yaml.Unmarshal(yamlFile, c.Categories)
	if err != nil {
		logrus.Errorf("Categorizer, cannot unmarshal category file: %s", err.Error())
		return err
	}

	return nil
}

func (c *Categorize) getCategory(keyword string) string {
	for k, v := range c.Categories {
		kLower := strings.ToLower(k)
		if strings.Contains(strings.ToLower(keyword), kLower) {
			return v
		}
	}
	return general
}
