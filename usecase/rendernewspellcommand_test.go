package usecase

import (
	"fmt"
	"testing"
)

func Test_addCommandConfig(t *testing.T) {
	commandName := "my-command"
	currentCode := `
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
				LongDescription: "Builds the config necessary for adding this plugin to the Golang Spell tool.
This command must be available in all Golang Spell plugins to make it possible the plugin addition to the platform.

Syntax: 
golangspell build-config
",
			},
			"init": &domain.Command{
				Name:             "init",
				ShortDescription: "The init command creates a new Golang application using the Golangspell base structure",
				LongDescription: "The init command creates a new Golang application using the Golangspell base structure
The Architectural Model is based in the Clean Architecture and is the basis to add more resources like domain models and repositories.
Args:
module: Module name (required) to initialize with 'Go Modules'. Example: mydomain.com/myapplication"
appname: App name (required) to initialize with 'Go Modules'. Example: myapplication

Syntax: 
golangspell init [module] [appname]
",
				ValidArgs: []string{"module", "name"},
			},
			"initspell": &domain.Command{
				Name:             "initspell",
				ShortDescription: "The initspell command creates a new Golang Spell using the Golangspell base structure",
				LongDescription: "The initspell command creates a new Golang Spell using the Golangspell base structure
A hello example command is included in the new Spell.
Args:
module: Module name (required) to initialize with 'Go Modules'. Example: mydomain.com/myapplication"
newspellname: App name (required) to initialize with 'Go Modules'. Example: my-new-spell
Use the Spell https://github.com/danilovalente/golangspell-core as a reference for templating

Syntax: 
golangspell initspell [module] [newspellname]
",
				ValidArgs: []string{"module", "newspellname"},
			},
			"addspellcommand": &domain.Command{
				Name:             "addspellcommand",
				ShortDescription: "The addspellcommand command creates a new command to the current Golangspell",
				LongDescription: "The addspellcommand command creates a new command to the current Golangspell
Args:
newSpellCommandName: New Spell's Command name (required). Example: my-spell-init
newCommandArgsNames: List of argument names for the command (separated by space)
		
Syntax: 
golangspell addspellcommand [NewSpellCommandName] [...argNames - optional]

Examples:
# Create a new Spell Command called my-spell-init (all commands must be preceeded by your spell name in order to avoid command name colision)
golangspell addspellcommand my-spell-init

# Create a new Spell Command called my-spell-init expecting the arguments author, createrepository and createreadme (all commands must be preceeded by your spell name in order to avoid command name colision)
golangspell addspellcommand my-spell-init author createrepository createreadme",
				ValidArgs: []string{"newSpellCommandName", "newCommandArgsNames"},
			},
		},
	}
}`

	type args struct {
		code string
		args []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Add Component and Delete",
			args: args{
				code: currentCode,
				args: []string{commandName}},
			want: fmt.Sprintf(`
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
				LongDescription: "Builds the config necessary for adding this plugin to the Golang Spell tool.
This command must be available in all Golang Spell plugins to make it possible the plugin addition to the platform.

Syntax: 
golangspell build-config
",
			},
			"init": &domain.Command{
				Name:             "init",
				ShortDescription: "The init command creates a new Golang application using the Golangspell base structure",
				LongDescription: "The init command creates a new Golang application using the Golangspell base structure
The Architectural Model is based in the Clean Architecture and is the basis to add more resources like domain models and repositories.
Args:
module: Module name (required) to initialize with 'Go Modules'. Example: mydomain.com/myapplication"
appname: App name (required) to initialize with 'Go Modules'. Example: myapplication

Syntax: 
golangspell init [module] [appname]
",
				ValidArgs: []string{"module", "name"},
			},
			"initspell": &domain.Command{
				Name:             "initspell",
				ShortDescription: "The initspell command creates a new Golang Spell using the Golangspell base structure",
				LongDescription: "The initspell command creates a new Golang Spell using the Golangspell base structure
A hello example command is included in the new Spell.
Args:
module: Module name (required) to initialize with 'Go Modules'. Example: mydomain.com/myapplication"
newspellname: App name (required) to initialize with 'Go Modules'. Example: my-new-spell
Use the Spell https://github.com/danilovalente/golangspell-core as a reference for templating

Syntax: 
golangspell initspell [module] [newspellname]
",
				ValidArgs: []string{"module", "newspellname"},
			},
			"addspellcommand": &domain.Command{
				Name:             "addspellcommand",
				ShortDescription: "The addspellcommand command creates a new command to the current Golangspell",
				LongDescription: "The addspellcommand command creates a new command to the current Golangspell
Args:
newSpellCommandName: New Spell's Command name (required). Example: my-spell-init
newCommandArgsNames: List of argument names for the command (separated by space)
		
Syntax: 
golangspell addspellcommand [NewSpellCommandName] [...argNames - optional]

Examples:
# Create a new Spell Command called my-spell-init (all commands must be preceeded by your spell name in order to avoid command name colision)
golangspell addspellcommand my-spell-init

# Create a new Spell Command called my-spell-init expecting the arguments author, createrepository and createreadme (all commands must be preceeded by your spell name in order to avoid command name colision)
golangspell addspellcommand my-spell-init author createrepository createreadme",
				ValidArgs: []string{"newSpellCommandName", "newCommandArgsNames"},
			},
			"%s": &domain.Command{
				Name:             "%s",
				ShortDescription: "The %s [TODO: PUT HERE THE NEW COMMAND FEATURE]",
				LongDescription: The %s [TODO: PUT HERE THE NEW COMMAND FEATURE EXTENDED DESCRIPTION]
Args:
[TODO: PUT HERE THE NEW COMMAND ARGS DESCRIPTION]

Syntax: 
golangspell [TODO: PUT HERE THE NEW COMMAND SYNTAX]

Examples:
[TODO: PUT HERE THE NEW COMMAND EXAMPLES IF NEEDED],
		%s
	},
},
}
}`, commandName, commandName, commandName, commandName, ""),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addCommandConfig(tt.args.code, tt.args.args); got != tt.want {
				t.Errorf("addCommandConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
