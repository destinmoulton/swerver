package server

import (
	"../../lib/configparser"
	"./routes"
	"github.com/gin-gonic/gin"
)

// Run starts the gin server
func Run(settings configparser.Configuration) {

	r := gin.Default()

	r.LoadHTMLGlob("./web/templates/**/*")

	r.Static("/assets", settings.AssetsPath)

	routes.HTMLRoutes(r)
	routes.AJAXRoutes(r)

	r.Run(settings.Port)
}
