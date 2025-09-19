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

func TestNonOwnerCannotPublish(t *testing.T) {
	g := setupRouterForTest()
	s := httptest.NewServer(g)
	defer s.Close()

	client := &http.Client{Timeout: 5 * time.Second}
	// use unique suffix to avoid collisions when tests share the same in-memory DB
	suffix := strconv.FormatInt(time.Now().UnixNano(), 10)

	// register user A
	regA := map[string]string{"username": "A" + suffix, "email": "a2" + suffix + "@example.com", "password": "pass123"}
	b, _ := json.Marshal(regA)
	resp, err := client.Post(s.URL+"/register", "application/json", bytes.NewReader(b))
	if err != nil { t.Fatalf("register A failed: %v", err) }
	resp.Body.Close()

	// login user A (use the same unique email)
	loginA := map[string]string{"email": "a2" + suffix + "@example.com", "password": "pass123"}
	b, _ = json.Marshal(loginA)
	resp, err = client.Post(s.URL+"/login", "application/json", bytes.NewReader(b))
	if err != nil { t.Fatalf("login A failed: %v", err) }
	var lresA map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&lresA)
	tokenA, _ := lresA["token"].(string)
	resp.Body.Close()

	// user A creates a project
	proj := map[string]interface{}{"title": "T2", "description": "D2", "goal": 50.0, "deadline": "2030-01-01"}
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

	// register user B
	regB := map[string]string{"username": "B" + suffix, "email": "b" + suffix + "@example.com", "password": "pass123"}
	b, _ = json.Marshal(regB)
	resp, err = client.Post(s.URL+"/register", "application/json", bytes.NewReader(b))
	if err != nil { t.Fatalf("register B failed: %v", err) }
	resp.Body.Close()

	// login user B (use the same unique email)
	loginB := map[string]string{"email": "b" + suffix + "@example.com", "password": "pass123"}
	b, _ = json.Marshal(loginB)
	resp, err = client.Post(s.URL+"/login", "application/json", bytes.NewReader(b))
	if err != nil { t.Fatalf("login B failed: %v", err) }
	var lresB map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&lresB)
	tokenB, _ := lresB["token"].(string)
	resp.Body.Close()

	// user B attempts to publish A's project -> should be forbidden
	req, _ = http.NewRequest("POST", s.URL+"/projects/"+strconv.Itoa(id)+"/publish", nil)
	req.Header.Set("Authorization", "Bearer "+tokenB)
	resp, err = client.Do(req)
	if err != nil { t.Fatalf("publish by B request failed: %v", err) }
	if resp.StatusCode != http.StatusForbidden { t.Fatalf("expected 403 for non-owner publish, got %d", resp.StatusCode) }
	resp.Body.Close()
}
