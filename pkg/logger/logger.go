package logger

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

var Log = logrus.New()

func init() {
	Log.Out = os.Stdout
	level, err := logrus.ParseLevel(viper.GetString("log.level"))
	if err != nil {
		level = logrus.InfoLevel
	}
	Log.SetLevel(level)
}
