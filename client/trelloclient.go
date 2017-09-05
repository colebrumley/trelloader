package client

import (
	"fmt"
	"log"
	"strings"

	"github.com/adlio/trello"
	"github.com/colebrumley/trelloader/tpl"
	"github.com/spf13/cobra"
)

// TrelloClient is a Client for Trello ;)
type TrelloClient struct {
	AppKey, Token string
	client        *trello.Client
}

// Initialize loads the flag values
func (c *TrelloClient) Initialize(cmd *cobra.Command) error {
	var err error
	if c.AppKey, err = cmd.Flags().GetString("appkey"); err != nil {
		return fmt.Errorf("must supply a valid appkey")
	}
	if c.Token, err = cmd.Flags().GetString("token"); err != nil {
		return fmt.Errorf("must supply a valid token")
	}
	return nil
}

// Apply creates the pre-populated Trello board
func (c *TrelloClient) Apply(b *tpl.Board) error {
	var (
		thisBoard trello.Board
		allLabels []*trello.Label
	)

	c.client = trello.NewClient(c.AppKey, c.Token)

	boardArgs, err := validateBoardData(b)
	if err != nil {
		return err
	}

	if err := c.client.Post("boards", boardArgs, &thisBoard); err != nil {
		return err
	}
	log.Printf("Created board %s", b.Name)

	for _, lbl := range b.Labels {
		var label trello.Label
		if len(lbl.Name) < 1 {
			return fmt.Errorf("label name is required")
		}
		if err := c.client.Post("labels", map[string]string{
			"name":    lbl.Name,
			"idBoard": thisBoard.ID,
			"color":   lbl.Color}, &label); err != nil {
			return err
		}
		allLabels = append(allLabels, &label)
		log.Printf("Created label %s", lbl.Name)
	}

	for listName, listCards := range b.Lists {
		var thisList trello.List
		if len(listName) < 1 {
			return fmt.Errorf("list names are required")
		}
		if err := c.client.Post("lists", map[string]string{
			"name":    listName,
			"idBoard": thisBoard.ID,
			"pos":     "bottom"}, &thisList); err != nil {
			return err
		}
		log.Printf("Created list %s", listName)

		for _, card := range listCards {
			var (
				wantLabels []string
				thisCard   trello.Card
			)

			for _, l := range allLabels {
				for _, wantLbl := range card.Labels {
					if l.Name == wantLbl {
						wantLabels = append(wantLabels, l.ID)
					}
				}
			}

			if len(card.Name) < 1 {
				return fmt.Errorf("card names are required")
			}

			cardData := map[string]string{
				"idList": thisList.ID,
				"name":   card.Name,
				"pos":    "bottom",
			}

			if len(card.Description) > 0 {
				cardData["desc"] = card.Description
			}

			if len(wantLabels) > 0 {
				cardData["idLabels"] = strings.Join(wantLabels, ",")
			}

			if err := c.client.Post("cards", cardData, &thisCard); err != nil {
				return err
			}
			log.Printf("Created card %s", card.Name)
		}

	}
	return nil
}

func validateBoardData(b *tpl.Board) (boardData map[string]string, err error) {
	err = nil
	boardData = map[string]string{
		"defaultLists":  "false",
		"defaultLabels": "false",
		"prefs_voting":  "members",
	}

	if len(b.Name) < 1 {
		err = fmt.Errorf("board name is required")
		return
	}

	boardData["name"] = b.Name

	if len(b.Description) > 0 {
		boardData["desc"] = b.Description
	}

	if len(b.Background) > 0 {
		boardData["prefs_background"] = b.Background
	}

	return
}
