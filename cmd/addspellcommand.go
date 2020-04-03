package cmd

import (
	"fmt"

	"github.com/danilovalente/golangspell-core/usecase"
	"github.com/spf13/cobra"
)

func init() {
	RunCommandFunctions["addspellcommand"] = runAddSpellCommandCommand
}

func runAddSpellCommandCommand(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		fmt.Println(`The command addspellcommand have one required parameter: NewSpellCommandName and optional new command's parameter names
Args:
NewSpellCommandName: New Spell's Command name (required). Example: my-spell-init
NewCommandArgsNames: List of argument names for the command (separated by space)
		
Syntax: 
golangspell addspellcommand [NewSpellCommandName] [...argNames - optional]

Examples:
# Create a new Spell Command called my-spell-init (all commands must be preceeded by your spell name in order to avoid command name colision)
golangspell addspellcommand my-spell-init

# Create a new Spell Command called my-spell-init expecting the arguments author, createrepository and createreadme (all commands must be preceeded by your spell name in order to avoid command name colision)
golangspell addspellcommand my-spell-init author createrepository createreadme`)
		return
	}

	err := usecase.RenderNewSpellCommandTemplate(args)
	if err != nil {
		fmt.Printf("An error occurred while trying to create the spell command. Error: %s\n", err.Error())
		return
	}
	fmt.Println("Spell command added!")
}
