package slacker

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"path"
	"strings"
)

type TemplateData struct {
	ProjectName string
	ProjectPath string
	MysqlConfig *mysql.Config
	Tables      []Table
}

type Table struct {
	Name    string
	Columns []Column
}

//首字母
func (t Table) Initials() string {
	return string(t.Name[0])
}

func (t Table) CamelCaseName() string {
	return CamelCase(t.Name)
}

func (t Table) LowerName() string {
	return strings.ToLower(t.CamelCaseName())
}

func (t Table) PrimaryKeyColumn() Column {
	for _, col := range t.Columns {
		if col.IsPrimaryKey() {
			return col
		}
	}
	return Column{ColumnName: "id"}
}

func (t Table) SQLColumns() string {
	var fields []string
	for _, col := range t.Columns {
		fields = append(fields, fmt.Sprintf("%v", col.ColumnName))
	}
	return strings.Join(fields, ",")
}

func (t Table) StdSQL() string {
	var fields []string
	for _, col := range t.Columns {
		if col.ColumnName == "id" || col.IsPrimaryKey() {
			continue
		}
		fields = append(fields, fmt.Sprintf("%v=?", col.ColumnName))
	}
	return strings.Join(fields, ",")
}

func (t Table) NamedSQL() string {
	var fields []string
	for _, col := range t.Columns {
		if col.ColumnName == "id" || col.IsPrimaryKey() {
			continue
		}
		fields = append(fields, fmt.Sprintf("%v=:%v", col.ColumnName, col.ColumnName))
	}
	return strings.Join(fields, ",")
}

func (t Table) SwitchCase() string {
	var fields []string
	for _, col := range t.Columns {
		fields = append(fields, fmt.Sprintf(`"%v"`, col.ColumnName))
	}
	return strings.Join(fields, ",")
}

func (t Table) ImportLibrary(dir string) string {
	return path.Join(ProjectPath(), dir)
}

type Column struct {
	ColumnName    string `json:"column_name" db:"column_name"`
	DataType      string `json:"data_type" db:"data_type"`
	ColumnType    string `json:"column_type" db:"column_type"`
	ColumnComment string `json:"column_comment" db:"column_comment"`
	ColumnKey     string `json:"column_key" db:"column_key"`
}

func (c Column) CamelCaseName() string {
	return CamelCase(c.ColumnName)
}

func (c Column) Tag() string {
	return fmt.Sprintf("`%v`", fmt.Sprintf(`json:"%v" db:"%v"`, c.ColumnName, c.ColumnName))
}

func (c Column) Comment() string {
	if strings.TrimSpace(c.ColumnComment) == "" {
		return ""
	}
	return fmt.Sprintf("// %v", c.ColumnComment)
}

func (c Column) Type() string {
	c.ColumnName = strings.ToLower(c.ColumnName)
	switch {
	case strings.Contains(c.DataType, "int"):
		if strings.Contains(c.ColumnName, "time") || strings.Contains(c.ColumnName, "update") {
			return "int64"
		}
		if strings.Contains(c.ColumnName, "ip") {
			return "uint32"
		}
		if c.ColumnKey == "PRI" {
			return "int64"
		}
		return "int"
	case strings.Contains(c.DataType, "char"):
		return "string"
	case strings.Contains(c.DataType, "text"):
		return "string"
	case "timestamp" == c.DataType:
		return "time.Time"
	default:
		return "interface{}"
	}
}

func (c Column) IsPrimaryKey() bool {
	return c.ColumnKey == "PRI"
}
