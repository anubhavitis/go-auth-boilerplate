package server

import (
	"library/controller"
	"library/services/users"

	"github.com/gin-gonic/gin"
)

func groupV1Endpoints(v1 *gin.RouterGroup) {
	v1.GET("/home", controller.AllBookHandler)
	v1.POST("/login", controller.LoginHandler)
	v1.POST("/addBook", users.IsAdminMiddleware(), controller.AddBookHandler)
	v1.DELETE("/deleteBook", users.IsAdminMiddleware(), controller.DeleteBookHandler)
}
