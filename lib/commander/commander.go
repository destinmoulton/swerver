package commander

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"../configparser"
)

// Run a shell command
func Run(parts ...string) (string, error) {

	cmd := exec.Command(parts[0], parts[1:]...)

	out, err := cmd.CombinedOutput()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			return "", fmt.Errorf("Error running command. Exited properly. %v", exitErr)
		}
		return "", fmt.Errorf("Error running command. Exited improperly. %v", err)
	}
	return string(out), nil
}

// SystemCtlStatus gets the status for a systemctl status
func SystemCtlStatus(service string) string {
	resp, err := Run("systemctl", "show", "-p", "ActiveState", service)

	if err != nil {
		return "ERROR"
	}
	return strings.Split(strings.TrimRight(resp, "\n"), "=")[1]
}

// RunScript runs the specified script
func RunScript(settings configparser.Configuration, scriptToRun string) (string, error) {

	path := settings.ScriptsPath + scriptToRun
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return "", fmt.Errorf("Script `%v` does not exist", path)
	}
	resp, err := Run("sh", "-c", path)

	if err != nil {
		return "", err
	}
	return resp, nil
}
