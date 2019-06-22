package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HTMLRoutes generates the homepage
func HTMLRoutes(router *gin.Engine) {
	router.LoadHTMLGlob("./web/templates/*")
	router.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Swerver",
		})
	})

}
