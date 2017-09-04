package client

import (
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

// Initialize creates the Trello board
func (c *TrelloClient) Initialize(cmd *cobra.Command) error {
	c.AppKey, _ = cmd.Flags().GetString("appkey")
	c.Token, _ = cmd.Flags().GetString("token")
	return nil
}

// Apply creates the pre-populated Trello board
func (c *TrelloClient) Apply(t *tpl.Board) error {
	var board trello.Board
	var labels []*trello.Label
	c.client = trello.NewClient(c.AppKey, c.Token)
	if err := c.client.Post("boards", map[string]string{
		"name":         t.Name,
		"desc":         t.Description,
		"defaultLists": "false"}, &board); err != nil {
		return err
	}
	log.Printf("Created board %s", t.Name)
	for _, lbl := range t.Labels {
		var label trello.Label
		if err := c.client.Post("labels", map[string]string{"name": lbl.Name, "idBoard": board.ID, "color": lbl.Color}, &label); err != nil {
			log.Fatal(err)
		}
		labels = append(labels, &label)
		log.Printf("Created label %s", lbl.Name)
	}
	for laneName, laneStories := range t.Lists {
		var lane trello.List
		c.client.Post("lists", map[string]string{"name": laneName, "idBoard": board.ID, "pos": "bottom"}, &lane)
		log.Printf("Created list %s", laneName)

		for _, story := range laneStories {
			// var wantLabels []*trello.Label

			// for _, l := range labels {
			// 	for _, wantLbl := range story.Labels {
			// 		if l.Name == wantLbl {
			// 			wantLabels = append(wantLabels, l)
			// 		}
			// 	}
			// }

			var wantLabels []string
			for _, l := range labels {
				for _, wantLbl := range story.Labels {
					if l.Name == wantLbl {
						wantLabels = append(wantLabels, l.ID)
					}
				}
			}

			// c.client.CreateCard(&trello.Card{
			// 	Name:   story.Name,
			// 	Desc:   story.Description,
			// 	IDList: lane.ID,
			// 	Labels: wantLabels}, map[string]string{"pos": "bottom"})

			var thisCard trello.Card

			c.client.Post("cards", map[string]string{
				"idList":   lane.ID,
				"name":     story.Name,
				"desc":     story.Description,
				"idLabels": strings.Join(wantLabels, ","),
				"pos":      "bottom"}, &thisCard)
			log.Printf("Created card %s", story.Name)
			// log.Printf("%+v", thisCard.Labels)
		}

	}
	return nil
}
