package prompts

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

// Port generates a prompt for the default port
func Port() string {

	prompt := promptui.Prompt{
		Label: "Port (default is 9090)",
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	return result
}
