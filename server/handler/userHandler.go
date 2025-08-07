package handler

import (
	"encoding/json"
	"net/http"

	"github.com/mitchan/gymbro/service"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	// TODO:
	type responsePayload struct {
		Success bool `json:"success"`
	}

	payload := responsePayload{Success: true}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payload)
}
