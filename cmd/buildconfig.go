package cmd

import (
	"fmt"

	"encoding/json"

	"github.com/danilovalente/golangspell/domain"
	"github.com/spf13/cobra"
)

func init() {
	RunCommandFunctions["build-config"] = runBuildConfigCommand
}

func runBuildConfigCommand(cmd *cobra.Command, args []string) {
	configBytes, err := json.MarshalIndent(buildSpellConfig(), "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(configBytes))
}

func buildSpellConfig() domain.Spell {
	return domain.Spell{
		Name: "golangspell-core",
		URL:  "https://github.com/danilovalente/golangspell-core",
		Commands: map[string]*domain.Command{
			"build-config": &domain.Command{
				Name:             "build-config",
				ShortDescription: "Builds the config necessary for adding this plugin to the Golang Spell tool",
				LongDescription: `Builds the config necessary for adding this plugin to the Golang Spell tool.
This command must be available in all Golang Spell plugins to make it possible the plugin addition to the platform.

Syntax: 
golangspell build-config
`,
			},
			"init": &domain.Command{
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
