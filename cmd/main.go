package main

import (
	"api-gateway/internal/broker"
	"api-gateway/internal/logger"
	"api-gateway/server"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		logger.Logger.DPanicln("Cannot load .env file")
		return
	}

	logger.InitLogger()

	brokerInstance, err := broker.GetBrokerInstance()

	if err != nil {
		logger.Logger.Fatalln("Failed to initialize broker:", err)
		return
	}

	defer brokerInstance.Close()

	server.NewGatewayServer().Run()
}
