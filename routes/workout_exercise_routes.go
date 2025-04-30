package routes

import (
	"gym-bro-backend/controllers"

	"github.com/gin-gonic/gin"
)

func WorkoutExerciseRoutes(router *gin.RouterGroup) {
	var apiGroup *gin.RouterGroup = router.Group("/workout_exercises")
	{
		apiGroup.GET("/", controllers.GetWorkoutExercises)
		apiGroup.GET("/:id", controllers.GetWorkoutExerciseByID)
		apiGroup.POST("/", controllers.CreateWorkoutExercise)
		apiGroup.PUT("/:id", controllers.UpdateWorkoutExercise)
		apiGroup.PATCH("/:id", controllers.UpdateWorkoutExercise)
		apiGroup.DELETE("/:id", controllers.DeleteWorkoutExercise)
	}
}
