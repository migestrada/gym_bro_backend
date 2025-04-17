package connection

import (
	"log"
	"tgn-backend/models"

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

	err = db.AutoMigrate(
		&models.TrainingPlan{},
		&models.Workout{},
		&models.TrainingPlanWorkout{})

	if err != nil {
		log.Fatal("Error during migration:", err)
	}

	DB = db
}
