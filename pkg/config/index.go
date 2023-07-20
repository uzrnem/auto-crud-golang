package config

import (
	"fmt"
	"io/ioutil"
	"strconv"

	"autocrud/pkg/db"
	"autocrud/pkg/integer"
	"gopkg.in/yaml.v3"
)

const (
	yamlFile = "/app/files/config.yaml"
)

var (
	Config Configuration
)

type Response struct {
	StatusCode int               `yaml:"status_code" json:"status_code"`
	Body       map[string]string `yaml:"body" json:"body"`
}
type Api struct {
	Type     string   `yaml:"type" json:"type"`
	Path     string   `yaml:"path" json:"path"`
	Method   string   `yaml:"method" json:"method"`
	Response Response `yaml:"response" json:"response"`
}
type Field struct {
	Type     string `yaml:"type" json:"type"`
	Required bool   `yaml:"required" json:"required"`
	Min      int    `yaml:"min" json:"min"`
	Max      int    `yaml:"max" json:"max"`
}
type ID struct {
	Name string `yaml:"name" json:"name"`
	Type string `yaml:"type" json:"type"`
}
type Documents struct {
	ID     ID               `yaml:"id" json:"id"`
	Path   string           `yaml:"path" json:"path"`
	Fields map[string]Field `yaml:"fields" json:"fields"`
	Apis   map[string]Api   `yaml:"apis" json:"apis"`
}
type Application struct {
	Name          string               `yaml:"name" json:"name"`
	Port          int                  `yaml:"port" json:"port"`
	PathExtention string               `yaml:"path_extention" json:"path_extention"`
	Documents     map[string]Documents `yaml:"documents" json:"documents"`
}
type Configuration struct {
	Application Application `yaml:"application" json:"application"`
}

func Load() error {

	data, err := ioutil.ReadFile(yamlFile)
	if err != nil {
		return err
	}

	// Unmarshal the YAML data into the struct
	err = yaml.Unmarshal(data, &Config)
	if err != nil {
		return err
	}
	for key, value := range Config.Application.Documents {
		if len(value.Fields) > 0 && value.ID.Type != "int" && value.ID.Type != "uuid" {
			return fmt.Errorf("Id type `%s` not allowed in Document[%s]", value.ID.Type, key)
		}
		if value.ID.Type == "int" {
			data, err := db.DB.GetDocuments(key)
			if err != nil {
				return err
			} else {
				for id, _ := range data {
					intVar, err := strconv.Atoi(id)
					if err != nil {
						return fmt.Errorf("Invalid Id[%s] found in in Document[%s]", id, key)
					}
					integer.SetCounter(intVar)
				}
			}

		}
	}
	return nil
}
