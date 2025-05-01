package models

type WorkoutExercise struct {
	ID         uint `gorm:"primaryKey"`
	WorkoutID  uint `gorm:"not null"`
	Workout    Workout
	ExerciseID uint `gorm:"not null"`
	Exercise   Exercise
}
