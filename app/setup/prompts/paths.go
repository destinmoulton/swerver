package prompts

import (
	"errors"
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
)

// ScriptsPath generates a prompt for the default port
func ScriptsPath(initialScriptPath string) string {

	prompt := promptui.Prompt{
		Default:  initialScriptPath,
		Label:    "ScriptsPath",
		Validate: validatePath,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	return result
}

// WebPath generates a prompt for the default port
func WebPath(initialPath string) string {

	prompt := promptui.Prompt{
		Default:  initialPath,
		Label:    "WebPath",
		Validate: validatePath,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	return result
}

func validatePath(input string) error {
	if len(input) == 0 {
		return errors.New("You must include a valid path")
	}

	_, err := os.Stat(input)
	if os.IsNotExist(err) {
		return errors.New("Directory doesn't exist")
	}
	return nil
}
