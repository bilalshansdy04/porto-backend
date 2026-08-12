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
		Name:       name,
		ImageURL:   imageURL,
		Status:     "Draft", // Default status
		IsComplete: false,
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

	models.DB.Model(&project).Select("Name", "Description", "Status", "IsComplete", "TechStack", "ProjectFlow", "JobDesc", "Link", "CarouselImages").Updates(input)

	// Check completeness: no null/empty required fields
	hasImage := project.ImageURL != "" || len(project.CarouselImages) > 0
	isComplete := project.Name != "" && project.Description != "" && hasImage && len(project.TechStack) > 0 && len(project.ProjectFlow) > 0 && len(project.JobDesc) > 0
	
	if project.IsComplete != isComplete {
		project.IsComplete = isComplete
		models.DB.Model(&project).Update("is_complete", isComplete)
	}

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

// UploadProjectImages handles uploading multiple images for the project carousel
func UploadProjectImages(c *gin.Context) {
	var project models.Project
	if err := models.DB.First(&project, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse form"})
		return
	}

	files := form.File["images"]
	if len(files) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No images provided"})
		return
	}

	var newImageURLs []string
	for _, file := range files {
		filename := fmt.Sprintf("%d_%s", time.Now().UnixNano(), filepath.Base(file.Filename))
		uploadPath := filepath.Join("public", "uploads", filename)
		if err := c.SaveUploadedFile(file, uploadPath); err != nil {
			continue // Skip failed uploads
		}
		newImageURLs = append(newImageURLs, "/uploads/"+filename)
	}

	// Append to existing carousel images
	project.CarouselImages = append(project.CarouselImages, newImageURLs...)
	models.DB.Save(&project)

	c.JSON(http.StatusOK, project)
}

// UpdateProjectThumbnail handles uploading a new main thumbnail
func UpdateProjectThumbnail(c *gin.Context) {
	var project models.Project
	if err := models.DB.First(&project, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No image provided"})
		return
	}

	filename := fmt.Sprintf("%d_%s", time.Now().UnixNano(), filepath.Base(file.Filename))
	uploadPath := filepath.Join("public", "uploads", filename)
	if err := c.SaveUploadedFile(file, uploadPath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
		return
	}

	project.ImageURL = "/uploads/" + filename
	models.DB.Save(&project)

	c.JSON(http.StatusOK, project)
}
