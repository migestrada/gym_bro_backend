package controllers

import (
	"gym-bro-backend/connection"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TrainingPlanExercise struct {
	ID             uint `gorm:"primaryKey"`
	Order          int  `json:"order" binding:"required"`
	TrainingPlanID uint `json:"training_plan_id" binding:"required"`
	ExerciseID     uint `json:"exercise_id" binding:"required"`
}

func GetTrainingPlanExercises(context *gin.Context) {
	var trainingPlanExercises []TrainingPlanExercise

	if err := connection.DB.Find(&trainingPlanExercises).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to retrieve training plan exercises",
			"error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Training plan exercises retrieved successfully",
		"data":    trainingPlanExercises,
	})
}
