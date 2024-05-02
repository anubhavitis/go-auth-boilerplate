package users

import (
	"fmt"
	"library/constants"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Middleware function to validate JWT token
func IsAdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		// Check if the token is missing
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Parse and validate the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Validate the signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(constants.JWTSecretKey), nil
		})

		// Check if there is an error parsing the token
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Check if the token is valid
		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token expired"})
			c.Abort()
			return
		}

		// Extract claims from the token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		// Extract username from claims and log it
		isAdmin, ok := claims["isAdmin"].(bool)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid isAdmin claim"})
			c.Abort()
			return
		}

		if !isAdmin {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "This request is only available for admin"})
			c.Abort()
		}

		// Token is valid, continue with the request
		c.Next()
	}
}
