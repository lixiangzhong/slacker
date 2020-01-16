package bindata

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"math/rand"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"golang.org/x/tools/imports"
)

// RestoreAsset restores an asset under the given directory
func restorefile(dir, name string, tpldata interface{}) error {

	data, err := Asset(name)
	if err != nil {
		return err
	}

	info, err := AssetInfo(name)
	if err != nil {
		return err
	}

	name = strings.Replace(name, "example_project", folderName(), -1)

	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}

	//exec template
	data, err = exectemplate(name, data, tpldata)
	if err != nil {
		return err
	}

	filename := _filePath(dir, name)

	err = ioutil.WriteFile(filename, data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(filename, info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

func Restore(tpldata interface{}) error {
	return restore("", "", tpldata)
}

func restore(dir, name string, tpldata interface{}) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return restorefile(dir, name, tpldata)
	}
	// Dir
	for _, child := range children {
		err = restore(dir, filepath.Join(name, child), tpldata)
		if err != nil {
			return err
		}
	}
	return nil
}

func folderName() string {
	dir, _ := os.Getwd()
	return path.Base(dir)
}

func exectemplate(name string, data []byte, tpldata interface{}) ([]byte, error) {
	var tpl = new(template.Template)
	var buff = new(bytes.Buffer)
	var b []byte

	_, err := tpl.Parse(string(data))
	if err != nil {
		return data, err
	}
	err = tpl.Execute(buff, tpldata)
	if err != nil {
		return buff.Bytes(), err
	}
	ext := path.Ext(name)
	switch ext {
	case ".go":
		b, err = imports.Process("", buff.Bytes(), nil)
		if err != nil {
			return b, err
		}
		b, err = format.Source(b)
		if err != nil {
			return b, err
		}
	case ".vue":
		b = buff.Bytes()
		b = bytes.Replace(b, []byte("{vue"), []byte("{{"), -1)
		b = bytes.Replace(b, []byte("vue}"), []byte("}}"), -1)
	case ".http":
		b = buff.Bytes()
		b = bytes.Replace(b, []byte("{http"), []byte("{{"), -1)
		b = bytes.Replace(b, []byte("http}"), []byte("}}"), -1)
		b = bytes.Replace(b, []byte("{RandPort}"), []byte(RandPort()), -1)
	case ".yaml", ".development", ".json":
		b = buff.Bytes()
		b = bytes.Replace(b, []byte("{RandPort}"), []byte(RandPort()), -1)
	default:
		b = buff.Bytes()
	}
	return b, nil
}

var randport int

func RandPort() string {
	if randport > 0 {
		return fmt.Sprintf("%v", randport)
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randport = r.Intn(20000) + 10000
	return fmt.Sprintf("%v", randport)
}
