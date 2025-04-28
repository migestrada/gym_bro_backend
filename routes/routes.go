package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	var apiGroup *gin.RouterGroup = router.Group("/api")
	ExerciseRoutes(apiGroup)
	SetRoutes(apiGroup)
	WorkoutRoutes(apiGroup)
	TrainingPlanRoutes(apiGroup)
	TrainingPlanExerciseRoutes(apiGroup)
	WorkoutSessionRoutes(apiGroup)
}
