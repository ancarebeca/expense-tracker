package services

import (
	"database/sql"
	"fmt"
	"github.com/ancarebeca/expense-tracker/config"
	"github.com/gocarina/gocsv"
	"os"
	"strings"
	"time"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type Statement struct {
	TransactionDate        string  `csv:"Transaction Date"`
	TransactionType        string  `csv:"Transaction Type"`
	TransactionDescription string  `csv:"Transaction Description"`
	Category               string  `csv:"Category"`
	DebitAmount            float64 `csv:"Debit Amount"`
	CreditAmount           float64 `csv:"Credit Amount"`
	Balance                float64 `csv:"Balance"`
}


type CsvReader struct {
	Conf config.Conf
}

func (c *CsvReader) ReadFromCsv(statements []*Statement) ([]*Statement, error) {

	csvFile, err := os.Open(c.Conf.FilePath)

	if err != nil {
		fmt.Printf("Error while opening configuration: %s", err.Error())
		return nil, err
	}
	defer csvFile.Close()

	if err := gocsv.UnmarshalFile(csvFile, &statements); err != nil {
		fmt.Printf("Error while parsing the CSV: %s", err.Error())
		return nil, err
	}
	return statements, nil
}


type Transformer struct{}

var layoutInput = "02/01/2006"
var layoutOutput = "2006-01-02"

func (t *Transformer) Transform(statements []*Statement) ([]*Statement, error) {
	for _, s := range statements {
		date, err := t.transformDate(s.TransactionDate)
		if err != nil {
			return nil, err
		}

		s.TransactionDate = date
		s.TransactionDescription = t.cleanString(s.TransactionDescription)
	}
	return statements, nil
}

func (t *Transformer) transformDate(date string) (string, error) {

	stringOutput, err := time.Parse(layoutInput, date)

	if err != nil {
		fmt.Printf("Error parsing date: %s", err.Error())
		return "", err
	}
	return stringOutput.Format(layoutOutput), nil
}

func (t *Transformer) cleanString(value string) string {
	valueLower := strings.ToLower(value)
	return t.standardizeSpaces(valueLower)
}

func (t *Transformer) standardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

//go:generate counterfeiter . DatabaseQueryConnection
type DatabaseQueryConnection interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
}

type Loader struct {
	DB DatabaseQueryConnection
}


func (l *Loader) Loader(statements []*Statement){
	for _, s := range statements {
		err := l.create(s)
		if err != nil {
			fmt.Printf("Statement %v cannot be loaded: %s", s, err.Error())
		}
	}
}

func (l *Loader) create(s *Statement) error {

	_, err := l.DB.Exec("INSERT INTO `statements` (`transaction_date`, `transaction_type`, `transaction_description`, `debit_amount`, `credit_amount`, `balance`, `category`) VALUES (?, ?, ?, ?, ?, ?, ?)",
		s.TransactionDate,
		s.TransactionType,
		s.TransactionDescription,
		s.DebitAmount,
		s.CreditAmount,
		s.Balance,
		s.Category,
	)

	if err != nil {
		fmt.Printf("Error while loading statement row for : ", err.Error())
		return err
	}

	return nil
}

//--------------


const general = "general"

type Categorize struct {
	Categories map[string]string
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


func (c *Categorize) Categorise(statements []*Statement) ([]*Statement, error)  {
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