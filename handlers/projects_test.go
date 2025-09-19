package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
	"strconv"

	"crowdfunding/database"
	"github.com/gin-gonic/gin"
)

func setupRouterForTest() *gin.Engine {
	_ = os.Setenv("DATABASE_URL", "file::memory:?cache=shared")
	database.InitDB()
	g := gin.Default()
	g.POST("/register", Register)
	g.POST("/login", Login)
	g.POST("/projects", AuthMiddleware(), CreateProject)
	g.POST("/projects/:id/publish", AuthMiddleware(), PublishProject)
	g.POST("/projects/:id/fund", AuthMiddleware(), FundProject)
	g.GET("/projects/:id/progress", GetProgress)
	return g
}

func TestPublishPermissionAndFunding(t *testing.T) {
	g := setupRouterForTest()
	s := httptest.NewServer(g)
	defer s.Close()

	client := &http.Client{Timeout: 5 * time.Second}

	// register user A
	reg := map[string]string{"username": "A", "email": "a@example.com", "password": "pass123"}
	b, _ := json.Marshal(reg)
	resp, err := client.Post(s.URL+"/register", "application/json", bytes.NewReader(b))
	if err != nil { t.Fatalf("register A failed: %v", err) }
	resp.Body.Close()

	// login user A
	login := map[string]string{"email": "a@example.com", "password": "pass123"}
	b, _ = json.Marshal(login)
	resp, err = client.Post(s.URL+"/login", "application/json", bytes.NewReader(b))
	if err != nil { t.Fatalf("login A failed: %v", err) }
	var lres map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&lres)
	tokenA, _ := lres["token"].(string)
	resp.Body.Close()

	// create project by A
	proj := map[string]interface{}{"title": "T", "description": "D", "goal": 100.0, "deadline": "2030-01-01"}
	b, _ = json.Marshal(proj)
	req, _ := http.NewRequest("POST", s.URL+"/projects", bytes.NewReader(b))
	req.Header.Set("Authorization", "Bearer "+tokenA)
	req.Header.Set("Content-Type", "application/json")
	resp, err = client.Do(req)
	if err != nil { t.Fatalf("create project failed: %v", err) }
	var pres map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&pres)
	resp.Body.Close()
	id := int(pres["ID"].(float64))

	// attempt fund before publish -> should be 403
	fund := map[string]float64{"amount": 10}
	b, _ = json.Marshal(fund)
	req, _ = http.NewRequest("POST", s.URL+"/projects/"+strconv.Itoa(id)+"/fund", bytes.NewReader(b))
	req.Header.Set("Authorization", "Bearer "+tokenA)
	req.Header.Set("Content-Type", "application/json")
	resp, err = client.Do(req)
	if err != nil { t.Fatalf("fund before publish failed request: %v", err) }
	if resp.StatusCode != http.StatusForbidden { t.Fatalf("expected 403, got %d", resp.StatusCode) }
	resp.Body.Close()

	// publish
	req, _ = http.NewRequest("POST", s.URL+"/projects/"+strconv.Itoa(id)+"/publish", nil)
	req.Header.Set("Authorization", "Bearer "+tokenA)
	resp, err = client.Do(req)
	if err != nil { t.Fatalf("publish request failed: %v", err) }
	if resp.StatusCode != http.StatusOK { t.Fatalf("publish failed status: %d", resp.StatusCode) }
	resp.Body.Close()

	// fund after publish -> should succeed
	b, _ = json.Marshal(fund)
	req, _ = http.NewRequest("POST", s.URL+"/projects/"+strconv.Itoa(id)+"/fund", bytes.NewReader(b))
	req.Header.Set("Authorization", "Bearer "+tokenA)
	req.Header.Set("Content-Type", "application/json")
	resp, err = client.Do(req)
	if err != nil { t.Fatalf("fund after publish request failed: %v", err) }
	if resp.StatusCode != http.StatusOK { t.Fatalf("expected 200 after publish, got %d", resp.StatusCode) }
	resp.Body.Close()
}
