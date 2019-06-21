package commander

import (
	"fmt"
	"os"
	"os/exec"
)

// Run a shell command
func Run(parts ...string) {

	cmd := exec.Command(parts[0], parts[0:]...)

	out, err := cmd.CombinedOutput()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			fmt.Printf("finished with non-zero: %v\n", exitErr)
		} else {
			fmt.Printf("failed to run: %v", err)
			os.Exit(1)
		}
	}
	fmt.Printf("Status is: %s\n", string(out))

}
