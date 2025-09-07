package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"crowdfunding/models"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("crowdfunding.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto migrate the schema
	DB.AutoMigrate(&models.User{}, &models.Project{}, &models.Funding{})
}
