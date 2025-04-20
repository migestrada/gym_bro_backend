package models

import "time"

type Set struct {
	ID              uint `gorm:"primaryKey"`
	Reps            int
	RestTime        int
	Weight          float32
	WeightUnit      string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	WorkoutSessions []WorkoutSession `gorm:"foreignKey:SetID"`
	ExerciseID      uint
	Exercise        Exercise
}
