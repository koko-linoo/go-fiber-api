package services

import (
	"github.com/koko-linoo/go-fiber-api/config"
	"github.com/koko-linoo/go-fiber-api/models"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user models.User) (*models.User, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	user.Password = string(hash)

	if result := config.DB.Create(&user); result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func UpdateUser(id int, user map[string]any) (*models.User, error) {
	usrModel := config.DB.Model(&models.User{ID: uint(id)})

	if result := usrModel.Updates(user); result.Error != nil {
		return nil, result.Error
	}

	return GetUserByID(uint(id))
}

func GetUserByID(id uint) (user *models.User, err error) {
	user = &models.User{ID: id}

	if result := config.DB.First(&user); result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func GetUserByName(username string) (user *models.User, err error) {

	usrModel := config.DB.Model(&models.User{})

	if result := usrModel.Where("username = ?", username).First(&user); result.Error != nil {
		return nil, result.Error
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
