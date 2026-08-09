package controllers

import (
	"backend-porto/models"
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

// GetProjects returns all projects
func GetProjects(c *gin.Context) {
	var projects []models.Project
	models.DB.Order("date_modified desc").Find(&projects)
	c.JSON(http.StatusOK, projects)
}

// GetProject returns a single project by ID
func GetProject(c *gin.Context) {
	var project models.Project
	if err := models.DB.First(&project, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}
	c.JSON(http.StatusOK, project)
}

// CreateProject creates a new project with image and name
func CreateProject(c *gin.Context) {
	// Parse multipart form
	name := c.PostForm("name")
	
	// Handle image upload
	file, err := c.FormFile("image")
	var imageURL string
	if err == nil {
		filename := fmt.Sprintf("%d_%s", time.Now().Unix(), filepath.Base(file.Filename))
		uploadPath := filepath.Join("public", "uploads", filename)
		if err := c.SaveUploadedFile(file, uploadPath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
			return
		}
		imageURL = "/uploads/" + filename // Route serving static files
	}

	project := models.Project{
		Name:     name,
		ImageURL: imageURL,
		Status:   "Draft", // Default status
	}

	models.DB.Create(&project)
	c.JSON(http.StatusCreated, project)
}

// UpdateProject updates project details (including flow, jobdesc, tech stack)
func UpdateProject(c *gin.Context) {
	var project models.Project
	if err := models.DB.First(&project, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}

	var input models.Project
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&project).Updates(input)
	c.JSON(http.StatusOK, project)
}

// DeleteProject removes a project
func DeleteProject(c *gin.Context) {
	var project models.Project
	if err := models.DB.First(&project, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}

	models.DB.Delete(&project)
	c.JSON(http.StatusOK, gin.H{"message": "Project deleted successfully"})
}
