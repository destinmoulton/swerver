package main

import (
	"fmt"

	"./app/lib/config"
	"./app/setup/prompts"
)

func main() {

	// fmt.Printf("The setup has been saved to .env")

	options := map[string]string{}
	confDefaults := config.LoadConfig()
	fmt.Println(confDefaults.Port)

	options["port"] = prompts.Port(confDefaults.Port)
	options["username"] = prompts.Username(confDefaults.Username)
	options["password"] = prompts.Password()
	prompts.ConfirmPassword()

	config.Save(options)
}
