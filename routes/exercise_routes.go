package routes

import (
	"tgn-backend/controllers"

	"github.com/gin-gonic/gin"
)

func ExerciseRoutes(r *gin.RouterGroup) {
	exercises := r.Group("/exercises")

	{
		exercises.GET("/", controllers.GetExercises)
	}
}
