package main

import (
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/atarte/go-init-proj/src"
)

func isProjectNameValid(name string) bool {
	r, _ := regexp.Compile("[A-Za-z0-9_.-]")
	return r.MatchString(name)
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
	u, err := src.GetGitUsername()
	if err == nil {
		username = u
	}

	p := src.ProjectParameters{
		Name:        projName,
		Username:    username,
		Repository:  src.GithubRepository,
		Path:        ".",
		ProjectType: src.HelloWorld,
	}

	fmt.Println(p)

	// createProject(p)
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

// func createProject(p src.ProjectParameters) {
// 	projectPath := p.getProjectPath()
// 	fmt.Println(projectPath)

// 	// check if the project dont already exist

// 	// Create the main directory of the project
// 	if src.isGitInstall() {
// 		err := exec.Command("git", "init", projectPath).Run()
// 		if err != nil {
// 			log.Fatal("Cannot git init the project :", err)
// 		}
// 	} else {
// 		// os.MkdirAll()
// 		// Create the folder mais la j'ai la flemme
// 	}

// 	cmd := exec.Command("go", "mod", "init", p.getProjectModule())
// 	cmd.Dir = projectPath
// 	err := cmd.Run()
// 	if err != nil {
// 		log.Fatal("Cannot create go mod :", err)
// 	}

// 	for _, file := range FileList {

// 		content := []byte("")
// 		switch file {
// 		case ".gitignore":
// 			content = gitignore
// 		case "README.md":
// 			content = []byte("# " + p.name)
// 		case "main.go":
// 			switch p.projectType {
// 			case HelloWorld:
// 				content = mainHelloWorld
// 			case BasicServer:
// 				content = mainBasicServer
// 			case EmptyMain:
// 				content = mainEmpty
// 			}
// 		}

// 		err := os.WriteFile(p.getProjectFilePath(file), content, 0644)
// 		if err != nil {
// 			fmt.Println("Cannot create the file because :", err)
// 		}
// 	}
// }

func main() {
	// if !isGolangInstall() {
	// 	fmt.Println("You need to have Golang to use this tool")
	// 	return
	// }
	//
	// if !isGitInstall() {
	// 	fmt.Println("/!\\ It would be better for you to install git")
	// }
	//
	// if !src.AreToolsInstall() {

	// }

	fmt.Println("arg", os.Args)

	if len(os.Args) == 1 {
		displayMenu()
	} else if len(os.Args) == 2 {
		createWithDefaultArgs()
	} else {
		createWithArgs()
	}
}
