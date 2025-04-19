package routes

import (
	"gym-bro-backend/controllers"

	"github.com/gin-gonic/gin"
)

func ExerciseRoutes(router *gin.RouterGroup) {
	exercises := router.Group("/exercises")

	{
		exercises.GET("/", controllers.GetExercises)
		exercises.POST("/", controllers.CreateExercise)
	}
}
