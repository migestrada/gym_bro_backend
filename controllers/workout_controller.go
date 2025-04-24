package controllers

import (
	"gym-bro-backend/connection"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Workout struct {
	ID    int    `json:"id" grom:"primaryKey"`
	Name  string `json:"name" binding:"required"`
	Order int    `json:"order" binding:"required"`
}

func GetWorkouts(context *gin.Context) {
	var workouts []Workout

	if err := connection.DB.Find(&workouts).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to retrieve workouts",
			"error":   err.Error(),
		})

		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Workouts retrieved successfully",
		"data":    workouts,
	})
}

func CreateWorkout(context *gin.Context) {
	var newWorkout Workout

	if err := context.ShouldBindJSON(&newWorkout); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
			"error":   err.Error(),
		})

		return
	}

	if err := connection.DB.Create(&newWorkout).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create workout",
			"error":   err.Error(),
		})

		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Workout created successfully",
		"data":    newWorkout,
	})
}

func DeleteWorkout(context *gin.Context) {
	var workoutId string = context.Param("id")
	var workout Workout

	if err := connection.DB.First(&workout, workoutId).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Workout not found",
			"error":   err.Error(),
		})

		return
	}

	if err := connection.DB.Delete(&workout).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to delete workout",
			"error":   err.Error(),
		})

		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Workout deleted successfully",
		"data":    workout,
	})
}

func GetWorkoutByID(context *gin.Context) {
	var workoutId string = context.Param("id")
	var workout Workout

	if err := connection.DB.First(&workout, workoutId).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Workout not found",
			"error":   err.Error(),
		})

		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Workout retrieved successfully",
		"data":    workout,
	})
}
