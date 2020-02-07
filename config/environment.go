package config

import (
	"fmt"

	coreconfig "github.com/danilovalente/golangspell/config"
	"github.com/danilovalente/golangspell/domain"
	"github.com/spf13/viper"
)

const (
	//configFileName defines the configuration file name
	configFileName = ".golangspell.json"
)

//Values stores the current configuration values
var (
	Values domain.Config
)

//ConfigFilePath contains the path of the config file
var ConfigFilePath = fmt.Sprintf("%s/%s", coreconfig.GetGolangspellHome(), configFileName)

func init() {
	viper.BindEnv("TestRun", "TESTRUN")
	viper.SetDefault("TestRun", false)
	viper.BindEnv("LogLevel", "LOG_LEVEL")
	viper.SetDefault("LogLevel", "INFO")
	viper.BindEnv("GoPath", "GOPATH")
	viper.SetDefault("GoPath", fmt.Sprintf("%s/go", coreconfig.GetHomeDir()))
	viper.Unmarshal(&Values)
}
