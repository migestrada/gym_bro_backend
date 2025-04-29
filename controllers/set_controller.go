package controllers

import (
	"gym-bro-backend/connection"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Set struct {
	ID         uint    `json:"id" gorm:"primaryKey"`
	Reps       int     `json:"reps" binding:"required"`
	RestTime   int     `json:"rest_time" binding:"required"`
	Weight     float32 `json:"weight" binding:"required"`
	WeightUnit string  `json:"weight_unit" binding:"required"`
	ExerciseID uint    `json:"exercise_id"`
}

func GetSets(context *gin.Context) {
	var sets []Set

	if err := connection.DB.Find(&sets).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to retrieve sets",
			"error":   err.Error(),
		})
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Sets retrieved successfully",
		"data":    sets,
	})
}

func CreateSet(context *gin.Context) {
	var newSet Set

	if err := context.ShouldBindJSON(&newSet); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
			"error":   err.Error(),
		})
		return
	}

	if err := connection.DB.Create(&newSet).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create set",
			"error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "Set created successfully",
		"data":    newSet,
	})
}

func DeleteSet(context *gin.Context) {
	var setId string = context.Param("id")

	var set Set

	if err := connection.DB.First(&set, setId).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Set fot found",
			"error":   err.Error(),
		})

		return
	}

	if err := connection.DB.Delete(&Set{}, set.ID).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to delete set",
			"error":   err.Error(),
		})

		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Set deleted successfully",
		"data":    set,
	})
}

func GetSetByID(context *gin.Context) {
	var setId string = context.Param("id")

	var set Set

	if err := connection.DB.First(&set, setId).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Set fot found",
			"error":   err.Error(),
		})

		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Set retrieved successfully",
		"data":    set,
	})
}

func UpdateSet(context *gin.Context) {
	var setId string = context.Param("id")

	var set Set

	if err := connection.DB.First(&set, setId).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Set fot found",
			"error":   err.Error(),
		})
		return
	}

	var updatedSet Set
	if err := context.ShouldBindJSON(&updatedSet); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
			"error":   err.Error(),
		})
		return
	}

	if err := connection.DB.Model(&set).Updates(&updatedSet).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to update set",
			"error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Set updated successfully",
		"data":    set,
	})
}
