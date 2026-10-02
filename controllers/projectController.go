package controllers

import (
	"backend-porto/models"
	"context"
	"net/http"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
)

// CleanupUnusedImages is a no-op for Cloudinary
func CleanupUnusedImages() {
}

// GetProjects returns all projects
func GetProjects(c *gin.Context) {
	var projects []models.Project
	models.DB.Preload("Screenshots").Order("date_modified desc").Find(&projects)
	c.JSON(http.StatusOK, projects)
}

// GetProject returns a single project by ID
func GetProject(c *gin.Context) {
	var project models.Project
	if err := models.DB.Preload("Screenshots").First(&project, c.Param("id")).Error; err != nil {
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
	fileHeader, err := c.FormFile("image")
	var imageURL string
	if err == nil {
		file, _ := fileHeader.Open()
		defer file.Close()
		cld, err := cloudinary.New()
		if err == nil {
			resp, err := cld.Upload.Upload(context.Background(), file, uploader.UploadParams{Folder: "portfolio"})
			if err == nil {
				imageURL = resp.SecureURL
			}
		}
	}

	project := models.Project{
		Name:       name,
		ImageURL:   imageURL,
		Status:     "Draft", // Default status
		IsVisible:  false,
	}

	models.DB.Create(&project)
	c.JSON(http.StatusCreated, project)
}

// UpdateProject updates project details (including flow, jobdesc, tech stack)
func UpdateProject(c *gin.Context) {
	var project models.Project
	if err := models.DB.Preload("Screenshots").First(&project, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}

	var input models.Project
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&project).Select("Name", "Description", "Status", "IsVisible", "TechStack", "ProjectFlow", "JobDesc", "Link", "CarouselImages").Updates(input)
	models.DB.Model(&project).Association("Screenshots").Replace(input.Screenshots)

	// Refetch to include updated screenshots in response
	models.DB.Preload("Screenshots").First(&project, project.ID)

	go CleanupUnusedImages()

	c.JSON(http.StatusOK, project)
}

// DeleteProject removes a project
func DeleteProject(c *gin.Context) {
	var project models.Project
	if err := models.DB.Preload("Screenshots").First(&project, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}

	models.DB.Delete(&project)
	go CleanupUnusedImages()
	c.JSON(http.StatusOK, gin.H{"message": "Project deleted successfully"})
}

// UploadProjectImages handles uploading multiple images for the project carousel
func UploadProjectImages(c *gin.Context) {
	var project models.Project
	if err := models.DB.Preload("Screenshots").First(&project, c.Param("id")).Error; err != nil {
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
	cld, _ := cloudinary.New()
	for _, fileHeader := range files {
		file, _ := fileHeader.Open()
		resp, err := cld.Upload.Upload(context.Background(), file, uploader.UploadParams{Folder: "portfolio"})
		file.Close()
		if err == nil {
			newImageURLs = append(newImageURLs, resp.SecureURL)
		}
	}

	// Append to existing screenshots
	var newScreenshots []models.ImageScreenshot
	for _, url := range newImageURLs {
		newScreenshots = append(newScreenshots, models.ImageScreenshot{
			ProjectID: project.ID,
			ImageURL:  url,
			Title:     "",
		})
	}
	if len(newScreenshots) > 0 {
		models.DB.Create(&newScreenshots)
	}

	// Reload project with screenshots
	models.DB.Preload("Screenshots").First(&project, project.ID)

	go CleanupUnusedImages()

	c.JSON(http.StatusOK, project)
}

// UpdateProjectThumbnail handles uploading a new main thumbnail
func UpdateProjectThumbnail(c *gin.Context) {
	var project models.Project
	if err := models.DB.Preload("Screenshots").First(&project, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}

	fileHeader, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No image provided"})
		return
	}

	file, _ := fileHeader.Open()
	defer file.Close()
	cld, err := cloudinary.New()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cloudinary config error"})
		return
	}

	resp, err := cld.Upload.Upload(context.Background(), file, uploader.UploadParams{Folder: "portfolio"})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload to Cloudinary"})
		return
	}

	project.ImageURL = resp.SecureURL
	models.DB.Save(&project)

	go CleanupUnusedImages()

	c.JSON(http.StatusOK, project)
}
