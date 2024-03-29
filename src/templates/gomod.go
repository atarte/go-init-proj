package templates

import (
	"fmt"
	"os/exec"

	"github.com/atarte/go-init-proj/utils"
)

// CreateGomod create a go.mod file
func CreateGomod(name string, path string) error {
	username, err := utils.GetGitUsername()
	if err == nil {
		username = "github.com/" + username + "/"
	}

	project_name := username + name

	cmd := exec.Command("go", "mod", "init", project_name)
	cmd.Dir = name + path
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("Cannot create the go.mod file: %s", err)
	}

	return nil
}
