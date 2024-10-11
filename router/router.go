package router

import (
	"github.com/Lynczera/goS/routes"
	"github.com/gin-gonic/gin"
)

// Adding the handlers for the routes
func SetupRoutes(r *gin.Engine) {
	//setup health
	healthGroup := r.Group("/health")
	routes.SetupHealthRoutes(healthGroup)

}
