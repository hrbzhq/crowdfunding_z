package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"crowdfunding/database"
	"crowdfunding/models"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	// Use default DB (or DATABASE_URL if set)
	database.InitDB()
	// ensure owner user exists
	email := "neuro@muscle.ai"
	username := "neuro_owner"
	password := "ChangeMe123!"

	var user models.User
	if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		// create user
		hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		user = models.User{
			Username: username,
			Email:    email,
			Password: string(hash),
		}
		tx := database.DB.Create(&user)
		if tx.Error != nil {
			log.Fatalf("failed to create user: %v", tx.Error)
		}
		fmt.Printf("Created user ID=%d email=%s\n", user.ID, user.Email)
	} else {
		fmt.Printf("Owner user already exists: ID=%d email=%s\n", user.ID, user.Email)
	}

	// create project
	proj := models.Project{
		Title:       "NeuroMuscleAI-MVP",
		Description: "MVP for NeuroMuscleAI: neurorehabilitation assistance using AI.",
		Goal:        50000,
		Deadline:    time.Now().AddDate(0, 3, 0).Format("2006-01-02"),
		Raised:      0,
		UserID:      user.ID,
	}

	if err := database.DB.Create(&proj).Error; err != nil {
		log.Fatalf("failed to create project: %v", err)
	}

	fmt.Printf("Created project ID=%d title=%s owner_id=%d\n", proj.ID, proj.Title, proj.UserID)

	os.Exit(0)
}
