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

	var req model.CreateUser
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("CreateUser - JSON decode error: %v", err)
		util.WriteError(w, http.StatusBadRequest, "invalid JSON payload")
		return
	}

	user, err := h.userService.CreateUser(req)
	if err != nil {
		log.Printf("CreateUser - Service error: %v", err)
		util.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		Value:    user.AccessToken,
		Path:     "/",
		MaxAge:   60 * 60 * 24,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})

	util.WriteJSON(w, http.StatusCreated, user)
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	util.EnableCors(w)

	var req model.LoginUser
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Login - JSON decode error: %v", err)
		util.WriteError(w, http.StatusBadRequest, "invalid JSON payload")
		return
	}

	user, err := h.userService.Login(req)
	if err != nil {
		log.Printf("Login - Service error: %v", err)
		util.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		Value:    user.AccessToken,
		Path:     "/",
		MaxAge:   60 * 60 * 24,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})

	util.WriteJSON(w, http.StatusOK, user)
}
