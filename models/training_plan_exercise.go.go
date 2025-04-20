package models

type TrainingPlanExercise struct {
	ID             uint `gorm:"primaryKey"`
	Index          int
	TrainingPlanID uint
	TrainingPlan   TrainingPlan
	ExerciseID     uint
	Exercise       Exercise
}
