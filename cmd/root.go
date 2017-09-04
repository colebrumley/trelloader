package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/colebrumley/trelloader/client"
	"github.com/colebrumley/trelloader/tpl"
	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "trelloader",
	Short: "Create a Trello board from JSON",
	Long:  "This utility creates a Trello board pre-populated with lists, cards, and labels from a JSON template.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client := new(client.TrelloClient)
		client.Initialize(cmd)
		for _, cfgFile := range args {
			template, err := tpl.LoadBoardTemplateFromFile(cfgFile)
			if err != nil {
				panic(err)
			}
			log.Println("Loaded file", cfgFile)
			if err = client.Apply(template); err != nil {
				panic(err)
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
