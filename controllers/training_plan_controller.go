package controllers

import (
	"gym-bro-backend/connection"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TrainingPlan struct {
	ID          int    `json:"id" gorm:"primary_key"`
	Name        string `json:"name" gorm:"required,not null"`
	Description string `json:"description" gorm:"required,not null"`
}

func GetTrainingPlans(context *gin.Context) {
	var trainingPlans []TrainingPlan

	if err := connection.DB.Find(&trainingPlans).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to retrieve training plans",
			"error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Training plans retrieved successfully",
		"data":    trainingPlans,
	})
}

func CreateTrainingPlan(context *gin.Context) {
	var trainingPlan TrainingPlan

	if err := context.ShouldBindJSON(&trainingPlan); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input",
			"error":   err.Error(),
		})
		return
	}

	if err := connection.DB.Create(&trainingPlan).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create training plan",
			"error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "Training plan created successfully",
		"data":    trainingPlan,
	})
}

func GetTrainingPlanByID(context *gin.Context) {
	var trainingPlanId string = context.Param("id")
	var trainingPlan TrainingPlan

	if err := connection.DB.First(&trainingPlan, trainingPlanId).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Training plan not found",
			"error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Training plan retrieved successfully",
		"data":    trainingPlan,
	})
}
