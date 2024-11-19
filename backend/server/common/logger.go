package common

import (
	"os"
	"time"

	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func init() {
	Logger = logrus.New()

	Logger.SetOutput(&lumberjack.Logger{
		Filename:   "./logs/ai-saas-schedular.log",
		MaxSize:    10,
		MaxAge:     2, 
		MaxBackups: 7,
		Compress:   true,
	})

	Logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
	})

	level, err := logrus.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		level = logrus.InfoLevel
	}
	Logger.SetLevel(level)
}
