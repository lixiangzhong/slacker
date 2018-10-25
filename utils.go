package slacker

import (
	"dns.com/log"
	"os"
	"path"
	"strings"
)

func FolderName() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Debug(err)
	}
	return path.Base(dir)
}

func Mkdir(name string) error {
	return os.Mkdir(name, os.FileMode(0755))
}

//gopath project path
func ProjectPath() string {
	GOPATH, ok := os.LookupEnv("GOPATH")
	if !ok {
		log.Debug("Lookup Env GOPATH", ok)
	}
	pwd, err := os.Getwd()
	if err != nil {
		log.Debug(err)
	}
	project := strings.TrimPrefix(pwd, path.Join(GOPATH, "src")+"/")
	return path.Clean(project)
}

func InGOPATH() bool {
	GOPATH, ok := os.LookupEnv("GOPATH")
	if !ok {
		log.Debug("Lookup Env GOPATH", ok)
	}
	pwd, err := os.Getwd()
	if err != nil {
		log.Debug(err)
	}
	return strings.HasPrefix(pwd, GOPATH)
}
