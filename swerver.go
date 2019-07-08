package main

import (
	"./lib/config"
	"./web/server"
)

func main() {

	settings := config.LoadConfig()
	server.Run(settings)
}
