package models

import "time"

type TrainingPlan struct {
	ID                   int
	Name                 string
	CreatedAt            time.Time
	UpdatedAt            time.Time
	TrainingPlanWorkouts []TrainingPlanWorkout `gorm:"foreignKey:TrainingPlanID"`
}
