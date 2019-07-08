package main

import (
	"./app/lib/config"
	"./app/server"
)

func main() {

	settings := config.LoadConfig()
	server.Run(settings)
}
