package server

import (
	"../../lib/configparser"
	"./routes"
	"github.com/gin-gonic/gin"
)

// Run starts the gin server
func Run(settings configparser.Configuration) {

	r := gin.Default()

	r.LoadHTMLGlob(settings.TemplatesGlob)

	r.Static("/assets", settings.AssetsPath)

	routes.HTMLRoutes(r)
	routes.AJAXRoutes(r, settings)

	r.Run(settings.Port)
}
