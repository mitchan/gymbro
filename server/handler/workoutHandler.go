package handler

import (
	"net/http"

	"github.com/mitchan/gymbro/service"
	"github.com/mitchan/gymbro/util"
)

type WorkoutHandler struct {
	workoutService *service.WorkoutService
}

func NewWorkoutHandler(workoutService *service.WorkoutService) *WorkoutHandler {
	return &WorkoutHandler{
		workoutService: workoutService,
	}
}

func (wh *WorkoutHandler) GetWorkouts(w http.ResponseWriter, r *http.Request) {
	workouts, err := wh.workoutService.GetWorkouts(r.Context())
	if err != nil {
		util.WriteError(w, http.StatusInternalServerError, "cannot retrieve user")
		return
	}

	util.WriteJSON(w, http.StatusOK, workouts)
}
