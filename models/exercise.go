package models

type Exercise struct {
	ID              uint   `gorm:"primaryKey"`
	Name            string `grom:"unique;not null"`
	Description     string
	PhotoURL        string `json:"photo_url"`
	VideoURL        string `json:"video_url"`
	CreatedAt       int64
	UpdatedAt       int64
	WorkoutSessions []WorkoutSession `gorm:"foreignKey:ExerciseID"`
	Sets            []Set            `gorm:"foreignKey:ExerciseID"`
}
