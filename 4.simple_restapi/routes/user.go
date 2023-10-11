package routes

import (
	userControllers "myrestapi/4.simple_restapi/controllers"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.RouterGroup) {
	userGroup := r.Group("/user")
	{
		userGroup.GET("/", userControllers.GetUser)
	}
}
