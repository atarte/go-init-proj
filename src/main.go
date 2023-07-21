package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/atarte/go-init-proj/templates"
	"github.com/atarte/go-init-proj/utils"
)

var (
	AppName string = "go-init-proj"
	Version string = "0.0.0"
	// Build   string = "abcd"
)

func displayHelp() {
	fmt.Println("help, on vera plus tard")
}

func displayVersion() {
	fmt.Println("App:", AppName)
	fmt.Println("Version:", Version)
	fmt.Println()
}

func displayUsage() {
	fmt.Println("Usage:", AppName)
	flag.PrintDefaults()
}

func createDirectory(name string) {
	err := os.Mkdir(name, 0750)
	if err != nil {
		os.Exit(1)
	}
}

// func createGomod(name string) {
// 	if !utils.IsGolangInstall() {
// 		return
// 	}

// 	var gomod_name string
// 	git_username, err := utils.GetGitUsername()
// 	if err != nil {
// 		git_username = ""
// 		gomod_name = "github.com/" + git_username + "/"
// 	}

// 	gomod_name += name

// 	cmd := exec.Command("cd", name)
// 	if err := cmd.Run(); err != nil {
// 		// log.Fatal(err)
// 		cmd = exec.Command("go", "mod", "init", gomod_name)
// 		if err := cmd.Run(); err != nil {
// 			// log.Fatal(err)
// 		}
// 	}
// }

func test() {
	// ls
	out, err := exec.Command("ls").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))

	// ls
	cmd := exec.Command("ls")
	cmd.Dir = "src"
	out, err = cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))
}

func main() {

	var help bool
	flag.BoolVar(&help, "help", false, "Display all the command available")
	flag.BoolVar(&help, "h", false, "Display all the command available (shorthand)")

	var version bool
	flag.BoolVar(&version, "version", false, "Display app version")
	flag.BoolVar(&version, "v", false, "Display app version (shorthand)")

	var project_name string
	flag.StringVar(&project_name, "name", "", "Enter the name for the project")
	flag.StringVar(&project_name, "n", "", "Enter the name for the project")

	flag.Parse()

	if help {
		displayHelp()
		return
	}
	if version {
		displayVersion()
		return
	}
	if project_name != "" {
		if !utils.IsProjectNameValid(project_name) {
			log.Fatalln("Project name invalid!")
		}

		createDirectory(project_name)
		templates.CreateGomod(project_name)
		templates.CreateGitIgnore(project_name)
		templates.CreateMain(project_name)
		templates.CreateReadme(project_name)

		return
	}

	displayUsage()
}
