package utils

import (
	"fmt"
	"os"
)

// CreateMainDirectory create the main directory where all the file will be created
func CreateMainDirectory(name string) error {
	err := os.Mkdir(name, 0750)
	if err != nil {
		return fmt.Errorf("Cannot create the main directory: %s", err)
	}

	return nil
}

// CreateScrDirectory create the src directory in case of the user select the gowork option
func CreateScrDirectory(name string) error {
	scr_path := name + "/src"

	err := os.Mkdir(scr_path, 0750)
	if err != nil {
		return fmt.Errorf("Cannot create the src directory: %s", err)
	}

	return nil
}
