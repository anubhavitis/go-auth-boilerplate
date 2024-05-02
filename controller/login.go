package controller

import (
	"library/services/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	var user LoginRequest
	var resp LoginResponse
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	token, err := users.Login(user.UserName, user.Password)

	if err != nil {
		resp.Error = err
		c.JSON(http.StatusUnauthorized, resp)
		return
	}

	// Send token in response
	resp.Message = "Success"
	resp.Token = token
	c.JSON(http.StatusOK, resp)
}
