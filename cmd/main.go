package main

import (
	"api-gateway/internal/logger"
	"api-gateway/server"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	logger.InitLogger()

	if err != nil {
		logger.Logger.DPanicln("Cannot load .env file")
	}

	server.NewGatewayServer().Run()
}
