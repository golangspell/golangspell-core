package cmd

import (
	"fmt"

	"encoding/json"

	"github.com/danilovalente/golangspell/domain"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(buildConfigCmd)
}

func buildSpellConfig() domain.Spell {
	return domain.Spell{
		Name: "core",
		URL:  "https://github.com/danilovalente/golangspell-core",
		Commands: map[string]Command{
			"init": Command{
				Name:             "init",
				ShortDescription: "The init command creates a new Golang application using the Golangspell base structure",
				LongDescription: `The init command creates a new Golang application using the Golangspell base structure
The Architectural Model is based in the Clean Architecture and is the basis to add more resources like domain models and repositories.
Args:
module: Module name (required) to initialize with 'Go Modules'. Example: mydomain.com/myapplication"
appname: App name (required) to initialize with 'Go Modules'. Example: myapplication

Syntax: 
golangspell init [module] [appname]
`,
				ValidArgs: []string{"module", "name"},
			},
		},
	}
}

var buildConfigCmd = &cobra.Command{
	Use:   "build-config",
	Short: "Builds the config necessary for adding this plugin to the Golang Spell tool",
	Long: `Builds the config necessary for adding this plugin to the Golang Spell tool.
This command must be available in all Golang Spell plugins to make it possible the plugin addition to the platform.`,
	Run: func(cmd *cobra.Command, args []string) {
		configBytes, err := json.MarshalIndent(buildSpellConfig(), "", "  ")
		if err != nil {
			panic(err)
		}
		fmt.Println(string(configBytes))
	},
}
