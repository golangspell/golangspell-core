package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initializes a base Golang Spell project structure",
	Long:  `initializes the current directory with a base Golang Spell project structure`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Golang Spell v1.0.0 -- HEAD")
	},
}
