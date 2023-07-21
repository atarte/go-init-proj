package templates

import (
	"fmt"
	"os"
)

const mainTemplate string = `package main

import "fmt"

func main() {
	fmt.Println("Hello world")
}`

// CreateMain create the main.go file
func CreateMain(name string) error {
	main_path := name + "/main.go"

	err := os.WriteFile(main_path, []byte(mainTemplate), 0666)
	if err != nil {
		return fmt.Errorf("Cannot create the main.go file: %s", err)
	}

	return nil
}
