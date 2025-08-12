package router

import (
	"log"
	"net/http"
	"time"

	"github.com/mitchan/gymbro/handler"
	"github.com/mitchan/gymbro/middleware"
)

func TestMiddlware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Println(r.Method, r.URL.Path, time.Since(start))
	})
}

func NewRouter(userHandler *handler.UserHandler) *http.ServeMux {
	router := http.NewServeMux()

	// user
	router.HandleFunc("POST /api/user", userHandler.CreateUser)
	router.HandleFunc("POST /api/user/login", userHandler.Login)

	// auth router
	authRouter := http.NewServeMux()
	authRouter.HandleFunc("GET /api/user/me", userHandler.Me)

	router.Handle("/", middleware.JwtAuth(authRouter))

	// TEST: route with parameter (to be removed)
	// router.HandleFunc("GET /api/user/{id}", func(w http.ResponseWriter, r *http.Request) {
	// 	id := r.PathValue("id")
	// 	w.Write([]byte("received request for id: " + id))
	// })

	return router
}
