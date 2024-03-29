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
	Version string = "0.1.0"
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
	utils.CustomBoolFlag(&help, "help", "Display all the command available.")

	var version bool
	utils.CustomBoolFlag(&version, "version", "Display app version.")

	var project_name string
	utils.CustomStringFlag(&project_name, "name", "Enter the name for the project.")

	var gowork bool
	flag.BoolVar(&gowork, "gowork", false, "Create a gowork environment.")

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

		var file_path string = ""
		if gowork {
			file_path = "/src"

			if err := utils.CreateScrDirectory(project_name); err != nil {
				log.Fatal(err)
			}

			templates.CreateGowork(project_name)
		}

		templates.CreateGitIgnore(project_name)
		templates.CreateReadme(project_name)
		templates.CreateMain(project_name, file_path)
		templates.CreateGomod(project_name, file_path)

		return
	}

	displayUsage()
}
