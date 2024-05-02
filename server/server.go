package server

import (
	"library/controller"

	"github.com/gin-gonic/gin"
)

func Build() *gin.Engine {
	router := gin.Default()

	router.GET("/health", controller.PingHander)

	v1 := router.Group("api/v1")
	groupV1Endpoints(v1)

	return router

}
