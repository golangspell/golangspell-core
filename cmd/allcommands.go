package cmd

import (
	"github.com/danilovalente/golangspell-core/appcontext"
	"github.com/danilovalente/golangspell/domain"
)

func init() {
	appcontext.Current.Add(appcontext.Spell, GetSpellConfig)
}

//GetSpellConfig lazy loads a Spell Config
func GetSpellConfig() appcontext.Component {
	return buildSpellConfig()
}

//RunCommandFunctions stores the available RunCommandFunctions in the Spell, to correlate with the Spell Commands using the Name as key
var RunCommandFunctions map[string]domain.RunCommandFunction = make(map[string]domain.RunCommandFunction)

func addInnerCommands() {
	rootCmd.AddCommand(versionCmd)
	spell := appcontext.Current.Get(appcontext.Spell).(domain.Spell)
	for key, command := range spell.Commands {
		rootCmd.AddCommand(command.CobraCommand(RunCommandFunctions[key]))
	}
}
