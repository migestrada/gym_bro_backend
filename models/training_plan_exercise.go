package models

type TrainingPlanExercise struct {
	ID             uint `gorm:"primaryKey"`
	Order          int  `gorm:"not null;unique"`
	TrainingPlanID uint
	TrainingPlan   TrainingPlan
	ExerciseID     uint
	Exercise       Exercise
}
