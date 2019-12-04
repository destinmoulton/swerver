package prompts

import (
	"errors"
	"fmt"

	"github.com/manifoldco/promptui"
)

// Username generates a prompt for the default port
func Username(currentUsername string) string {

	prompt := promptui.Prompt{
		Default:  currentUsername,
		Label:    "Username",
		Validate: validateUsername,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	return result
}

func validateUsername(input string) error {
	if len(input) == 0 {
		return errors.New("You must include a username")
	}
	return nil
}
