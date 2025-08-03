package logger

import (
	"github.com/assimon/ai-anti-bot/config"
	"github.com/sirupsen/logrus"
	"os"
)

var Log = logrus.New()

func init() {
	Log.Out = os.Stdout
	level, err := logrus.ParseLevel(config.Cfg.Log.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	Log.SetLevel(level)
}
