// controllers books.go

package controllers

import (
	"net/http"

	"is-my-website-down/models"

	"github.com/gin-gonic/gin"
)

// GetWebsites - GET /websites
func GetWebsites(c *gin.Context) {
	var websites []models.Website
	models.DB.Find(&websites)

	c.JSON(http.StatusOK, gin.H{"websites": websites})
}

// CreateWebsite - POST /websites
func CreateWebsite(c *gin.Context) {
	var input models.CreateWebsiteInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	website := models.Website{Name: input.Name, URL: input.URL}
	models.DB.Create(&website)

	c.JSON(http.StatusOK, gin.H{"website": website})
}

// GetOneWebsite - GET /website/:id
func GetOneWebsite(c *gin.Context) {
	var website models.Website

	if err := models.DB.Where("id = ?", c.Param("id")).First(&website).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Website not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"website": website})
}

// UpdateWebsite - PATCH /website/:id
func UpdateWebsite(c *gin.Context) {
	var website models.Website
	if err := models.DB.Where("id = ?", c.Param("id")).First(&website).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Website not found!"})
		return
	}

	var input models.UpdateWebsiteInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&website).Updates(input)

	c.JSON(http.StatusOK, gin.H{"website": website})
}

// DeleteWebsite - DELETE - /website/:id
func DeleteWebsite(c *gin.Context) {
	// Get model if exist
	var website models.Website
	if err := models.DB.Where("id = ?", c.Param("id")).First(&website).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Website not found!"})
		return
	}

	models.DB.Delete(&website)

	c.JSON(http.StatusOK, gin.H{"website": website})
}

// LiveCheck - GET /live-check
func LiveCheck(c *gin.Context) {
	var websites []models.Website
	models.DB.Find(&websites)

	channel := make(chan string)
	websiteStatuses := []string{}

	for _, website := range websites {
		go IsWebsiteLive(website, channel)
	}

	for i := 0; i < len(websites); i++ {
		websiteStatuses = append(websiteStatuses, <-channel)
	}

	c.JSON(http.StatusOK, websiteStatuses)
}

func IsWebsiteLive(website models.Website, c chan string) {
	_, err := http.Get(website.URL)
	if err != nil {
		c <- website.Name + " is DOWN!"
		return
	}

	c <- website.Name + " is OK."
}
