package controllers

import (
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

func DeleteExercise(context *gin.Context) {
	var exerciseId string = context.Param("id")

	var exercise Exercise

	if err := connection.DB.First(&exercise, exerciseId).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Exercise not found",
		})
		return
	}

	if err := connection.DB.Delete(&Exercise{}, exercise.ID).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to delete exercise",
			"error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Exercise deleted successfully",
		"data":    exercise,
	})
}

func GetExerciseByID(context *gin.Context) {
	var exerciseId string = context.Param("id")

	var exercise Exercise

	if err := connection.DB.First(&exercise, exerciseId).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Exercise not found",
			"error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Exercise retieved successfully",
		"data":    exercise,
	})
}

func UpdateExercise(context *gin.Context) {
	var exerciseId string = context.Param("id")

	var exercise Exercise

	if err := connection.DB.First(&exercise, exerciseId).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Exercise not found",
		})
		return
	}

	var updatedExercise Exercise
	if err := context.ShouldBindJSON(&updatedExercise); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
			"error":   err.Error(),
		})
		return
	}

	if err := connection.DB.Model(&exercise).Updates(&updatedExercise).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to update exercise",
			"error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Exercise updated successfully",
		"data":    exercise,
	})
}
