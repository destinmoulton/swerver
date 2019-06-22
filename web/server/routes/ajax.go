package routes

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AJAXRoutes creates the basic routes for ajax calls
func AJAXRoutes(router *gin.Engine) {
	prefix := "/ajax"
	router.GET(prefix+"/ip", func(c *gin.Context) {
		resp, err := http.Get("https://ipinfo.io/ip")
		defer resp.Body.Close()
		error := ""
		ip := ""
		if err != nil {
			error = "Unable to get IP Address."
		} else {
			body, rerr := ioutil.ReadAll(resp.Body)
			if rerr != nil {
				error = "Error parsing IP address response."
			} else {
				ip = string(body)
			}
		}
		c.HTML(http.StatusOK, "ajax/ip.html", gin.H{
			"ip":    ip,
			"error": error,
		})
	})
}
