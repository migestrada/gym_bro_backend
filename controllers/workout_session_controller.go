package controllers

import (
	"gym-bro-backend/connection"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type WorkoutSession struct {
	ID                uint      `json:"id" gorm:"primaryKey"`
	Date              time.Time `json:"date" binding:"required"`
	WorkoutExerciseID uint      `json:"workout_exercise_id" binding:"required"`
	SetID             uint      `json:"set_id" binding:"required"`
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

func GetWorkoutSessionByID(context *gin.Context) {
	var workoutSession WorkoutSession
	var workoutSessionID string = context.Param("id")

	if err := connection.DB.First(&workoutSession, workoutSessionID).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Workout session not found",
			"error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Workout session retrieved successfully",
		"data":    workoutSession,
	})
}

func CreateWorkoutSession(context *gin.Context) {
	var newWorkoutSession WorkoutSession

	if err := context.ShouldBindJSON(&newWorkoutSession); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
			"error":   err.Error(),
		})
		return
	}

	if err := connection.DB.Create(&newWorkoutSession).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create workout session",
			"error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "Workout session created successfully",
		"data":    newWorkoutSession,
	})
}

func UpdateWorkoutSession(context *gin.Context) {
	var workoutSession WorkoutSession
	var workoutSessionID string = context.Param("id")

	if err := connection.DB.First(&workoutSession, workoutSessionID).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Workout session not found",
			"error":   err.Error(),
		})
		return
	}

	if err := context.ShouldBindJSON(&workoutSession); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
			"error":   err.Error(),
		})
		return
	}

	if err := connection.DB.Save(&workoutSession).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to update workout session",
			"error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Workout session updated successfully",
		"data":    workoutSession,
	})
}

func DeleteWorkoutSession(context *gin.Context) {
	var workoutSession WorkoutSession
	var workoutSessionID string = context.Param("id")

	if err := connection.DB.First(&workoutSession, workoutSessionID).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Workout session not found",
			"error":   err.Error(),
		})
		return
	}

	if err := connection.DB.Delete(&workoutSession).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to delete workout session",
			"error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Workout session deleted successfully",
		"data":    workoutSession,
	})
}
