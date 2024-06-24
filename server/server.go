package server

import (
	"api-gateway/internal/logger"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type GatewayServer struct {
}

func NewGatewayServer() *GatewayServer {
	return &GatewayServer{}
}

func (server *GatewayServer) Run() {
	router := mux.NewRouter()

	router.PathPrefix("/auth/{action}").Handler(NewAuthHandler())

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	logger.Logger.Infoln("Server running on port: ", port)

	err := http.ListenAndServe(":"+port, router)
	logger.Logger.Panicln(err)
}
