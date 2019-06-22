package main

import (
	"./lib/config"
	"./web/server"
)

func main() {

	settings := config.LoadConfig("./config/config.json")
	server.Run(settings)
}
