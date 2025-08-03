package config

import (
	"github.com/spf13/viper"
	"log"
)

const (
	CfgFileName = "config.yml"
	CfgFileType = "yml"
)

func init() {
	viper.SetDefault("retry.times", 3)
	viper.SetDefault("retry.delay", 5)
	viper.SetDefault("log.level", "info")
	viper.SetConfigName(CfgFileName)
	viper.SetConfigType(CfgFileType)
	viper.AddConfigPath(".")
	viper.AddConfigPath("./data")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("%s\n", err.Error())
	}
}
