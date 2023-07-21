package main

import (
	"flag"
	"fmt"
	"log"

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
			log.Fatalln("Project name invalid! It must follow the `[A-Za-z0-9_.-]` regex.")
		}

		if err := utils.CreateMainDirectory(project_name); err != nil {
			log.Fatal(err)
		}
		templates.CreateGomod(project_name)
		templates.CreateGitIgnore(project_name)
		templates.CreateMain(project_name)
		templates.CreateReadme(project_name)

		return
	}

	displayUsage()
}
