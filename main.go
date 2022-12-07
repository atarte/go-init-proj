package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
)

var FileList = []string{
	".gitignore",
	"README.md",
	"main.go",
}

type ProjectType int64

const (
	HelloWorld ProjectType = iota
	BasicServer
)

func (p ProjectType) String() string {
	switch p {
	case HelloWorld:
		return "HelloWorld"
	case BasicServer:
		return "BasicServer"
	default:
		return "Undefined"
	}
}

type projectParameters struct {
	name        string
	username    string
	repository  string
	path        string
	projectType ProjectType
}

func isGitInstall() bool {
	out, err := exec.Command("git", "--version").Output()
	if err != nil || !strings.HasPrefix(string(out), "git version") {
		return false
	}
	return true
}

func isGolangInstall() bool {
	out, err := exec.Command("go", "version").Output()
	if err != nil || !strings.HasPrefix(string(out), "go version") {
		return false
	}
	return true
}

func isProjectNameValid(name string) bool {
	r, _ := regexp.Compile("[A-Za-z0-9_.-]")
	return r.MatchString(name)
}

func getGitUsername() (string, error) {
	out, err := exec.Command("git", "config", "--global", "user.name").Output()
	if err != nil || string(out) == "" {
		return "", errors.New("No username configure")
	}
	return string(out[:len(out)-1]), nil
}

func getDefaultPath() string {
	os := runtime.GOOS
	switch os {
	case "windows":
		return ".\\"
	case "darwin":
		return ".//"
	case "linux":
		return ".//"
	default:
		return ".//"
	}
}

func displayMenu() {
	fmt.Println("Menu")
}

func createWithDefaultArgs() {
	projName := os.Args[1]
	if !isProjectNameValid(projName) {
		log.Fatal("Project name invalid!")
	}

	username := "username"
	if isGitInstall() {
		u, err := getGitUsername()
		if err == nil {
			username = u
		}
	}

	p := projectParameters{
		name:        projName,
		username:    username,
		repository:  "github.com",
		path:        getDefaultPath(),
		projectType: HelloWorld,
	}

	createProject(p)
}

func createWithArgs() {
	/*
		--help -h
		--name -n
		--username -u
		--path -p
		--repository -r
		--type -t
	*/
}

func createProject(p projectParameters) {
	projectPath := filepath.Join(p.path, p.name)

	// Create the main directory of the project
	if isGitInstall() {
		err := exec.Command("git", "init", projectPath)
		if err != nil {
			log.Fatal("Cannot git init the project")
		}
	} else {
		// Create the folder mais la j'ai la flemme
	}

}

func main() {
	// args := os.Args[1:]
	// fmt.Println(args)

	// git := GitIsInstall()
	// fmt.Println(git)

	// golang := GolangIsInstall()
	// fmt.Println(golang)

	if len(os.Args) == 1 {
		displayMenu()
	} else if len(os.Args) == 2 {

	} else {
		createWithArgs()
	}

	// username, err := GetGitUsername()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// // fmt.Println(username)
	// fmt.Printf("[%v]", username)

	// param := projectParameters{
	// 	name:        "test",
	// 	username:    "atarte",
	// 	ProjectType: HelloWorld,
	// }

	// fmt.Println(param.ProjectType == HelloWorld)
	// fmt.Println(param.ProjectType == BasicServer)

}
