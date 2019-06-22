package server

import (
	"../../lib/config"
	"./routes"
	"github.com/gin-gonic/gin"
)

// Run starts the gin server
func Run(settings config.Configuration) {

	r := gin.Default()

	r.LoadHTMLGlob("./web/templates/*")
	r.Static("/assets", settings.AssetsPath)
	routes.HTMLRoutes(r)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(settings.Port)
}
