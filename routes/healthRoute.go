package routes

import (
	"github.com/Lynczera/goS/handlers"
	"github.com/gin-gonic/gin"
)

func SetupHealthRoutes(rg *gin.RouterGroup) {
	rg.GET("/check", handlers.HandleHealthCheck)
	rg.GET("/fail", handlers.HandleHealthCheckFail)
}
