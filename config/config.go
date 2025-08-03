package config

import (
	"github.com/spf13/viper"
	"log"
)

const (
	CfgFileName = "config.yml"
	CfgFileType = "yml"
)

var Cfg *Config

type Config struct {
	Retry  RetryConfig
	Log    LogConfig
	Ai     AiConfig
	Prompt PromptConfig
}

type RetryConfig struct {
	Times int
	Delay int
}

type LogConfig struct {
	Level string
}

type AiConfig struct {
	Timeout int
}

type PromptConfig struct {
	Text  string
	Image string
}

func init() {
	viper.SetDefault("retry.times", 3)
	viper.SetDefault("retry.delay", 5)
	viper.SetDefault("log.level", "info")
	viper.SetDefault("ai.timeout", 10)
	viper.SetConfigName(CfgFileName)
	viper.SetConfigType(CfgFileType)
	viper.AddConfigPath(".")
	viper.AddConfigPath("./data")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("%s\n", err.Error())
	}
	err = viper.Unmarshal(&Cfg)
	if err != nil {
		log.Fatalf("%s\n", err.Error())
	}
}
