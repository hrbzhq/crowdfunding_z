package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"crowdfunding/database"
	"crowdfunding/handlers"
)

func main() {
	// Use an in-memory SQLite database for isolation and speed in tests
	// This avoids overwriting the developer's crowdfunding.db file.
	// Use shared cache so multiple connections see the same DB.
	// Format: file::memory:?cache=shared
	// For a temporary file instead: file:selftest.db?cache=shared&mode=rwc
	// Note: set DATABASE_URL externally if you prefer a different path.
	// Set before calling InitDB so it picks up the test DB.
	_ = os.Setenv("DATABASE_URL", "file::memory:?cache=shared")
	database.InitDB()

	// Build router like main.go
	r := gin.Default()
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)
	r.GET("/projects", handlers.GetProjects)
	a := r.Group("/")
	a.Use(handlers.AuthMiddleware())
	a.POST("/projects", handlers.CreateProject)
	a.POST("/projects/:id/fund", handlers.FundProject)
	r.GET("/projects/:id/progress", handlers.GetProgress)

	// Start test server
	ts := httptest.NewServer(r)
	defer ts.Close()

	client := &http.Client{Timeout: 10 * time.Second}

	fmt.Println("Test server URL:", ts.URL)

	// 1) Register
	token := ""
	var projectID int

	regBody := map[string]interface{}{
		"username": "selftester",
		"email":    fmt.Sprintf("selftester-%d@example.com", time.Now().Unix()),
		"password": "password123",
	}
	printStep("Register", func() error {
		b, _ := json.Marshal(regBody)
		resp, err := client.Post(ts.URL+"/register", "application/json", bytes.NewReader(b))
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		fmt.Println(string(body))
		if resp.StatusCode != http.StatusCreated {
			return fmt.Errorf("register failed: %s", resp.Status)
		}
		return nil
	})

	// 2) Login
	printStep("Login", func() error {
		loginBody := map[string]string{"email": regBody["email"].(string), "password": "password123"}
		b, _ := json.Marshal(loginBody)
		resp, err := client.Post(ts.URL+"/login", "application/json", bytes.NewReader(b))
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		var res map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
			return err
		}
		if tok, ok := res["token"].(string); ok {
			token = tok
			fmt.Println("token:", token)
			return nil
		}
		return fmt.Errorf("token missing in login response")
	})

	// 3) Create project (authenticated)
	printStep("CreateProject", func() error {
		proj := map[string]interface{}{"title": "SelfTest Project", "description": "A test project", "goal": 1000.0, "deadline": "2026-01-01"}
		b, _ := json.Marshal(proj)
		req, _ := http.NewRequest("POST", ts.URL+"/projects", bytes.NewReader(b))
		req.Header.Set("Authorization", "Bearer "+token)
		req.Header.Set("Content-Type", "application/json")
		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		var res map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
			return err
		}
		fmt.Println(res)
		if idf, ok := res["ID"].(float64); ok {
			projectID = int(idf)
			return nil
		}
		return fmt.Errorf("project ID missing")
	})

	// 4) Fund project (authenticated)
	printStep("FundProject", func() error {
		fund := map[string]interface{}{"amount": 100.0}
		b, _ := json.Marshal(fund)
		req, _ := http.NewRequest("POST", fmt.Sprintf(ts.URL+"/projects/%d/fund", projectID), bytes.NewReader(b))
		req.Header.Set("Authorization", "Bearer "+token)
		req.Header.Set("Content-Type", "application/json")
		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		fmt.Println(string(body))
		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("fund failed: %s", resp.Status)
		}
		return nil
	})

	// 5) Get progress
	printStep("GetProgress", func() error {
		resp, err := client.Get(fmt.Sprintf(ts.URL+"/projects/%d/progress", projectID))
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		fmt.Println(string(body))
		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("progress failed: %s", resp.Status)
		}
		return nil
	})

	fmt.Println("Self-test finished")
}

func printStep(name string, fn func() error) {
	fmt.Printf("--- %s ---\n", name)
	if err := fn(); err != nil {
		log.Fatalf("%s failed: %v", name, err)
	}
}
