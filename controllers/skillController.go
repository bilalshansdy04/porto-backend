package controllers

import (
	"backend-porto/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetSkills returns all skills
func GetSkills(c *gin.Context) {
	var skills []models.Skill
	models.DB.Find(&skills)
	c.JSON(http.StatusOK, skills)
}

// CreateSkill adds a new skill
func CreateSkill(c *gin.Context) {
	var skill models.Skill
	if err := c.ShouldBindJSON(&skill); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Create(&skill)
	c.JSON(http.StatusCreated, skill)
}

// UpdateSkill updates an existing skill
func UpdateSkill(c *gin.Context) {
	var skill models.Skill
	if err := models.DB.First(&skill, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Skill not found"})
		return
	}

	var input models.Skill
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&skill).Updates(input)
	c.JSON(http.StatusOK, skill)
}

// DeleteSkill removes a skill
func DeleteSkill(c *gin.Context) {
	var skill models.Skill
	if err := models.DB.First(&skill, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Skill not found"})
		return
	}

	models.DB.Delete(&skill)
	c.JSON(http.StatusOK, gin.H{"message": "Skill deleted successfully"})
}
