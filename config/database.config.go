package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/koko-linoo/go-fiber-api/models"
)

var DB *gorm.DB

func InitDatabase() {
	var err error
	dsn := "host=localhost user=apple dbname=cms_db port=5432"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	DB.AutoMigrate(&models.User{})
}
