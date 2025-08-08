package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mitchan/gymbro/model"
	"github.com/mitchan/gymbro/service"
	"github.com/mitchan/gymbro/util"
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
	util.EnableCors(w)

	type responsePayload struct {
		Success bool `json:"success"`
	}

	var req model.CreateUser
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("CreateUser - JSON decode error: %v", err)
		util.WriteError(w, http.StatusBadRequest, "invalid JSON payload")
		return
	}

	err := h.userService.CreateUser(req)
	if err != nil {
		log.Printf("CreateUser - Service error: %v", err)
		util.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	payload := responsePayload{Success: false}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payload)
}
