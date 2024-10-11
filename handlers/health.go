package handlers

import (
	"github.com/gin-gonic/gin"
)

func HandleHealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "ok",
	})
}

func HandleHealthCheckFail(c *gin.Context) {
	c.JSON(500, gin.H{
		"msg": "Forced server fail",
	})
}
