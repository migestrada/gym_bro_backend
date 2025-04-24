package controllers

import (
	"gym-bro-backend/connection"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Workout struct {
	ID    int    `json:"id" grom:"primaryKey"`
	Index int    `json:"index" binding:"required"`
	Name  string `json:"name" binding:"required"`
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
