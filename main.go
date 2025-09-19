package main

import (
	"context"
	"log"
	"os"
	"strings"
	"time"

	"crowdfunding/database"
	"crowdfunding/handlers"
	autoupdater "crowdfunding/tools/autoupdater"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database
	database.InitDB()

	// Optionally start autoupdater scheduler if enabled via env
	if strings.ToLower(os.Getenv("ENABLE_AUTOUPDATE")) == "true" {
		// build fetcher
		var fetcher autoupdater.Fetcher
		urls := os.Getenv("AUTOFETCH_URLS")
		if urls != "" {
			parts := strings.Split(urls, ",")
			fetcher = autoupdater.NewHTTPFetcher(parts, 10*time.Second)
		} else {
			fetcher = &autoupdater.MockFetcher{}
		}

		// build analyzer and updater
		analyzer := &autoupdater.MockAnalyzer{}
		var updater autoupdater.Updater = &autoupdater.MockUpdater{}
		ghToken := os.Getenv("GITHUB_TOKEN")
		ghRepo := os.Getenv("GITHUB_REPO")
		if ghToken != "" && ghRepo != "" {
			updater = autoupdater.NewGitHubUpdater(ghToken, ghRepo)
		}

		sched := autoupdater.NewScheduler(fetcher, analyzer, updater, 1*time.Hour)
		ctx := context.Background()
		sched.Start(ctx)
		defer sched.Stop()
		log.Println("Autoupdater started")
	}

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
