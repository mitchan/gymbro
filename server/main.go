package main

import (
	"log"
	"net/http"

	"github.com/mitchan/gymbro/db"
	"github.com/mitchan/gymbro/db/migrations"
	"github.com/mitchan/gymbro/handler"
	"github.com/mitchan/gymbro/middleware"
	"github.com/mitchan/gymbro/repository"
	"github.com/mitchan/gymbro/router"
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
	workoutRepo := repository.NewWorkoutRepository(dbConn)

	// services
	userService := service.NewUserService(userRepo)
	workoutService := service.NewWorkoutService(workoutRepo)

	// handlers
	userHandler := handler.NewUserHandler(userService)
	workoutHandler := handler.NewWorkoutHandler(workoutService)

	// middlewares
	authMiddleware := middleware.NewAuthMiddleware(userRepo)

	r := router.NewRouter(authMiddleware, userHandler, workoutHandler)

	server := http.Server{
		Addr:    ":8080",
		Handler: router.TestMiddlware(r),
	}

	log.Fatal(server.ListenAndServe())
}
