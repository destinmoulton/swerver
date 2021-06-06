package server

import (
	"path"

	"github.com/destinmoulton/swerver/app/lib/pw"
	"github.com/destinmoulton/swerver/app/server/routes"
	"github.com/destinmoulton/swerver/app/settings"
	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

// Run starts the gin server
func Run(settings settings.Configuration) {

	r := gin.Default()

	// Setup cookie storage setting
	store := cookie.NewStore([]byte(settings.CryptoSecret))
	r.Use(sessions.Sessions("swerver.session", store))

	glob := path.Join(settings.WebTemplatesPath, "**", "*.html")
	r.LoadHTMLGlob(glob)

	r.Static("/static", settings.WebStaticPath)

	username := settings.Username
	password := pw.DecryptPassword(settings.CryptoSecret, settings.Password)
	authorized := r.Group("", gin.BasicAuth(gin.Accounts{username: password}))

	routes.HTMLRoutes(authorized)
	routes.AJAXRoutes(authorized, settings)

	r.Run(":" + settings.Port)
}
