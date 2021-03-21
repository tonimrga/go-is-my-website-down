package main

import (
	"github.com/gin-gonic/gin"

	"is-my-website-down/controllers"
	"is-my-website-down/utils"
)

func main() {
	router := gin.Default()
	utils.ConnectDataBase()

	router.GET("/websites", controllers.GetWebsites)
	router.GET("/websites/:id", controllers.GetOneWebsite)
	router.POST("/websites", controllers.CreateWebsite)
	router.PATCH("/websites/:id", controllers.UpdateWebsite)
	router.DELETE("/websites/:id", controllers.DeleteWebsite)

	router.GET("/down-check", controllers.DownCheck)

	router.Run()
}
