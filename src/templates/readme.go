package templates

import (
	"fmt"
	"log"
	"os"
)

func CreateReadme(name string) {
	file_path := name + "/.gitignore"

	var readmeTemplate string = fmt.Sprintf("# %s", name)

	err := os.WriteFile(file_path, []byte(readmeTemplate), 0666)
	if err != nil {
		log.Fatal(err)
	}
}
