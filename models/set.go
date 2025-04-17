package models

import "time"

type Set struct {
	ID          uint
	Quantity    int
	Repetitions int
	ExerciseID  uint
	Exercise    Exercise
	WorkoutID   uint
	Workout     Workout
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
