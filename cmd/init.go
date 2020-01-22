package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	initCmd.Flags().StringVarP(&Module, "module", "m", "", "Module name (required) to initialize with 'Go Modules'. Example: mydomain.com/myapplication")
	initCmd.MarkFlagRequired("module")
	initCmd.Flags().StringVarP(&AppName, "appname", "n", "", "App name (required) to initialize with 'Go Modules'. Example: myapplication")
	initCmd.MarkFlagRequired("appname")
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
}
