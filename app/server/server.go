package server

import (
	"path"

	"../lib/config"
	"./routes"
	"github.com/gin-gonic/gin"
)

// Run starts the gin server
func Run(settings config.Configuration) {

	r := gin.Default()

	glob := path.Join(settings.TemplatesPath, "**", "*.html")
	r.LoadHTMLGlob(glob)

	r.Static("/static", settings.StaticPath)

	routes.HTMLRoutes(r)
	routes.AJAXRoutes(r, settings)

	r.Run(":" + settings.Port)
}
