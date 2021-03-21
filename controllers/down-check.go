// controllers live-check.go

package controllers

import (
	"is-my-website-down/models"
	"is-my-website-down/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// DownCheck - GET /live-check
func DownCheck(c *gin.Context) {
	var websites []models.Website
	utils.DB.Find(&websites)

	channel := make(chan string)
	websiteStatuses := []string{}

	for _, website := range websites {
		go utils.IsWebsiteDown(website, channel)
	}

	for i := 0; i < len(websites); i++ {
		w := <-channel
		if w != "" {
			websiteStatuses = append(websiteStatuses, w)
		}
	}

	c.JSON(http.StatusOK, websiteStatuses)
}
