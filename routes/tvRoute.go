package routes

import (
	"github.com/Lynczera/goS/handlers"
	"github.com/gin-gonic/gin"
)

func SetupTvRoutes(rg *gin.RouterGroup) {
	rg.POST("/getlisting", handlers.GetListing)
}
