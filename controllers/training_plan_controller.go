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
