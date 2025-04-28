package routes

import (
	"gym-bro-backend/controllers"

	"github.com/gin-gonic/gin"
)

func WorkoutSessionRoutes(router *gin.RouterGroup) {
	var apiGroup *gin.RouterGroup = router.Group("/workout_sessions")

	{
		apiGroup.GET("/", controllers.GetWorkoutSessions)
	}
}
