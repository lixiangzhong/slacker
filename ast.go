package slacker

import (
	"bytes"
	"dns.com/log"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"golang.org/x/tools/go/ast/astutil"
	"io/ioutil"
	"os"
)

var (
	fset = token.NewFileSet()
)

func addInitCreateTableFile(table string) {
	var b []byte //backup old file
	file := "gosrc/models/init.go"
	fmt.Println("添加create table -> init.go")
	fi, err := os.Lstat(file)
	if err != nil {
		log.Error(err)
		return
	}
	b, err = ioutil.ReadFile(file)
	if err != nil {
		log.Error(err)
		return
	}
	f, err := parser.ParseFile(fset, file, nil, 0)
	if err != nil {
		log.Error(err)
		return
	}
	astutil.Apply(f, addTable(table), nil)
	var buff = new(bytes.Buffer)
	err = printer.Fprint(buff, fset, f)
	if err != nil {
		log.Error(err)
		return
	}
	err = ioutil.WriteFile(file, buff.Bytes(), fi.Mode().Perm())
	if err != nil {
		log.Info(string(b))
		log.Error(err)
		return
	}
	fmt.Println("Successfully")
}

func addTable(table string) astutil.ApplyFunc {
	return func(cur *astutil.Cursor) bool {
		n := cur.Node()
		switch o := n.(type) {
		case *ast.FuncDecl:
			if o.Name.Name != "CreateTable" {
				return true
			}
			var filenames = make(map[string]bool)
			for _, stmt := range o.Body.List {
				ds, ok := stmt.(*ast.DeclStmt)
				if !ok {
					continue
				}
				gd := ds.Decl.(*ast.GenDecl)
				for _, spec := range gd.Specs {
					vspec, ok := spec.(*ast.ValueSpec)
					if !ok {
						continue
					}
					for _, vs := range vspec.Values {
						cl, ok := vs.(*ast.CompositeLit)
						if !ok {
							continue
						}
						for _, elt := range cl.Elts {
							bl, ok := elt.(*ast.BasicLit)
							if !ok {
								continue
							}
							filenames[bl.Value] = true
						}
						filenames[fmt.Sprintf(`"%v.sql"`, table)] = true
						cl.Elts = cl.Elts[:0]
						for filename := range filenames {
							cl.Elts = append(cl.Elts, &ast.BasicLit{
								Value: filename,
							})
						}
						return false
					}
				}
			}
		}
		return true
	}
}
