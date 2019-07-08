package commander

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"../config"
)

// Run a shell command
func Run(parts ...string) (string, error) {

	cmd := exec.Command(parts[0], parts[1:]...)

	out, err := cmd.CombinedOutput()
	if err != nil {
		if _, ok := err.(*exec.ExitError); ok {

			return "", fmt.Errorf("%v", string(out))
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

// SystemCtlCommand runs a systemctl command on a service
func SystemCtlCommand(service string, command string) error {

	_, err := Run("sudo", "systemctl", command, service)

	if err != nil {
		return err
	}

	return nil
}

// RunScript runs the specified script
func RunScript(settings config.Configuration, scriptToRun string) (string, error) {

	path := filepath.Join(settings.ScriptsPath, scriptToRun)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return "", fmt.Errorf("Script `%v` does not exist", path)
	}
	resp, err := Run("sh", "-c", path)

	if err != nil {
		return "", err
	}
	return resp, nil
}
