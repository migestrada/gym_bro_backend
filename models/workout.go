package models

import "time"

type Workout struct {
	ID                   int
	Index                int
	Name                 string
	TrainingPlanWorkouts []TrainingPlanWorkout `gorm:"foreignKey:WorkoutID"`
	WorkoutSessions      []WorkoutSession      `gorm:"foreignKey:WorkoutID"`
	Sets                 []Set                 `gorm:"foreignKey:WorkoutID"`
	CreatedAt            time.Time
	UpdatedAt            time.Time
}
