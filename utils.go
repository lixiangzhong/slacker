package slacker

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	gobindata "github.com/go-bindata/go-bindata"
	"github.com/lixiangzhong/log"
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

func GoModName() string {
	b, err := ioutil.ReadFile("go.mod")
	if err == nil {
		rows := strings.Split(string(b), "\n")
		for _, row := range rows {
			cols := strings.Fields(row)
			if len(cols) == 2 {
				if cols[0] == "module" {
					return cols[1]
				}
			}
		}
	}
	log.Debug("can't read go.mod", err)
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	log.Debug("use workdir as module name:", pwd)
	return pwd
}

// func InGOPATH() bool {
// 	GOPATH, ok := os.LookupEnv("GOPATH")
// 	if !ok {
// 		log.Debug("Lookup Env GOPATH", ok)
// 	}
// 	pwd, err := os.Getwd()
// 	if err != nil {
// 		log.Debug(err)
// 	}
// 	return strings.HasPrefix(pwd, GOPATH)
// }

//for sql bindata
func ReBindata(dir string, suffix string) {
	if dir == MVCDefaultDir {
		dir = "gosrc/bindata/"
	}
	fmt.Printf("go-bindata *.%v\n", suffix)
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

func StringInSlice(s string, slice []string) bool {
	for _, v := range slice {
		if s == v {
			return true
		}
	}
	return false
}

func Quote(s string) string {
	return "`" + s + "`"
}
