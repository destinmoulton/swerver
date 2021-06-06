package main

import (
	"github.com/destinmoulton/swerver/app/lib/config"
	"github.com/destinmoulton/swerver/app/server"
)

func main() {

	settings := config.LoadConfig()
	server.Run(settings)
}
