package templates

import (
	"fmt"
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
go.work.sum

# Editor directories and files
.idea
.vscode
*.suo
*.ntvs*
*.njsproj
*.sln
*.sw?`

// CreateGitIgnore create the .gitignore file
func CreateGitIgnore(name string) error {
	gitignore_path := name + "/.gitignore"

	err := os.WriteFile(gitignore_path, []byte(gitIgnoreTemplate), 0666)
	if err != nil {
		return fmt.Errorf("Cannot create the .gitignore file: %s", err)
	}

	return nil
}
