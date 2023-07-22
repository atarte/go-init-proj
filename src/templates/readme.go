package templates

import (
	"fmt"
	"os"
)

// CreateReadme create the README.md file
func CreateReadme(name string) error {
	readme_path := name + "/README.md"

	var readmeTemplate string = fmt.Sprintf("# %s", name)

	err := os.WriteFile(readme_path, []byte(readmeTemplate), 0666)
	if err != nil {
		return fmt.Errorf("Cannot create the readme.md file: %s", err)
	}

	return nil
}
