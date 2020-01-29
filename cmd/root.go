package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile     string
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
	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	viper.SetDefault("license", "Apache")
}

func initConfig() {
	viper.AutomaticEnv()
}
