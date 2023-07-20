package templates

import (
	"log"
	"os"
)

const mainTemplate string = `package main

import "fmt"

func main() {
	fmt.Println("Hello world")
}`

func CreateMain(name string) {
	file_path := name + "/main.go"
	err := os.WriteFile(file_path, []byte(mainTemplate), 0666)
	if err != nil {
		log.Fatal(err)
	}
}
