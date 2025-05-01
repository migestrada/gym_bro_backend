package models

import "time"

type Workout struct {
	ID               uint              `gorm:"primaryKey"`
	Order            int               `gorm:"not null;uniqueIndex:idx_training_plan_order,composite:order_plan"`
	Name             string            `gorm:"not null"`
	WorkoutSessions  []WorkoutSession  `gorm:"foreignKey:WorkoutID"`
	TrainingPlanID   uint              `gorm:"not null;uniqueIndex:idx_training_plan_order,composite:order_plan"`
	TrainingPlan     TrainingPlan      `gorm:"foreignKey:TrainingPlanID"`
	WorkoutExercises []WorkoutExercise `gorm:"foreignKey:WorkoutID"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
