package server

import "github.com/gin-gonic/gin"
import "../../lib/config"

// Run starts the gin server
func Run(settings config.Configuration) {

	r := gin.Default()

	r.Static("/assets", settings.AssetsPath)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(settings.Port)
}
