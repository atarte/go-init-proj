package src

import (
	"errors"
	"os/exec"
	"strings"
)

const (
	GithubRepository string = "github.com"
	GitlabRepository string = "gitlab.com"
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

// isGitInstall check if git is install localy on the computer
func isGitInstall() bool {
	out, err := exec.Command("git", "--version").Output()

	if err != nil || !strings.HasPrefix(string(out), "git version") {
		return false
	}

	return true
}

// isGolangInstall check if golang is install localy on the computer
func isGolangInstall() bool {
	out, err := exec.Command("go", "version").Output()

	if err != nil || !strings.HasPrefix(string(out), "go version") {
		return false
	}

	return true
}

// AreToolsInstall check if every required tools for this application is installed, if not it send a error message
func AreToolsInstall() (bool, error) {
	// can easely add other tools required if needed
	switch false {
	case isGitInstall():
		return false, errors.New("Git not installed")
	case isGolangInstall():
		return false, errors.New("Golang not installed")
	default:
		return true, nil
	}
}

// GetGitUsername retreve the git user name if set in the git config
func GetGitUsername() (string, error) {
	if isGitInstall() {
		return "", errors.New("Git is not installed")
	}

	out, err := exec.Command("git", "config", "--global", "user.name").Output()

	if err != nil || string(out) == "" {
		return "", errors.New("No username configure")
	}

	return string(out[:len(out)-1]), nil
}
