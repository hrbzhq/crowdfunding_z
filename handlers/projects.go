package handlers

import (
	"net/http"
	"strconv"
	"time"

	"crowdfunding/database"
	"crowdfunding/models"
	"github.com/gin-gonic/gin"
)

type CreateProjectRequest struct {
	Title       string  `json:"title" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Goal        float64 `json:"goal" binding:"required"`
	Deadline    string  `json:"deadline" binding:"required"`
}

func GetProjects(c *gin.Context) {
	var projects []models.Project
	// Optional status filter: ?status=published or ?status=draft
	status := c.Query("status")
	if status != "" {
		// If a status is requested, only allow querying drafts when authenticated and owner
		if status == "draft" {
			// require authentication
			uidRaw, ok := c.Get("user_id")
			if !ok {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization required to view drafts"})
				return
			}
			uid, ok := uidRaw.(uint)
			if !ok {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user id in context"})
				return
			}
			database.DB.Where("status = ? AND user_id = ?", status, uid).Find(&projects)
			c.JSON(http.StatusOK, projects)
			return
		}

		database.DB.Where("status = ?", status).Find(&projects)
		c.JSON(http.StatusOK, projects)
		return
	}

	// By default, show only published projects to unauthenticated users.
	// If the requester is authenticated and wants drafts, allow fetching drafts owned by the user.
	// If authenticated, show user's projects (both draft and published)
	if uidRaw, ok := c.Get("user_id"); ok {
		if uid, ok := uidRaw.(uint); ok {
			database.DB.Where("user_id = ?", uid).Find(&projects)
			c.JSON(http.StatusOK, projects)
			return
		}
	}

	// Default: only show published projects
	database.DB.Where("status = ?", "published").Find(&projects)
	c.JSON(http.StatusOK, projects)
}

func CreateProject(c *gin.Context) {
	var req CreateProjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	project := models.Project{
		Title:       req.Title,
		Description: req.Description,
		Goal:        req.Goal,
		Deadline:    req.Deadline,
		Raised:      0,
		Status:      "draft",
	}
	// If user_id is available in context (from AuthMiddleware), set it
	if uid, exists := c.Get("user_id"); exists {
		if u, ok := uid.(uint); ok {
			project.UserID = u
		}
	}
	if err := database.DB.Create(&project).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create project"})
		return
	}

	c.JSON(http.StatusCreated, project)
}

// PublishProject marks a draft project as published. Requires authentication and ownership.
func PublishProject(c *gin.Context) {
	id := c.Param("id")
	projectID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	var project models.Project
	if err := database.DB.First(&project, projectID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}

	// check ownership
	var uid uint
	if ctxUID, ok := c.Get("user_id"); ok {
		if u, ok := ctxUID.(uint); ok {
			uid = u
		}
	}
	if uid == 0 || project.UserID != uid {
		c.JSON(http.StatusForbidden, gin.H{"error": "not authorized to publish this project"})
		return
	}

	now := time.Now()
	project.Status = "published"
	project.PublishedAt = &now
	if err := database.DB.Save(&project).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to publish project"})
		return
	}

	// notify websocket clients that a project was published
	BroadcastJSON(map[string]interface{}{"type": "project_published", "project_id": project.ID, "status": project.Status})

	c.JSON(http.StatusOK, project)
}

func FundProject(c *gin.Context) {
	id := c.Param("id")
	projectID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	// Only amount is required in the request body. The user ID is taken from the
	// auth middleware (JWT) which sets "user_id" in the context.
	var req struct {
		Amount float64 `json:"amount" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var project models.Project
	if err := database.DB.First(&project, projectID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}

	// Only allow funding on published projects
	if project.Status != "published" {
		c.JSON(http.StatusForbidden, gin.H{"error": "project is not open for funding"})
		return
	}

	// user_id must be provided by auth middleware
	var uid uint
	if ctxUID, ok := c.Get("user_id"); ok {
		if u, ok := ctxUID.(uint); ok {
			uid = u
		}
	}
	if uid == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization required"})
		return
	}

	// Update raised amount and create funding record in a transaction
	tx := database.DB.Begin()
	project.Raised += req.Amount
	if err := tx.Save(&project).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to apply funding"})
		return
	}

	funding := models.Funding{
		ProjectID: uint(projectID),
		UserID:    uid,
		Amount:    req.Amount,
	}
	if err := tx.Create(&funding).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to record funding"})
		return
	}
	tx.Commit()

	// notify websocket clients about funding update
	BroadcastJSON(map[string]interface{}{"type": "project_funded", "project_id": uint(projectID), "raised": project.Raised, "goal": project.Goal})

	c.JSON(http.StatusOK, gin.H{"message": "Funding successful"})
}

func GetProgress(c *gin.Context) {
	id := c.Param("id")
	projectID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	var project models.Project
	if err := database.DB.First(&project, projectID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}

	progress := (project.Raised / project.Goal) * 100
	c.JSON(http.StatusOK, gin.H{"progress": progress, "raised": project.Raised, "goal": project.Goal})
}
