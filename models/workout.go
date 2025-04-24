package models

import "time"

type Workout struct {
	ID              int              `grom:"primaryKey"`
	Order           int              `gorm:"unique;not null"`
	Name            string           `gorm:"unique;not null"`
	WorkoutSessions []WorkoutSession `gorm:"foreignKey:WorkoutID"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
