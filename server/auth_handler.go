package server

import (
	"api-gateway/internal/logger"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

type AuthHandler struct{}

func (h *AuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	command := vars["action"]
	logger.Logger.Infoln(command)
	switch r.Method {
	case http.MethodPost:
		h.handlePost(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	// action := strings.ToLower(vars["action"])
}

func (h *AuthHandler) handlePost(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}

}
