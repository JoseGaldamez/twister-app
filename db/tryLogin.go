package db

import (
	"twister/app/models"

	"golang.org/x/crypto/bcrypt"
)

func TryLogin(email string, password string) (models.User, bool) {

	user, found, _ := CheckUserExists(email)

	if !found {
		return user, false
	}

	passwordBytes := []byte(password)
	passwordBD := []byte(user.Password)

	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)

	if err != nil {
		return user, false
	}

	return user, true

}
