package routes

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"../../../lib/commander"
	"../../../lib/configparser"
)

type service struct {
	Name           string
	ResponseString string
	IsActive       bool
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
		error := ""
		for _, serviceName := range settings.Services {

			response := commander.SystemCtlStatus(serviceName)
			error = response
			isActive := false
			if response == "active" {
				isActive = true
			}
			services = append(services, service{serviceName, response, isActive})
		}
		c.HTML(http.StatusOK, "ajax/services.html", gin.H{
			"services": services,
			"error":    error,
		})
	})

	router.GET(prefix+"/scripts", func(c *gin.Context) {

		c.HTML(http.StatusOK, "ajax/scripts.html", gin.H{
			"scripts": settings.Scripts,
			"path":    settings.ScriptsPath,
		})
	})

	router.GET(prefix+"/run-script", func(c *gin.Context) {
		script := c.Query("script")

		output, err := commander.RunScript(settings, script)

		isError := false
		var lines []string
		if err != nil {
			isError = true

			lines = strings.Split(err.Error(), "\n")
		} else {

			lines = strings.Split(output, "\n")
		}
		c.HTML(http.StatusOK, "ajax/tty.html", gin.H{
			"lines":   lines,
			"isError": isError,
		})

	})

	router.GET(prefix+"/memory-usage", func(c *gin.Context) {

		output, err := commander.Run("free", "-m")

		if err != nil {

		}

		memline := strings.Split(output, "\n")[1]
		memparts := strings.Fields(memline)

		c.HTML(http.StatusOK, "ajax/memory.html", gin.H{

			"total":     memparts[1],
			"used":      memparts[2],
			"free":      memparts[3],
			"buffcache": memparts[5],
			"available": memparts[6],
		})
	})

	router.GET(prefix+"/sysinfo", func(c *gin.Context) {

		output, err := commander.Run("uptime")

		if err != nil {

		}

		parts := strings.Fields(output)

		c.HTML(http.StatusOK, "ajax/sysinfo.html", gin.H{

			"daysup":         parts[2],
			"timeup":         strings.TrimRight(parts[4], ","),
			"oneminload":     strings.TrimRight(parts[9], ","),
			"fiveminload":    strings.TrimRight(parts[10], ","),
			"fifteenminload": strings.TrimRight(parts[11], ","),
		})
	})
}
