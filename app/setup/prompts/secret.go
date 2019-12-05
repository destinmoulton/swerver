package prompts

import (
	"errors"
	"fmt"

	"github.com/manifoldco/promptui"
)

// Secret generates a prompt for the crypto secret
func Secret(currentSecret string) string {

	prompt := promptui.Prompt{
		Default:  currentSecret,
		Label:    "Secret Key",
		Validate: validateSecret,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	return result
}

func validateSecret(input string) error {
	if len(input) == 0 {
		return errors.New("You must include a secret")
	}
	return nil
}
