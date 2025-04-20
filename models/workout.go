package models

import "time"

type Workout struct {
	ID              int `grom:"primaryKey"`
	Index           int
	Name            string
	WorkoutSessions []WorkoutSession `gorm:"foreignKey:WorkoutID"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
