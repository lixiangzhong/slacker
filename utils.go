package slacker

import (
	"dns.com/log"
	"fmt"
	gobindata "github.com/go-bindata/go-bindata"
	"os"
	"path"
	"path/filepath"
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

//for sql bindata
func ReBindata(dir string, suffix string) {
	if dir == MVCDefaultDir {
		dir = "gosrc/bindata/"
	}
	fmt.Println("Re go-bindata...")
	cfg := gobindata.NewConfig()
	cfg.Output = filepath.Join(dir, "bindata.go")
	cfg.Package = "bindata"
	cfg.Prefix = dir
	files := hasSuffixFiles(dir, suffix)
	for _, file := range files {
		cfg.Input = append(cfg.Input, gobindata.InputConfig{
			Path: file,
		})
	}
	err := gobindata.Translate(cfg)
	if err != nil {
		log.Error(err)
	}
}

func hasSuffixFiles(dir, suffix string) []string {
	var files = make([]string, 0)
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() && path != dir {
			return filepath.SkipDir
		}
		if strings.HasSuffix(path, suffix) {
			files = append(files, path)
		}
		return err
	})
	return files
}
