package cmd

import (
	"fmt"

	"encoding/json"

	"github.com/danilovalente/golangspell/domain"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
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
	spell := domain.Spell{}
	spellYaml := `
		Name: golangspell-core
		URL:  "https://github.com/danilovalente/golangspell-core"
		Commands: 
			- build-config: 
				Name: build-config
				ShortDescription: Builds the config necessary for adding this plugin to the Golang Spell tool
				LongDescription: "Builds the config necessary for adding this plugin to the Golang Spell tool.\n
This command must be available in all Golang Spell plugins to make it possible the plugin addition to the platform.\n
\n
Syntax:\n 
golangspell build-config
"
			- init: 
				Name: init
				ShortDescription: The init command creates a new Golang application using the Golangspell base structure
				LongDescription: "The init command creates a new Golang application using the Golangspell base structure\n
The Architectural Model is based in the Clean Architecture and is the basis to add more resources like domain models and repositories.\n
Args:\n
module: Module name (required) to initialize with 'Go Modules'. Example: mydomain.com/myapplication\n
appname: App name (required) to initialize with 'Go Modules'. Example: myapplication\n
\n
Syntax:\n
golangspell init [module] [appname]\n
"
				ValidArgs: 
					- module
					- name
			- initspell: 
				Name: initspell
				ShortDescription: The initspell command creates a new Golang Spell using the Golangspell base structure
				LongDescription: "The initspell command creates a new Golang Spell using the Golangspell base structure\n
A hello example command is included in the new Spell.\n
Args:\n
module: Module name (required) to initialize with 'Go Modules'. Example: mydomain.com/myapplication\n
newspellname: App name (required) to initialize with 'Go Modules'. Example: my-new-spell\n
Use the Spell https://github.com/danilovalente/golangspell-core as a reference for templating\n
\n
Syntax:\n
golangspell initspell [module] [newspellname]\n
"
				ValidArgs: 
					- module
					- newspellname
			- addspellcommand: 
				Name: addspellcommand
				ShortDescription: The addspellcommand command adds a new command to the current Golang Spell (the one located in the current directory)
				LongDescription: "The addspellcommand command adds a new command to the current Golang Spell (the one located in the current directory)\n
Args:\n
NewSpellCommandName: New Spell's Command name (required). Example: my-spell-init\n
NewCommandArgsNames: List of argument names for the command (separated by space)\n
\n		
Syntax:\n
golangspell addspellcommand [NewSpellCommandName] [...argNames - optional]\n
\n
Examples:\n
# Create a new Spell Command called my-spell-init (all commands must be preceeded by your spell name in order to avoid command name colision)\n
golangspell addspellcommand my-spell-init\n
\n
# Create a new Spell Command called my-spell-init expecting the arguments author, createrepository and createreadme (all commands must be preceeded by your spell name in order to avoid command name colision)\n
golangspell addspellcommand my-spell-init author createrepository createreadme
"
				ValidArgs: 
					- NewSpellCommandName
					- NewCommandArgsNames
`
	err := yaml.Unmarshal([]byte(spellYaml), &spell)
	if err != nil {
		fmt.Println("An error occurred while trying to unmarshal the Spell Yaml specification. Message: " + err.Error())
		panic(err)
	}

	return spell
}
