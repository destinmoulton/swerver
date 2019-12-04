package main

import (
	"./app/lib/config"
	"./app/setup/prompts"
)

func main() {

	// fmt.Printf("The setup has been saved to .env")

	options := map[string]string{}

	options["port"] = prompts.Port(config.GetSingle("port"))
	options["scripts_path"] = prompts.ScriptsPath(config.GetSingle("scripts_path"))
	options["web_path"] = prompts.WebPath(config.GetSingle("web_path"))
	options["services_to_monitor"] = prompts.Services(config.GetSingle("services_to_monitor"))
	options["username"] = prompts.Username(config.GetSingle("username"))
	options["password"] = prompts.Password()
	prompts.ConfirmPassword()

	config.Save(options)
}
