package main

import (
	"github.com/gin-gonic/gin"

	"is-my-website-down/controllers"
	"is-my-website-down/models"
)

func main() {
	router := gin.Default()
	models.ConnectDataBase()

	router.GET("/websites", controllers.GetWebsites)
	router.GET("/websites/:id", controllers.GetOneWebsite)
	router.POST("/websites", controllers.CreateWebsite)
	router.PATCH("/websites/:id", controllers.UpdateWebsite)
	router.DELETE("/websites/:id", controllers.DeleteWebsite)

	router.GET("/live-check", controllers.LiveCheck)

	router.Run()
}
