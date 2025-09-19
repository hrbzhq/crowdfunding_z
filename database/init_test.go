package database

import (
	"os"
	"path/filepath"
	"testing"

	"crowdfunding/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// TestInitDB creates a temporary sqlite file, opens it with GORM, runs AutoMigrate
// and verifies the DB handle and migrations succeed.
func TestInitDB(t *testing.T) {
	dir := t.TempDir()
	dbPath := filepath.Join(dir, "test_crowdfunding.db")

	// open a new gorm DB with sqlite pointing to temp file
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open temp sqlite db: %v", err)
	}

	// Run AutoMigrate for the models
	if err := db.AutoMigrate(&models.User{}, &models.Project{}, &models.Funding{}); err != nil {
		t.Fatalf("AutoMigrate failed: %v", err)
	}

	// Ensure the DB file exists on disk
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		t.Fatalf("expected sqlite db file to exist at %s", dbPath)
	}

	// Close underlying sql DB to release file lock before cleanup
	sqlDB, err := db.DB()
	if err == nil {
		_ = sqlDB.Close()
	}
}
