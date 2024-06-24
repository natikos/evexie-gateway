package server

import (
	"api-gateway/internal/broker"
	"api-gateway/internal/logger"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

type AuthHandler struct {
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (h *AuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	command := vars["action"]
	logger.Logger.Infoln(command)
	switch r.Method {
	case http.MethodPost:
		h.handlePost(w, r, command)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *AuthHandler) handlePost(w http.ResponseWriter, r *http.Request, command string) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}

	message := map[string]interface{}{
		"pattern": command,
		"data":    json.RawMessage(body),
	}

	messageBytes, err := json.Marshal(message)

	if err != nil {
		logger.Logger.Errorln(err)
	}

	broker.Broker.Publish(messageBytes, "auth")
}
