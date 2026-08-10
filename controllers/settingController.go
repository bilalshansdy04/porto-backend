package controllers

import (
	"backend-porto/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetSettings retrieves the global settings. If it doesn't exist, it creates a default one.
func GetSettings(c *gin.Context) {
	var setting models.Setting
	
	// Try to get the first record
	result := models.DB.First(&setting)
	
	// If no record exists, create default settings
	if result.Error != nil {
		setting = models.Setting{
			ShowExperience: true,
			Language:       "id",
		}
		if err := models.DB.Create(&setting).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create default settings"})
			return
		}
	}

	c.JSON(http.StatusOK, setting)
}

// UpdateSettings updates the global settings
func UpdateSettings(c *gin.Context) {
	var input models.Setting
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var setting models.Setting
	if err := models.DB.First(&setting).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Settings not found"})
		return
	}

	models.DB.Model(&setting).Select("ShowExperience", "Language").Updates(input)

	c.JSON(http.StatusOK, setting)
}
