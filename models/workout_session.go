package models

import (
	"time"
)

type WorkoutSession struct {
	ID          uint
	Date        time.Time
	Quantity    int
	Repetitions int
	Weight      float32
	WeightUnit  string
	WorkoutID   uint
	Workout     Workout
	ExerciseID  uint
	Exercise    Exercise
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
