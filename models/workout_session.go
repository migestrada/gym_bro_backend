package models

import (
	"time"
)

type WorkoutSession struct {
	ID                uint
	Date              time.Time
	WorkoutExerciseID uint
	WorkoutExercise   WorkoutExercise
	SetID             uint
	Set               Set
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
