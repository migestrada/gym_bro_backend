package controllers

import (
	"gym-bro-backend/connection"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type WorkoutSession struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	Date           time.Time `json:"date" binding:"required"`
	WorkoutID      uint      `json:"workout_id" binding:"required"`
	ExerciseID     uint      `json:"exercise_id" binding:"required"`
	SetID          uint      `json:"set_id" binding:"required"`
	TrainingPlanID uint      `json:"training_plan_id" binding:"required"`
}

func GetWorkoutSessions(context *gin.Context) {
	var workoutSessions []WorkoutSession

	if err := connection.DB.Find(&workoutSessions).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to retrieve workout sessions",
			"error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Workout sessions retrieved successfully",
		"data":    workoutSessions,
	})
}
