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

func GetTrainingPlanExerciseByID(context *gin.Context) {
	var trainingPlanExercise TrainingPlanExercise
	var id string = context.Param("id")

	if err := connection.DB.First(&trainingPlanExercise, id).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Training plan exercise not found",
			"error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Training plan exercise retrieved successfully",
		"data":    trainingPlanExercise,
	})
}

func CreateTrainingPlanExercise(context *gin.Context) {
	var trainingPlanExercise TrainingPlanExercise

	if err := context.ShouldBindJSON(&trainingPlanExercise); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input data",
			"error":   err.Error(),
		})
		return
	}

	if err := connection.DB.Create(&trainingPlanExercise).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create training plan exercise",
			"error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "Training plan exercise created successfully",
		"data":    trainingPlanExercise,
	})
}
