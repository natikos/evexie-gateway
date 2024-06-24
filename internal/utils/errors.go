package utils

import "api-gateway/internal/logger"

func FailOnError(err error, message string) {
	if err != nil {
		logger.Logger.Fatalln(message)
	}
}
