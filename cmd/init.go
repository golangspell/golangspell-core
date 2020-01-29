package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	RunCommandFunctions["init"] = runInitCommand
}

//Module name to initialize with 'Go Modules'
var Module string

//AppName used to define the application's directory and the default value to the config variable with the same name
var AppName string

func runInitCommand(cmd *cobra.Command, args []string) {
	fmt.Printf("Args: %s\n", args)
}
