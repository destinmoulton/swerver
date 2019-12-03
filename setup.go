package main

import (
	"errors"
	"fmt"

	"github.com/manifoldco/promptui"
)

func main() {

	port := promptPort()
	password := promptPassword()

	fmt.Printf("The setup has been saved to .env")

}

func promptPort() string {

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
func promptPassword() string {

	prompt := promptui.Prompt{
		Label:    "Password",
		Validate: validatePassword,
		Mask:     '*',
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}
	return result
}
func validatePassword(input string) error {
	if len(input) == 0 {
		return errors.New("You must include a password")
	}
	return nil
}
