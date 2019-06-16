package tpl

import (
	"encoding/json"
	"io/ioutil"

	yaml "gopkg.in/yaml.v3"
)

// Card represents a single work item
type Card struct {
	Name        string   `yaml:"Name"`
	Description string   `yaml:"Description"`
	Labels      []string `yaml:"Labels"`
}

// Label represents a Trello label
type Label struct {
	Name  string `yaml:"Name"`
	Color string `yaml:"Color"`
	ID    string `yaml:"ID"`
}

// Board represents a board
type Board struct {
	Name        string            `yaml:"Name"`
	Description string            `yaml:"Description"`
	Background  string            `yaml:"Background"`
	Labels      []Label           `yaml:"Labels"`
	Lists       map[string][]Card `yaml:"Lists"`
}

// LoadBoardTemplateFromJSONFile loads a JSON board
func LoadBoardTemplateFromJSONFile(path string) (t *Board, err error) {
	var contentBytes []byte
	t = new(Board)
	contentBytes, err = ioutil.ReadFile(path)
	if err != nil {
		return
	}
	err = json.Unmarshal(contentBytes, t)
	return
}

// LoadBoardTemplateFromYAMLFile loads a YAML board
func LoadBoardTemplateFromYAMLFile(path string) (t *Board, err error) {
	var contentBytes []byte
	t = new(Board)
	contentBytes, err = ioutil.ReadFile(path)
	if err != nil {
		return
	}
	err = yaml.Unmarshal(contentBytes, t)
	return
}
