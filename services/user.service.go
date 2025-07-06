package services

import (
	"github.com/koko-linoo/go-fiber-api/config"
	"github.com/koko-linoo/go-fiber-api/models"
)

func CreateUser(user models.User) (*models.User, error) {
	if result := config.DB.Create(&user); result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func UpdateUser(user models.User) (*models.User, error) {
	if result := config.DB.Save(&user); result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func GetUser(username string) (user models.User, err error) {

	user = models.User{Username: username}

	if result := config.DB.First(&user); result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

func GetAllUsers() (users []models.User, err error) {
	if result := config.DB.Find(&users); result.Error != nil {
		return users, result.Error
	}
	return users, nil
}

func DeleteUser(id uint) (err error) {
	user := models.User{ID: id}
	if result := config.DB.Delete(&user); result.Error != nil {
		return result.Error
	}
	return nil
}
