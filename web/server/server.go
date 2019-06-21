package server

import "github.com/gin-gonic/gin"

// Run starts the gin server
func Run() {

	r := gin.Default()

	r.Static("/assets", "../assets")
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
