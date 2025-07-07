package services

import (
	"fmt"

	"github.com/koko-linoo/go-fiber-api/models"
	"golang.org/x/crypto/bcrypt"
)

func Login(username string, password string) (user *models.User, err error) {

	user, err = GetUserByName(username)

	fmt.Println(user.Password, user.Username)

	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		fmt.Println("Error")
		return nil, err
	}

	return user, nil
}
