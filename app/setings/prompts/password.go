package prompts

import (
	"errors"
	"fmt"

	"github.com/manifoldco/promptui"
)

var password = ""

// Password generates a prompt for the password
func Password() string {

	prompt := promptui.Prompt{
		Label:    "Password",
		Validate: validatePassword,
		Mask:     '*',
	}

	pw, err := prompt.Run()
	password = pw

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}
	return password
}

// ConfirmPassword prompts to re type the password
func ConfirmPassword() string {
	prompt := promptui.Prompt{
		Label:    "Repeat Password",
		Validate: validateRepeatPassword,
		Mask:     '*',
	}

	_, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}
	return password
}

func validatePassword(input string) error {
	if len(input) == 0 {
		return errors.New("You must include a password")
	}
	return nil
}

func validateRepeatPassword(input string) error {
	if password != input {
		return errors.New("Password doesn't match")
	}
	return nil
}
