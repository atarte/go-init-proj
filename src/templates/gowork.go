package templates

import (
	"fmt"
	"os/exec"
)

// CreateGowork create a go.work file
func CreateGowork(name string) error {
	cmd := exec.Command("go", "work", "init", "src")
	cmd.Dir = name
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("Cannot create the go.work file: %s", err)
	}

	return nil
}
