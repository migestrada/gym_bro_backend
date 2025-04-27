package routes

import (
	"gym-bro-backend/controllers"

	"github.com/gin-gonic/gin"
)

func TrainingPlanExerciseRoutes(router *gin.RouterGroup) {
	var apiGroup *gin.RouterGroup = router.Group("/training_plan_exercises")
	{
		apiGroup.GET("/", controllers.GetTrainingPlanExercises)
		apiGroup.GET("/:id", controllers.GetTrainingPlanExerciseByID)
	}
}
