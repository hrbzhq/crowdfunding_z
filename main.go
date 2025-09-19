package main

import (
	"context"
	"flag"
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
	// CLI flags
	autoFlag := flag.Bool("autoupdate", false, "Enable autoupdater (overrides ENABLE_AUTOUPDATE env)")
	autoInterval := flag.String("autoupdate-interval", "", "Autoupdater interval duration (e.g. 30m, 1h). Overrides AUTOSCHED_INTERVAL env")
	flag.Parse()

	// Initialize database
	database.InitDB()
	// Optionally start autoupdater scheduler if enabled via CLI flag or env
	enabled := shouldEnableAutoupdate(os.Args[1:], os.Getenv("ENABLE_AUTOUPDATE")) || *autoFlag
	if enabled {
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

		// determine interval: CLI flag > env var > default 1h
		interval := 1 * time.Hour
		if *autoInterval != "" {
			if d, err := time.ParseDuration(*autoInterval); err == nil {
				interval = d
			}
		} else if env := os.Getenv("AUTOSCHED_INTERVAL"); env != "" {
			if d, err := time.ParseDuration(env); err == nil {
				interval = d
			}
		}
		sched := autoupdater.NewScheduler(fetcher, analyzer, updater, interval)
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

// shouldEnableAutoupdate decides whether autoupdater should be enabled based on
// provided CLI args and the ENABLE_AUTOUPDATE env value. This helper keeps the
// logic testable. It returns true if args contain "--autoupdate" or if
// envEnable equals "true" (case-insensitive).
func shouldEnableAutoupdate(args []string, envEnable string) bool {
	for _, a := range args {
		if a == "--autoupdate" {
			return true
		}
	}
	return strings.ToLower(envEnable) == "true"
}
