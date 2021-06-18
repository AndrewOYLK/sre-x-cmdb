package utils

import (
	"os"
	"sync"

	"github.com/sirupsen/logrus"
)

const (
	LOGLEVEL = "debug"
)

var Logger *logrus.Logger

func init() {
	if Logger != nil {
		return
	}

	var once sync.Once
	once.Do(func() {
		Logger = &logrus.Logger{}

		level, err := logrus.ParseLevel(LOGLEVEL)
		if err != nil {
			panic(err)
		}
		Logger.SetLevel(level)
		Logger.SetOutput(os.Stdout)

		formatter := &logrus.JSONFormatter{
			TimestampFormat: "2006/01/02 15:04:05",
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyTime:  "@timestamp",
				logrus.FieldKeyLevel: "@level",
				logrus.FieldKeyMsg:   "@message",
				logrus.FieldKeyFunc:  "@caller",
			},
		}
		Logger.SetFormatter(formatter)
	})
}
