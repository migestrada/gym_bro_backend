package routes

import (
	"gym-bro-backend/controllers"

	"github.com/gin-gonic/gin"
)

func TrainingPlanRoutes(router *gin.RouterGroup) {
	var apiGroup *gin.RouterGroup = router.Group("/training_plans")

	{
		apiGroup.GET("/", controllers.GetTrainingPlans)
		apiGroup.POST("/", controllers.CreateTrainingPlan)
		apiGroup.GET("/:id", controllers.GetTrainingPlanByID)
		apiGroup.PUT("/:id", controllers.UpdateTrainingPlan)
		apiGroup.PATCH("/:id", controllers.UpdateTrainingPlan)
		apiGroup.DELETE("/:id", controllers.DeleteTrainingPlan)
	}
}
