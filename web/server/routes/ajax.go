package routes

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"

	"../../../lib/commander"
	"../../../lib/configparser"
)

type service struct {
	Name     string
	Response string
}

// AJAXRoutes creates the basic routes for ajax calls
func AJAXRoutes(router *gin.Engine, settings configparser.Configuration) {
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

	router.GET(prefix+"/services", func(c *gin.Context) {

		var services []service
		for _, serviceName := range settings.Services {

			response := commander.SystemCtlStatus(serviceName)
			services = append(services, service{serviceName, response})
		}
		error := ""
		c.HTML(http.StatusOK, "ajax/services.html", gin.H{
			"services": services,
			"error":    error,
		})
	})
}
