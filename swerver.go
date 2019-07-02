package main

import (
	"./lib/configparser"
	"./web/server"
)

func main() {

	settings := configparser.LoadConfig("./config/config.json")
	server.Run(settings)
}
