package router

import (
	"net/http"

	"github.com/mitchan/gymbro/handler"
)

func SetupRouter(userHandler *handler.UserHandler) {
	// user
	http.HandleFunc("POST /api/user", userHandler.CreateUser)
}
