package models

import "time"

type TrainingPlanWorkout struct {
	ID             uint
	CreatedAt      time.Time
	UpdatedAt      time.Time
	TrainingPlanID uint
	WorkoutID      uint
	TrainingPlan   TrainingPlan
	Workout        Workout
}
