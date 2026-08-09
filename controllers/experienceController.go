package controllers

import (
	"backend-porto/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetExperiences returns all experiences
func GetExperiences(c *gin.Context) {
	var experiences []models.Experience
	models.DB.Order("start_date desc").Find(&experiences)
	c.JSON(http.StatusOK, experiences)
}

// CreateExperience adds a new experience
func CreateExperience(c *gin.Context) {
	var experience models.Experience
	if err := c.ShouldBindJSON(&experience); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Create(&experience)
	c.JSON(http.StatusCreated, experience)
}

// UpdateExperience updates an existing experience
func UpdateExperience(c *gin.Context) {
	var experience models.Experience
	if err := models.DB.First(&experience, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Experience not found"})
		return
	}

	var input models.Experience
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&experience).Updates(input)
	c.JSON(http.StatusOK, experience)
}

// DeleteExperience removes an experience
func DeleteExperience(c *gin.Context) {
	var experience models.Experience
	if err := models.DB.First(&experience, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Experience not found"})
		return
	}

	models.DB.Delete(&experience)
	c.JSON(http.StatusOK, gin.H{"message": "Experience deleted successfully"})
}
