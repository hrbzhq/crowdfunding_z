package database

import (
	"log"
	"os"

	"crowdfunding/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB initializes the database connection. It respects the DATABASE_URL
// environment variable. If not set, it falls back to "crowdfunding.db".
func InitDB() {
	var err error
	dbPath := os.Getenv("DATABASE_URL")
	if dbPath == "" {
		dbPath = "crowdfunding.db"
	}

	DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto migrate the schema
	DB.AutoMigrate(&models.User{}, &models.Project{}, &models.Funding{})
}
