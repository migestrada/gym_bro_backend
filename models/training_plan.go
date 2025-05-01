package models

import "time"

type TrainingPlan struct {
	ID          int
	Name        string `gorm:"unique;not null"`
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Workouts    []Workout `gorm:"foreignKey:TrainingPlanID"`
}
