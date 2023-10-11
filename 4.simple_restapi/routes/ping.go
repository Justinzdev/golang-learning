package routes

import (
	pingControllers "myrestapi/4.simple_restapi/controllers"

	"github.com/gin-gonic/gin"
)

func SetupPingRoutes(r *gin.RouterGroup) {
	pingGroup := r.Group("/ping")
	{
		pingGroup.GET("/", pingControllers.PangOut)
	}
}
