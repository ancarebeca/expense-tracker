package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Conf struct {
	UserDb     string `yaml:"user_db"`
	PassDb     string `yaml:"pass_db"`
	Database   string `yaml:"database"`
	FilePath   string `yaml:"file_path"`
}

func (c *Conf) LoadConfig(path string) error {

	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Configuration file not found or was not provided [%s]", path)
		return err
	}

	err = yaml.Unmarshal(data, c)
	if err != nil {
		fmt.Println("Configuration file invalid")
		return err
	}

	return nil
}
