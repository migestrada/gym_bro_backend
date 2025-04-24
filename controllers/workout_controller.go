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
