package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Workout struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type WorkoutRepository struct {
	db *sql.DB
}

func NewWorkoutRepository(db *sql.DB) *WorkoutRepository {
	return &WorkoutRepository{db: db}
}

func (repo *WorkoutRepository) GetWorkouts(ctx context.Context, userId uuid.UUID) ([]Workout, error) {
	workouts := make([]Workout, 0)
	rows, err := repo.db.QueryContext(ctx, "SELECT id from workouts where user_id = $1", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		workout := new(Workout)
		if err := rows.Scan(&workout.ID); err != nil {
			return nil, err
		}
		workouts = append(workouts, *workout)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return workouts, nil
}
