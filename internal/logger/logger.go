package logger

import (
	"os"

	"go.uber.org/zap"
)

var Logger *zap.SugaredLogger

func InitLogger() {
	env := os.Getenv("ENV")

	var err error
	var zapLogger *zap.Logger

	if env == "dev" {
		zapLogger, err = zap.NewDevelopment()
	} else {
		zapLogger, err = zap.NewProduction()
	}

	if err != nil {
		panic(err)
	}

	Logger = zapLogger.Sugar()

	defer zapLogger.Sync()
}
