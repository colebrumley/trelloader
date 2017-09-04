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
func (c *TrelloClient) Apply(t *tpl.Board) error {
	var (
		thisBoard trello.Board
		allLabels []*trello.Label
	)

	c.client = trello.NewClient(c.AppKey, c.Token)

	boardArgs := map[string]string{
		"name":         t.Name,
		"desc":         t.Description,
		"defaultLists": "false",
	}

	if len(t.Background) > 0 {
		boardArgs["prefs_background"] = t.Background
	}

	if err := c.client.Post("boards", boardArgs, &thisBoard); err != nil {
		return err
	}
	log.Printf("Created board %s", t.Name)

	for _, lbl := range t.Labels {
		var label trello.Label
		if err := c.client.Post("labels", map[string]string{
			"name":    lbl.Name,
			"idBoard": thisBoard.ID,
			"color":   lbl.Color}, &label); err != nil {
			log.Fatal(err)
		}
		allLabels = append(allLabels, &label)
		log.Printf("Created label %s", lbl.Name)
	}

	for laneName, laneStories := range t.Lists {
		var lane trello.List
		if err := c.client.Post("lists", map[string]string{
			"name":    laneName,
			"idBoard": thisBoard.ID,
			"pos":     "bottom"}, &lane); err != nil {
			return err
		}
		log.Printf("Created list %s", laneName)

		for _, story := range laneStories {
			var (
				wantLabels []string
				thisCard   trello.Card
			)

			for _, l := range allLabels {
				for _, wantLbl := range story.Labels {
					if l.Name == wantLbl {
						wantLabels = append(wantLabels, l.ID)
					}
				}
			}

			if err := c.client.Post("cards", map[string]string{
				"idList":   lane.ID,
				"name":     story.Name,
				"desc":     story.Description,
				"idLabels": strings.Join(wantLabels, ","),
				"pos":      "bottom"}, &thisCard); err != nil {
				return err
			}
			log.Printf("Created card %s", story.Name)
		}

	}
	return nil
}
