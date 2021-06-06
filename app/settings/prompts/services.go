package prompts

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

// Services generates a prompt for the default port
func Services(currentServices string) string {

	prompt := promptui.Prompt{
		Default: currentServices,
		Label:   "Comma Separated Linux Services",
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	return result
}
