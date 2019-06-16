package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/colebrumley/trelloader/client"
	"github.com/colebrumley/trelloader/tpl"
	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:     "trelloader",
	Short:   "Create a Trello board from JSON or YAML",
	Example: "trelloader -k YOUR_KEY -t YOUR_TOKEN examples/example.json examples/example.yaml",
	Long: `This utility creates a Trello board pre-populated with a background,
lists, cards, and labels from a JSON or YAML template.
	
A Trello API AppKey and token are required, to generate new ones see
https://trello.com/app-key`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client := new(client.TrelloClient)
		if err := client.Initialize(cmd); err != nil {
			log.Fatal(err)
		}
		for _, cfgFile := range args {
			var template *tpl.Board
			var err error
			if strings.HasSuffix(cfgFile, ".yml") || strings.HasSuffix(cfgFile, ".yaml") {
				template, err = tpl.LoadBoardTemplateFromYAMLFile(cfgFile)
			}
			if strings.HasSuffix(cfgFile, ".json") {
				template, err = tpl.LoadBoardTemplateFromJSONFile(cfgFile)
			}
			if err != nil {
				log.Fatalf("failed to load board template %s: %s", cfgFile, err.Error())
			}
			log.Println("Loaded file", cfgFile)
			if err = client.Apply(template); err != nil {
				log.Fatalf("failed to create board %s: %s", template.Name, err.Error())
			}
			log.Printf("Built board %s", template.Name)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	RootCmd.Flags().StringP("appkey", "k", "", "Trello App Key")
	RootCmd.Flags().StringP("token", "t", "", "Trello token")
}
