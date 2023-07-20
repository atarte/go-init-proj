package templates

import (
	"log"
	"os"
)

const gitIgnoreTemplate string = `# Golang default gitignore
/build
.DS_Store

# Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary, built with 'go test -c'
*.test

# Output of the go coverage tool, specifically when used with LiteIDE
*.out

# Dependency directories (remove the comment below to include it)
# vendor/

# Go workspace file
go.work

# Editor directories and files
.idea
.vscode
*.suo
*.ntvs*
*.njsproj
*.sln
*.sw?`

func CreateGitIgnore(name string) {
	file_path := name + "/.gitignore"
	err := os.WriteFile(file_path, []byte(gitIgnoreTemplate), 0666)
	if err != nil {
		log.Fatal(err)
	}
}
