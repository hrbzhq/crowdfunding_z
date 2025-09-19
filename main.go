package main

import (
	"log"

	"crowdfunding/database"
	"crowdfunding/handlers"
	"github.com/gin-gonic/gin"
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
	// Protected routes
	auth := r.Group("/")
	auth.Use(handlers.AuthMiddleware())
	auth.POST("/projects", handlers.CreateProject)
	auth.POST("/projects/:id/fund", handlers.FundProject)
	auth.POST("/projects/:id/publish", handlers.PublishProject)
	r.GET("/projects/:id/progress", handlers.GetProgress)

	// WebSocket for real-time updates
	r.GET("/ws", handlers.WebSocketHandler)

	// Start server
	log.Println("Server starting on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
