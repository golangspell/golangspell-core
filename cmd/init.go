package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

//Module name to initialize with 'Go Modules'
var Module string

//AppName used to define the application's directory and the default value to the config variable with the same name
var AppName string

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initializes a base Golang Spell project structure",
	Long:  `initializes the current directory with a base Golang Spell project structure`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Golang Spell v1.0.0 -- HEAD")
	},
	ValidArgs: []string{"module", "appname"},
}
