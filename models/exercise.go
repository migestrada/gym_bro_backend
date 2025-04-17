package models

type Exercise struct {
	ID              uint
	Name            string
	PhotoURL        string
	WorkoutSessions []WorkoutSession `gorm:"foreignKey:ExerciseID"`
	Sets            []Set            `gorm:"foreignKey:ExerciseID"`
}
