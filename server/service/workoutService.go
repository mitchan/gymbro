package service

import (
	"context"
	"log"

	"github.com/mitchan/gymbro/middleware"
	"github.com/mitchan/gymbro/model"
	"github.com/mitchan/gymbro/repository"
)

type WorkoutService struct {
	workoutRepo *repository.WorkoutRepository
}

func NewWorkoutService(workoutRepo *repository.WorkoutRepository) *WorkoutService {
	return &WorkoutService{
		workoutRepo: workoutRepo,
	}
}

func (ws *WorkoutService) GetWorkouts(ctx context.Context) ([]repository.Workout, error) {
	userID, err := middleware.GetUserIDFromContext(ctx)
	if err != nil {
		log.Printf("WorkoutService.GetWorkouts - cannot retrieve AuthUser from context")
		return nil, model.UnauthedError
	}

	workouts, err := ws.workoutRepo.GetWorkouts(ctx, userID)
	if err != nil {
		log.Printf("WorkoutService.GetWorkouts - cannot retrieve workouts from userID: %s", err)
	}

	return workouts, nil
}
