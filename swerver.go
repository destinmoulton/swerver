package main

import (
	"github.com/destinmoulton/swerver/app/server"
	"github.com/destinmoulton/swerver/app/settings"
)

func main() {

	settings := settings.LoadConfig()
	server.Run(settings)
}
