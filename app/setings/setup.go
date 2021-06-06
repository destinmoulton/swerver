package setup

import (
	"github.com/destinmoulton/swerver/app/lib/config"
	"github.com/destinmoulton/swerver/app/lib/pw"
	"github.com/destinmoulton/swerver/app/setup/prompts"
)

func PromptConfig() {

	options := map[string]string{}

	options["port"] = prompts.Port(config.GetSingle("port"))
	options["scripts_path"] = prompts.ScriptsPath(config.GetSingle("scripts_path"))
	options["web_path"] = prompts.WebPath(config.GetSingle("web_path"))
	options["services_to_monitor"] = prompts.Services(config.GetSingle("services_to_monitor"))
	options["crypto_secret"] = prompts.Secret(config.GetSingle("crypto_secret"))
	options["username"] = prompts.Username(config.GetSingle("username"))
	password := prompts.Password()
	if prompts.ConfirmPassword() != "" {
		options["password"] = pw.GenerateHash(password)
	}

	config.Save(options)
}
