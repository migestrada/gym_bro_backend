package routes

import (
	"gym-bro-backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetRoutes(router *gin.RouterGroup) {
	var set_route *gin.RouterGroup = router.Group("/sets")

	{
		set_route.GET("/", controllers.GetSets)
		set_route.POST("/", controllers.CreateSet)
		set_route.DELETE("/:id", controllers.DeleteSet)
		set_route.GET("/:id", controllers.GetSetByID)
		set_route.PUT("/:id", controllers.UpdateSet)
		set_route.PATCH("/:id", controllers.UpdateSet)
	}
}
