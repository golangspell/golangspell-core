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
		Commands: []domain.Command{
			domain.Command{
				Name: "init",
				Flags: []domain.Flag{
					domain.Flag{
						Name:      "module",
						Shorthand: "m",
						Value:     "",
						Usage:     "Module name (required) to initialize with 'Go Modules'. Example: mydomain.com/myapplication",
					},
					domain.Flag{
						Name:      "appname",
						Shorthand: "n",
						Value:     "",
						Usage:     "App name (required) to initialize with 'Go Modules'. Example: myapplication",
					},
				},
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
