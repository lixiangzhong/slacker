package slacker

import (
	"bytes"
	"fmt"
	"go/format"
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/lixiangzhong/log"
	"github.com/lixiangzhong/slacker/mvctemplate"
	"golang.org/x/tools/imports"
)

var (
	FuncMap = template.FuncMap{
		"Contains": func(s, sub string) bool { return strings.Contains(strings.ToLower(s), sub) },
	}
)

func Tables(dbname string) (tables []Table) {
	tables = make([]Table, 0)
	var tablenames = make([]string, 0)
	err := db.Select(&tablenames, "show tables")
	if err != nil {
		log.Error(err)
		return
	}
	for _, tablename := range tablenames {
		var table Table
		table.Name = tablename
		table.Columns = make([]Column, 0)
		err = db.Select(&table.Columns, "select column_name,data_type,column_type,column_comment,column_key from information_schema.columns where table_schema =?  and table_name = ?", dbname, tablename)
		if err != nil {
			log.Error(err)
			continue
		}
		table.DBName = dbname
		table.ShowCreateTable()
		tables = append(tables, table)
	}
	return
}

func (t Table) ExecTemplate(mvc string, dir string) {
	var filename string
	var tpl = new(template.Template)
	var b []byte
	var formatter func([]byte) ([]byte, error)
	switch mvc {
	case "m", "models":
		if dir == MVCDefaultDir {
			dir = "gosrc/models"
		}
		filename = fmt.Sprintf("%v_%v.go", DBName, t.Name)
		b = mvctemplate.MustAsset("models.tpl")
		formatter = gofmt
	case "c", "controller":
		if dir == MVCDefaultDir {
			dir = "gosrc/controllers"
		}
		filename = fmt.Sprintf("%v_%v.go", DBName, t.Name)
		b = mvctemplate.MustAsset("controllers.tpl")
		formatter = gofmt
	case "v", "vue", "view":
		if dir == MVCDefaultDir {
			dir = "src/views"
		}
		filename = fmt.Sprintf("%v%v.vue", CamelCase(DBName), t.CamelCaseName())
		b = mvctemplate.MustAsset("vue.tpl")
		formatter = vuedelims
	case "js":
		if dir == MVCDefaultDir {
			dir = "src/api"
		}
		filename = fmt.Sprintf("%v_%v.js", DBName, t.Name)
		b = mvctemplate.MustAsset("js.tpl")
	case "sql":
		if dir == MVCDefaultDir {
			dir = "gosrc/bindata"
		}
		filename = fmt.Sprintf("%v_%v.sql", DBName, t.Name)
		b = mvctemplate.MustAsset("sql.tpl")
	}

	tpl.Funcs(FuncMap)
	_, err := tpl.Parse(string(b))
	if err != nil {
		log.Error(err)
		return
	}
	var writer = new(bytes.Buffer)

	err = tpl.Execute(writer, t)
	if err != nil {
		log.Error(err)
		return
	}
	b = writer.Bytes()
	if formatter != nil {
		b, err = formatter(b)
		if err != nil {
			log.Error(string(writer.Bytes()))
			log.Error(dir, filename, err)
			return
		}
	}
	f, err := os.OpenFile(path.Join(dir, filename), os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_EXCL, 0644)
	if err != nil {
		log.Error(err)
		return
	}
	defer f.Close()
	f.Write(b)
}

func gofmt(in []byte) ([]byte, error) {
	b, err := imports.Process("", in, nil)
	if err != nil {
		return nil, err
	}
	return format.Source(b)
}

func vuedelims(in []byte) ([]byte, error) {
	in = bytes.Replace(in, []byte("<<<"), []byte("{{"), -1)
	in = bytes.Replace(in, []byte(">>>"), []byte("}}"), -1)
	return in, nil
}
