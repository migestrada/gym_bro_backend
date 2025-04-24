package routes

import (
	"gym-bro-backend/controllers"

	"github.com/gin-gonic/gin"
)

func WorkoutRoutes(router *gin.RouterGroup) {
	var workoutRouter = router.Group("/workouts")

	{
		workoutRouter.GET("/", controllers.GetWorkouts)
		workoutRouter.POST("/", controllers.CreateWorkout)
		workoutRouter.DELETE("/:id", controllers.DeleteWorkout)
		workoutRouter.GET("/:id", controllers.GetWorkoutByID)
	}
}
