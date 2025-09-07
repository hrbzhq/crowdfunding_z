package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"crowdfunding/handlers"
	"crowdfunding/database"
)

func main() {
	// Initialize database
	database.InitDB()

	// Create Gin router
	r := gin.Default()

	// Routes
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)
	r.GET("/projects", handlers.GetProjects)
	r.POST("/projects", handlers.CreateProject)
	r.POST("/projects/:id/fund", handlers.FundProject)
	r.GET("/projects/:id/progress", handlers.GetProgress)

	// WebSocket for real-time updates
	r.GET("/ws", handlers.WebSocketHandler)

	// Start server
	log.Println("Server starting on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
