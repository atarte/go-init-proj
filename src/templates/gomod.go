package templates

import (
	"log"
	"os/exec"

	"github.com/atarte/go-init-proj/utils"
)

func CreateGomod(name string) {
	username, err := utils.GetGitUsername()
	if err == nil {
		username = "github.com/" + username + "/"
		// log.Fatal(err)
	}

	project_name := username + name

	cmd := exec.Command("go", "mod", "init", project_name)
	cmd.Dir = name
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
