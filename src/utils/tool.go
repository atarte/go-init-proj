package utils

import (
	"errors"
	"os/exec"
	"regexp"
	"strings"
)

// isProjectNameValid check is the project name given is valid
func IsProjectNameValid(name string) bool {
	r, _ := regexp.Compile("[A-Za-z0-9_.-]")
	return r.MatchString(name)
}

// isGolangInstall check if golang is install localy on the computer
func IsGolangInstall() bool {
	out, err := exec.Command("go", "version").Output()

	if err != nil || !strings.HasPrefix(string(out), "go version") {
		return false
	}

	return true
}

// isGitInstall check if git is install localy on the computer
func IsGitInstall() bool {
	out, err := exec.Command("git", "--version").Output()

	if err != nil || !strings.HasPrefix(string(out), "git version") {
		return false
	}

	return true
}

// GetGitUsername retreve the git user name if set in the git config
func GetGitUsername() (string, error) {
	if !IsGitInstall() {
		return "", errors.New("Git is not found")
	}

	out, err := exec.Command("git", "config", "--global", "user.name").Output()

	if err != nil || string(out) == "" {
		return "", errors.New("No username found in your gitconfig")
	}

	return string(out[:len(out)-1]), nil
}

// TODO: get go version to check if the go work or go mod is availble
