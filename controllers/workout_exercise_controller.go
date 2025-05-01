package controllers

import (
	"gym-bro-backend/connection"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WorkoutExercise struct {
	ID         uint `gorm:"primaryKey"`
	WorkoutID  uint `json:"workout_id" binding:"required"`
	ExerciseID uint `json:"exercise_id" binding:"required"`
}

func GetWorkoutExercises(context *gin.Context) {
	var workoutExercises []WorkoutExercise

	if err := connection.DB.Find(&workoutExercises).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to retrieve workout exercises",
			"error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Workout exercises retrieved successfully",
		"data":    workoutExercises,
	})
}

func GetWorkoutExerciseByID(context *gin.Context) {
	var workoutExercise WorkoutExercise
	var id string = context.Param("id")

	if err := connection.DB.First(&workoutExercise, id).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Workout exercise not found",
			"error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Workout exercise retrieved successfully",
		"data":    workoutExercise,
	})
}

func CreateWorkoutExercise(context *gin.Context) {
	var workoutExercise WorkoutExercise

	if err := context.ShouldBindJSON(&workoutExercise); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input data",
			"error":   err.Error(),
		})
		return
	}

	if err := connection.DB.Create(&workoutExercise).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create workout exercise",
			"error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "Workout exercise created successfully",
		"data":    workoutExercise,
	})
}

func UpdateWorkoutExercise(context *gin.Context) {
	var workoutExercise WorkoutExercise
	var workoutExerciseId string = context.Param("id")

	if err := connection.DB.First(&workoutExercise, workoutExerciseId).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Workout exercise not found",
			"error":   err.Error(),
		})
		return
	}

	if err := context.ShouldBindJSON(&workoutExercise); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input data",
			"error":   err.Error(),
		})
		return
	}

	if err := connection.DB.Save(&workoutExercise).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to update workout exercise",
			"error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Workout exercise updated successfully",
		"data":    workoutExercise,
	})
}

func DeleteWorkoutExercise(context *gin.Context) {
	var workoutExercise WorkoutExercise
	var workoutExerciseId string = context.Param("id")

	if err := connection.DB.First(&workoutExercise, workoutExerciseId).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Workout exercise not found",
			"error":   err.Error(),
		})
		return
	}

	if err := connection.DB.Delete(&workoutExercise).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to delete workout exercise",
			"error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Workout exercise deleted successfully",
	})
}
