package users

import (
	"fmt"
	"library/constants"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type User struct {
	UUID     uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	IsAdmin  bool      `json:"isAdmin"`
}

var users = []User{
	{UUID: uuid.New(), Username: "user1", Password: "password1"},
	{UUID: uuid.New(), Username: "user2", Password: "password2"},
	{UUID: uuid.New(), Username: "admin", Password: "password3", IsAdmin: true},
}

func generateJWTToken(user User) (string, error) {

	claims := jwt.MapClaims{
		"id":       user.UUID,
		"username": user.Username,
		"isAdmin":  user.IsAdmin,
		"exp":      time.Now().Add(time.Hour * 1).Unix(), // Token expires in 1 hours
	}

	// Create token object with claims and signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token string using the secret key
	tokenString, err := token.SignedString([]byte(constants.JWTSecretKey))
	fmt.Printf("token %v and err %v", token, err)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
