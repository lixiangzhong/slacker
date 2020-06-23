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
		"Contains":    func(s, sub string) bool { return strings.Contains(strings.ToLower(s), sub) },
		"NotContains": func(s, sub string) bool { return !strings.Contains(strings.ToLower(s), sub) },
	}
)

func Tables(projectname, dbname string) (tables []Table) {
	tables = make([]Table, 0)
	var tablenames = make([]string, 0)
	err := db.Select(&tablenames, "show tables")
	if err != nil {
		log.Error(err)
		return
	}
	for _, tablename := range tablenames {
		var table Table
		table.ProjectName = projectname
		table.Name = tablename
		table.Columns = make([]Column, 0)
		err = db.Select(&table.Columns, "select column_name,data_type,column_type,column_comment,column_key,column_default from information_schema.columns where table_schema =?  and table_name = ?", dbname, tablename)
		if err != nil {
			log.Error(err)
			continue
		}
		// for i := 0; i < len(table.Columns); i++ {
		// 	col := table.Columns[i]
		// 	err = db.Get(&col, "select referenced_table_schema,referenced_table_name,referenced_column_name from information_schema.key_column_usage where table_schema=? and table_name=? and column_name=?", dbname, tablename, col.ColumnName)
		// 	if err == nil {
		// 		db.Get(&col, "select delete_rule,update_rule from information_schema.referential_constraints where constraint_schema=? and table_name=? and referenced_table_name=?", dbname, tablename, col.ReferencedTable.String)
		// 	}
		// }
		table.DBName = dbname
		//table.ShowCreateTable()
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
			dir = "gosrc/models/" + t.Name
		}
		filename = fmt.Sprintf("%v.go", t.Name)
		b = mvctemplate.MustAsset("models.tpl")
		formatter = goimports
	case "c", "controller":
		if dir == MVCDefaultDir {
			dir = "gosrc/controllers"
		}
		filename = fmt.Sprintf("%v.go", t.Name)
		b = mvctemplate.MustAsset("controllers.tpl")
		formatter = goimports
	case "v", "vue", "view":
		if dir == MVCDefaultDir {
			dir = "src/views"
		}
		filename = fmt.Sprintf("%v.vue", t.CamelCaseName())
		b = mvctemplate.MustAsset("vue.tpl")
		formatter = vuedelims
	case "comp", "component":
		if dir == MVCDefaultDir {
			dir = "src/components"
		}
		filename = fmt.Sprintf("%v-select.vue", t.Name)
		b = mvctemplate.MustAsset("component.tpl")
		formatter = vuedelims
	case "js":
		if dir == MVCDefaultDir {
			dir = "src/api"
		}
		filename = fmt.Sprintf("%v.js", t.Name)
		b = mvctemplate.MustAsset("js.tpl")
	case "md":
		if dir == MVCDefaultDir {
			dir = "markdowndoc"
		}
		filename = fmt.Sprintf("%v.md", t.Name)
		b = mvctemplate.MustAsset("markdown.tpl")
	case "dao":
		if dir == MVCDefaultDir {
			dir = "gosrc/dao"
		}
		filename = fmt.Sprintf("%v.go", t.Name)
		b = mvctemplate.MustAsset("dao.tpl")
		formatter = goimports
	case "service":
		if dir == MVCDefaultDir {
			dir = "gosrc/service"
		}
		filename = fmt.Sprintf("%v.go", t.Name)
		b = mvctemplate.MustAsset("service.tpl")
		formatter = goimports
	}

	tpl.Funcs(FuncMap)
	_, err := tpl.Parse(string(b))
	if err != nil {
		log.Error(mvc, t.Name, err)
		return
	}
	var writer = new(bytes.Buffer)

	err = tpl.Execute(writer, t)
	if err != nil {
		log.Error(mvc, t.Name, err)
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
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		log.Error(mvc, dir, err)
		return
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
	return format.Source(in)
}

func goimports(in []byte) ([]byte, error) {
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
