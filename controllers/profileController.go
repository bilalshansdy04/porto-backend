package controllers

import (
	"backend-porto/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetDashboardStats returns stats for the dashboard
func GetDashboardStats(c *gin.Context) {
	var profile models.Profile
	models.DB.First(&profile) // gets the first record

	var totalProjects int64
	models.DB.Model(&models.Project{}).Count(&totalProjects)

	var recentProjects []models.Project
	models.DB.Order("date_modified desc").Limit(3).Find(&recentProjects)

	c.JSON(http.StatusOK, gin.H{
		"total_projects":      totalProjects,
		"years_of_experience": profile.YearsOfExperience,
		"summary":             profile.Summary,
		"recent_projects":     recentProjects,
	})
}

// CreateProfile handles creating the initial profile
func CreateProfile(c *gin.Context) {
	var profile models.Profile
	if err := c.ShouldBindJSON(&profile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Create(&profile)
	c.JSON(http.StatusCreated, profile)
}

// UpdateProfile updates the profile data
func UpdateProfile(c *gin.Context) {
	var profile models.Profile
	if err := models.DB.First(&profile).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Profile not found, create one first"})
		return
	}

	var input models.Profile
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&profile).Updates(input)
	c.JSON(http.StatusOK, profile)
}
