package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

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

func displayMenu() {
	fmt.Println("Menu")
}

func createWithDefaultArgs() {
	projName := os.Args[1]
	if !isProjectNameValid(projName) {
		log.Fatal("Project name invalid!")
	}

	username := "username"
	// if isGitInstall() {
	// 	if username, err := getGitUsername(); err != nil {
	// 		username = "username"
	// 	}
	// }

	p := projectParameters{
		name:        projName,
		username:    username,
		repository:  "github.com",
		projectType: HelloWorld,
	}

	createProject(p)
}

func createWithArgs() {

}

func createProject(p projectParameters) {

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
