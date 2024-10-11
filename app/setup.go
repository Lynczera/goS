package app

import (
	"github.com/Lynczera/goS/router"
	"github.com/gin-gonic/gin"
)

func Setup() error {

	r := gin.Default()

	router.SetupRoutes(r)

	r.Run()

	return nil

}
