package cmd

import (
	"fmt"

	"github.com/danilovalente/golangspell-core/config"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "golangspell-core-version",
	Short: "golangspell-core version number",
	Long:  `Shows the golangspell-core current installed version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("golangspell-core v%s -- HEAD\n", config.Version)
	},
}
