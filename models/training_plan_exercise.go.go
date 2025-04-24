package models

type TrainingPlanExercise struct {
	ID             uint `gorm:"primaryKey"`
	Order          int
	TrainingPlanID uint
	TrainingPlan   TrainingPlan
	ExerciseID     uint
	Exercise       Exercise
}
