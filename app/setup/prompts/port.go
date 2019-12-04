package prompts

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

// Port generates a prompt for the default port
func Port(defaultPort string) string {

	prompt := promptui.Prompt{
		Default: defaultPort,
		Label:   "Port",
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	return result
}
