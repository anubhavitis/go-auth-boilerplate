package users

import (
	"errors"
)

func Login(username string, password string) (token string, err error) {
	// Find user by username and password (for demonstration purposes)
	var authenticatedUser User
	for _, u := range users {
		if u.Username == username && u.Password == password {
			authenticatedUser = u
			break
		}
	}
	// Check if user is found
	if authenticatedUser.Username == "" {
		return token, errors.New("user not found")
	}

	// Generate JWT token
	token, err = generateJWTToken(authenticatedUser)
	return
}
