package settings

import (
	"github.com/destinmoulton/swerver/app/lib/pw"
	"github.com/destinmoulton/swerver/app/settings/prompts"
)

func PromptConfig() {

	options := map[string]string{}

	options["port"] = prompts.Port(GetSingleConfigValue("port"))
	options["scripts_path"] = prompts.ScriptsPath(GetSingleConfigValue("scripts_path"))
	options["web_path"] = prompts.WebPath(GetSingleConfigValue("web_path"))
	options["services_to_monitor"] = prompts.Services(GetSingleConfigValue("services_to_monitor"))
	options["crypto_secret"] = prompts.Secret(GetSingleConfigValue("crypto_secret"))
	options["username"] = prompts.Username(GetSingleConfigValue("username"))
	password := prompts.Password()
	if prompts.ConfirmPassword() != "" {
		options["password"] = pw.GenerateHash(password)
	}

	SaveConfigToFile(options)
}
