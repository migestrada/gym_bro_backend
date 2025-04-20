package models

import (
	"time"
)

type WorkoutSession struct {
	ID             uint
	Date           time.Time
	WorkoutID      uint
	Workout        Workout
	ExerciseID     uint
	Exercise       Exercise
	SetID          uint
	Set            Set
	TrainingPlanID uint
	TrainingPlan   TrainingPlan
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
