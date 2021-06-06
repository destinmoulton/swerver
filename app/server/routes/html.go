package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HTMLRoutes generates the homepage
func HTMLRoutes(router *gin.RouterGroup) {
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "layout/index.html", gin.H{
			"title": "Swerver",
		})
	})

}
