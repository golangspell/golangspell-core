package cmd

import (
	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	userLicense string

	rootCmd = &cobra.Command{
		Use:   "golangspell-core",
		Short: "Golang Spell - Core: Plugin with the main Golang Spell commands",
		Long: `Golang Spell - Core is the main Plugin for the Golang Spell tool. 
Golang Spell makes it possible to build lightning fast Microservices in Go 
in an easy and productive way.
Welcome to the tool that will kick out the boilerplate code 
and drive you through new amazing possibilities`,
		TraverseChildren: true,
	}
)

// Execute executes the root command.
func Execute() error {
	addInnerCommands()
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringP("author", "a", "", "author name for copyright attribution")
	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "Apache", "name of license for the project")
}

func initConfig() {
}
