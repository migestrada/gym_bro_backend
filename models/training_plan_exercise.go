package models

type TrainingPlanExercise struct {
	ID             uint `gorm:"primaryKey"`
	Order          int  `gorm:"not null;unique"`
	TrainingPlanID uint `gorm:"not null"`
	TrainingPlan   TrainingPlan
	ExerciseID     uint `gorm:"not null"`
	Exercise       Exercise
}
