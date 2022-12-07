package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	githubRepository string = "github.com"
	gitlabRepository string = "gitlab.com"
)

var (
	FileList = []string{
		".gitignore",
		"README.md",
		"main.go",
	}
	gitignore = []byte(
		"*.exe\n/build")
	mainHelloWorld = []byte(
		"package main\n\nimport \"fmt\"\n\nfunc main() {\n\tfmt.Println(\"Hello World\")\n}")
	mainBasicServer = []byte(
		"package main\n\nimport (\"fmt\"\n\"net/http\"\n)\n\nfunc mainHandler(w http.ResponseWriter, r *http.Request) {\n\tw.Write([]byte(\"Hello World from the Web\"))\n}\n\nfunc main() {\n\thttp.HandleFunc(\"/\", mainHandler)\n\n\tfmt.Println(\"Server starting on port 8080\")\n\thttp.ListenAndServe(\":8080\", nil)\n}")
	mainEmpty = []byte(
		"package main\n\nfunc main() {\n\t\n}")
)

type ProjectType int64

const (
	HelloWorld ProjectType = iota
	BasicServer
	EmptyMain
)

func (p ProjectType) String() string {
	switch p {
	case HelloWorld:
		return "HelloWorld"
	case BasicServer:
		return "BasicServer"
	case EmptyMain:
		return "EmptyMain"
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

func (p projectParameters) getProjectPath() string {
	return filepath.Join(p.path, p.name)
}

func (p projectParameters) getProjectModule() string {
	return p.repository + "/" + p.username + "/" + p.name
}

func (p projectParameters) getProjectFilePath(file string) string {
	return filepath.Join(p.path, p.name, file)
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
	if isGitInstall() {
		return "", errors.New("Git is not installed")
	}
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
	u, err := getGitUsername()
	if err == nil {
		username = u
	}

	p := projectParameters{
		name:        projName,
		username:    username,
		repository:  githubRepository,
		path:        ".",
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
	projectPath := p.getProjectPath()
	fmt.Println(projectPath)

	// check if the project dont already exist

	// Create the main directory of the project
	if isGitInstall() {
		err := exec.Command("git", "init", projectPath).Run()
		if err != nil {
			log.Fatal("Cannot git init the project :", err)
		}
	} else {
		// os.MkdirAll()
		// Create the folder mais la j'ai la flemme
	}

	cmd := exec.Command("go", "mod", "init", p.getProjectModule())
	cmd.Dir = projectPath
	err := cmd.Run()
	if err != nil {
		log.Fatal("Cannot create go mod :", err)
	}

	for _, file := range FileList {

		content := []byte("")
		switch file {
		case ".gitignore":
			content = gitignore
		case "README.md":
			content = []byte("# " + p.name)
		case "main.go":
			switch p.projectType {
			case HelloWorld:
				content = mainHelloWorld
			case BasicServer:
				content = mainBasicServer
			case EmptyMain:
				content = mainEmpty
			}
		}

		err := os.WriteFile(p.getProjectFilePath(file), content, 0644)
		if err != nil {
			fmt.Println("Cannot create the file because :", err)
		}
	}
}

func main() {
	if !isGolangInstall() {
		fmt.Println("You need to have Golang to use this tool")
		return
	}

	if !isGitInstall() {
		fmt.Println("/!\\ It would be better for you to install git")
	}

	if len(os.Args) == 1 {
		displayMenu()
	} else if len(os.Args) == 2 {
		createWithDefaultArgs()
	} else {
		createWithArgs()
	}
}
