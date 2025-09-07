package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"crowdfunding/models"
	"crowdfunding/database"
)

type CreateProjectRequest struct {
	Title       string  `json:"title" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Goal        float64 `json:"goal" binding:"required"`
	Deadline    string  `json:"deadline" binding:"required"`
}

func GetProjects(c *gin.Context) {
	var projects []models.Project
	database.DB.Find(&projects)
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
	}

	if err := database.DB.Create(&project).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create project"})
		return
	}

	c.JSON(http.StatusCreated, project)
}

func FundProject(c *gin.Context) {
	id := c.Param("id")
	projectID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	var req struct {
		Amount float64 `json:"amount" binding:"required"`
		UserID uint    `json:"user_id" binding:"required"`
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

	project.Raised += req.Amount
	database.DB.Save(&project)

	// Create funding record
	funding := models.Funding{
		ProjectID: uint(projectID),
		UserID:    req.UserID,
		Amount:    req.Amount,
	}
	database.DB.Create(&funding)

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
