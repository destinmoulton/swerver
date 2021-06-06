package server

import (
	"path"

	"github.com/destinmoulton/swerver/app/lib/config"
	"github.com/destinmoulton/swerver/app/server/routes"
	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

// Run starts the gin server
func Run(settings config.Configuration) {

	r := gin.Default()

	// Setup cookie storage setting
	store := cookie.NewStore([]byte(settings.CryptoSecret))
	r.Use(sessions.Sessions("swerver.session", store))

	glob := path.Join(settings.WebTemplatesPath, "**", "*.html")
	r.LoadHTMLGlob(glob)

	r.Static("/static", settings.WebStaticPath)

	routes.HTMLRoutes(r)
	routes.AJAXRoutes(r, settings)

	r.Run(":" + settings.Port)
}
