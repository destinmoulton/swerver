package prompts

import (
	"errors"
	"fmt"

	"github.com/manifoldco/promptui"
)

// Port generates a prompt for the default port
func Port(defaultPort string) string {

	prompt := promptui.Prompt{
		Default:  defaultPort,
		Label:    "Port",
		Validate: validatePort,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	return result
}

func validatePort(input string) error {
	if len(input) == 0 {
		return errors.New("You must include a port")
	}
	return nil
}
