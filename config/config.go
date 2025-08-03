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
	vper.SetDefault("retry.times", 3)
	vper.SetDefault("retry.delay", 5)
	vper.SetDefault("log.level", "info")
	vper.SetDefault("ai.timeout", 10)
	vper.SetConfigName(CfgFileName)
	vper.SetConfigType(CfgFileType)
	vper.AddConfigPath(".")
	vper.AddConfigPath("./data")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("%s\n", err.Error())
	}
	err = viper.Unmarshal(&Cfg)
	if err != nil {
		log.Fatalf("%s\n", err.Error())
	}
}
