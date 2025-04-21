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
		exercises.DELETE("/:id", controllers.DeleteExercise)
		exercises.GET("/:id", controllers.GetExerciseByID)
		exercises.PUT("/:id", controllers.UpdateExercise)
		exercises.PATCH("/:id", controllers.UpdateExercise)
	}
}
