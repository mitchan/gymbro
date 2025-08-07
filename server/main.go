package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mitchan/gymbro/db"
	"github.com/mitchan/gymbro/db/migrations"
	"github.com/mitchan/gymbro/handler"
	"github.com/mitchan/gymbro/repository"
	"github.com/mitchan/gymbro/service"
)

func main() {
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("Could not initialize DB connection: %s", err)
	}
	defer dbConn.Close()

	if err := dbConn.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
	log.Println("Connected to database successfully")

	// Run migrations
	if err := migrations.RunMigrations(dbConn); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// repos
	userRepo := repository.NewUserRepository(dbConn)

	// services
	userService := service.NewUserService(userRepo)

	// handlers
	userHandler := handler.NewUserHandler(userService)

	http.HandleFunc("POST /api/user", userHandler.CreateUser)

	http.HandleFunc("/api/test", func(w http.ResponseWriter, r *http.Request) {
		type responsePayload struct {
			Success bool `json:"success"`
		}

		payload := responsePayload{Success: true}

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(payload)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
