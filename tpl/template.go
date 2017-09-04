package tpl

import (
	"encoding/json"
	"io/ioutil"
)

// Card represents a single work item
type Card struct {
	Name, Description string
	Labels            []string
}

// Label represents a Trello label
type Label struct {
	Name, Color, ID string
}

// Board represents a board
type Board struct {
	Name, Description, Background string
	Labels                        []Label
	Lists                         map[string][]Card
}

// LoadBoardTemplateFromFile loads a JSON board
func LoadBoardTemplateFromFile(path string) (t *Board, err error) {
	var contentBytes []byte
	t = new(Board)
	contentBytes, err = ioutil.ReadFile(path)
	if err != nil {
		return
	}
	err = json.Unmarshal(contentBytes, t)
	return
}
