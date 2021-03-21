// controllers websites.go

package controllers

import (
	"is-my-website-down/models"
	"is-my-website-down/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetWebsites - GET /websites
func GetWebsites(c *gin.Context) {
	var websites []models.Website
	utils.DB.Find(&websites)

	c.JSON(http.StatusOK, gin.H{"websites": websites})
}

// CreateWebsite - POST /websites
func CreateWebsite(c *gin.Context) {
	var input models.CreateWebsiteInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !utils.IsURL(input.URL) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Website URL not valid!"})
		return
	}

	website := models.Website{Name: input.Name, URL: input.URL}
	utils.DB.Create(&website)

	c.JSON(http.StatusOK, gin.H{"website": website})
}

// GetOneWebsite - GET /website/:id
func GetOneWebsite(c *gin.Context) {
	var website models.Website

	if err := utils.DB.Where("id = ?", c.Param("id")).First(&website).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Website not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"website": website})
}

// UpdateWebsite - PATCH /website/:id
func UpdateWebsite(c *gin.Context) {
	var website models.Website
	if err := utils.DB.Where("id = ?", c.Param("id")).First(&website).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Website not found!"})
		return
	}

	var input models.UpdateWebsiteInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !utils.IsURL(input.URL) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Website URL not valid!"})
		return
	}

	utils.DB.Model(&website).Updates(input)

	c.JSON(http.StatusOK, gin.H{"website": website})
}

// DeleteWebsite - DELETE - /website/:id
func DeleteWebsite(c *gin.Context) {
	var website models.Website
	if err := utils.DB.Where("id = ?", c.Param("id")).First(&website).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Website not found!"})
		return
	}

	utils.DB.Delete(&website)

	c.JSON(http.StatusOK, gin.H{"website": website})
}
