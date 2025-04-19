package controllers

import (
	"log"
	"net/http"

	"gym-bro-backend/connection"

	"github.com/gin-gonic/gin"
)

type Exercise struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	PhotoURL    string
	VideoURL    string
}

func GetExercises(context *gin.Context) {
	var exercises []Exercise

	if err := connection.DB.Find(&exercises).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to retrieve exercises",
			"error":   err.Error(),
		})
	}

	context.JSON(http.StatusOK, gin.H{
		"message":   "Exercises retrieved successfully",
		"exercises": exercises,
	})
}

func CreateExercise(context *gin.Context) {
	var newExercise Exercise

	if err := context.ShouldBindJSON(&newExercise); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request ",
			"error":   err.Error(),
		})
		return
	}

	log.Println("Database connection:", connection.DB)

	// Check if the database connection is initialized
	if connection.DB == nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Database connection is not initialized",
		})
		return
	}

	if err := connection.DB.Create(&newExercise).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create exercise",
			"error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Exercise created successfully",
		"data":    newExercise,
	})
}
