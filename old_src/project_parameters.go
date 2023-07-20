package src

import "path/filepath"

type ProjectType int64

const (
	HelloWorld ProjectType = iota
	BasicServer
	EmptyMain
)

func (p ProjectType) String() string {
	switch p {
	case HelloWorld:
		return "HelloWorld"
	case BasicServer:
		return "BasicServer"
	case EmptyMain:
		return "EmptyMain"
	default:
		return "Undefined"
	}
}

type ProjectParameters struct {
	Name        string
	Username    string
	Repository  string
	Path        string
	ProjectType ProjectType
}

func (p ProjectParameters) GetProjectPath() string {
	return filepath.Join(p.Path, p.Name)
}

func (p ProjectParameters) GetProjectModule() string {
	return p.Repository + "/" + p.Username + "/" + p.Name
}

func (p ProjectParameters) GetProjectFilePath(file string) string {
	return filepath.Join(p.Path, p.Name, file)
}
