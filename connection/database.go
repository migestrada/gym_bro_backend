package connection

import (
	"gym-bro-backend/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func CreateConnection() {
	dsn := "host=localhost user=go_user password=go_user dbname=go_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("DB connection failed:", err)
	}

	modelsToMigrate := []interface{}{
		&models.TrainingPlan{},
		&models.Workout{},
		&models.Exercise{},
		&models.Set{},
		&models.WorkoutSession{},
		&models.TrainingPlanExercise{},
	}

	for _, model := range modelsToMigrate {
		err = db.AutoMigrate(model)
		if err != nil {
			log.Fatal("Error during migration:", err)
		}
	}

	DB = db
}
