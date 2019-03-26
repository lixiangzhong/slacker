package slacker

import (
	"fmt"
	"path"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/lixiangzhong/log"
)

type TemplateData struct {
	ProjectName string
	ProjectPath string
	MysqlConfig *mysql.Config
	DBName      string
	Tables      []Table
}

func (t TemplateData) HasUserTable() bool {
	for _, table := range t.Tables {
		if table.IsUserTable() {
			return true
		}
	}
	return false
}

func (t TemplateData) UserTable() Table {
	for _, table := range t.Tables {
		if table.IsUserTable() {
			return table
		}
	}
	return Table{Name: "user"}
}

type Table struct {
	Name           string
	CreateTableSQL string
	Columns        []Column
}

//首字母
func (t Table) Initials() string {
	return strings.ToLower(string(t.Name[0]))
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

func (t Table) StateColumn() Column {
	for _, col := range t.Columns {
		if StringInSlice(col.ColumnName, StateFields) {
			return col
		}
	}
	return Column{ColumnName: "state"}
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
		fields = append(fields, fmt.Sprintf("%v=:%v", Quote(col.ColumnName), col.ColumnName))
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

func (t *Table) ShowCreateTable() {
	err := db.QueryRow("show create table "+t.Name).Scan(&t.Name, &t.CreateTableSQL)
	if err != nil {
		log.Error(err)
		return
	}
	t.CreateTableSQL = strings.Replace(t.CreateTableSQL, "CREATE TABLE", "CREATE TABLE IF NOT EXISTS", 1)
	fields := strings.Fields(t.CreateTableSQL)
	for i, v := range fields {
		if strings.Contains(v, "AUTO_INCREMENT=") {
			fields[i] = ""
			break
		}
	}
	t.CreateTableSQL = strings.Join(fields, " ")
}

func (t Table) AutomaticCreateUpdateExpression(obj string) string {
	var exp string
	fields := append(AutoAssignCreateFields, AutoAssignUpdateFields...)
	for _, col := range t.Columns {
		if StringInSlice(col.ColumnName, fields) {
			switch {
			case col.Type() == "time.Time":
				exp += fmt.Sprintf("%v.%v = %v\n", obj, col.CamelCaseName(), "timeNow")
			case strings.Contains(col.Type(), "int"):
				exp += fmt.Sprintf("%v.%v = %v\n", obj, col.CamelCaseName(), "timeNow.Unix()")
			}
		}
	}
	if exp != "" {
		exp = "var timeNow = time.Now()\n" + exp
	}
	return exp
}

func (t Table) AutomaticUpdateExpression(obj string) string {
	var exp string
	for _, col := range t.Columns {
		if StringInSlice(col.ColumnName, AutoAssignUpdateFields) {
			switch {
			case col.Type() == "time.Time":
				exp += fmt.Sprintf("%v.%v = %v\n", obj, col.CamelCaseName(), "timeNow")
			case strings.Contains(col.Type(), "int"):
				exp += fmt.Sprintf("%v.%v = %v\n", obj, col.CamelCaseName(), "timeNow.Unix()")
			}
		}
	}
	if exp != "" {
		exp = "var timeNow = time.Now()\n" + exp
	}
	return exp
}

func (t Table) AutomaticUpdateMapExpression() string {
	var exp string
	for _, col := range t.Columns {
		if StringInSlice(col.ColumnName, AutoAssignUpdateFields) {
			switch {
			case col.Type() == "time.Time":
				exp += fmt.Sprintf(`update["%v"] = %v`, col.ColumnName, "timeNow\n")
			case strings.Contains(col.Type(), "int"):
				exp += fmt.Sprintf(`update["%v"] = %v`, col.ColumnName, "timeNow.Unix()\n")
			}
		}
	}
	if exp != "" {
		exp = "var timeNow = time.Now()\n" + exp
	}
	return exp
}

func (t Table) HasStateColumn() bool {
	for _, col := range t.Columns {
		if StringInSlice(col.ColumnName, StateFields) {
			return true
		}
	}
	return false
}

func (t Table) MethodTake() string {
	var s = fmt.Sprintf("func (%v %v) Take() (%v,error) {\n",
		t.Initials(), t.CamelCaseName(), t.CamelCaseName(),
	)
	if t.HasStateColumn() {
		s += fmt.Sprintf(`err:=db.Get(&%v,"select * from %v where %v=? and %v!=? limit 1",%v.%v,StateDel)`,
			t.Initials(), t.Name, t.PrimaryKeyColumn().ColumnName, t.StateColumn().ColumnName, t.Initials(), t.PrimaryKeyColumn().CamelCaseName(),
		)
	} else {
		s += fmt.Sprintf(`err:=db.Get(&%v,"select * from %v where %v=? limit 1",%v.%v)`,
			t.Initials(), t.Name, t.PrimaryKeyColumn().ColumnName, t.Initials(), t.PrimaryKeyColumn().CamelCaseName(),
		)
	}
	s += fmt.Sprintf("\n return %v,err", t.Initials())
	s += "}"
	return s
}

func (t Table) MethodDelete() string {
	var s = fmt.Sprintf("func (%v %v) Delete() error {\n",
		t.Initials(), t.CamelCaseName(),
	)
	var updatedfields []string
	for _, col := range t.Columns {
		if StringInSlice(col.ColumnName, AutoAssignUpdateFields) {
			updatedfields = append(updatedfields, col.ColumnName)
		}
	}

	switch {
	case t.HasStateColumn() && updatedfields == nil,
		!t.HasStateColumn():
		s += fmt.Sprintf(` _,err := db.Exec("delete from %v where %v=?",%v.%v)`,
			t.Name, t.PrimaryKeyColumn().ColumnName, t.Initials(), t.PrimaryKeyColumn().CamelCaseName(),
		)
	case t.HasStateColumn() && updatedfields != nil:
		s += "var timeNow = time.Now()\n"
		var namedfields = make([]string, 0)
		for _, field := range updatedfields {
			namedfields = append(namedfields, fmt.Sprintf("`%v`=:%v", field, field))

			for _, col := range t.Columns {
				if col.ColumnName == field {
					switch col.Type() {
					case "time.Time":
						s += fmt.Sprintf("%v.%v=timeNow\n", t.Initials(), CamelCase(field))
					case "int64":
						s += fmt.Sprintf("%v.%v=timeNow.Unix()\n", t.Initials(), CamelCase(field))
					}
				}
			}
		}
		s += fmt.Sprintf("%v.%v=StateDel\n", t.Initials(), t.StateColumn().CamelCaseName())
		s += fmt.Sprintf(` _,err := db.NamedExec("update %v set %v=:%v,`+strings.Join(namedfields, ",")+` where %v=:%v",%v)`,
			t.Name,
			Quote(t.StateColumn().ColumnName), t.StateColumn().ColumnName,
			Quote(t.PrimaryKeyColumn().ColumnName), t.PrimaryKeyColumn().ColumnName,
			t.Initials(),
		)
	}

	s += fmt.Sprintf("\n return err")
	s += "}"
	return s
}

func (t Table) IsUserTable() bool {
	var hasusername bool
	var haspassword bool
	for _, c := range t.Columns {
		if StringInSlice(c.ColumnName, UsernameFields) {
			hasusername = true
			continue
		}
		if StringInSlice(c.ColumnName, PasswordFields) {
			haspassword = true
			continue
		}
	}
	return hasusername && haspassword
}

func (t Table) UsernameColumn() Column {
	for _, c := range t.Columns {
		if StringInSlice(c.ColumnName, UsernameFields) {
			return c
		}
	}
	return Column{ColumnName: "username"}
}

func (t Table) PasswordColumn() Column {
	for _, c := range t.Columns {
		if StringInSlice(c.ColumnName, PasswordFields) {
			return c
		}
	}
	return Column{ColumnName: "password"}
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
	return fmt.Sprintf("`%v`", fmt.Sprintf(`json:"%v" db:"%v" form:"%v"`, c.ColumnName, c.ColumnName, c.ColumnName))
}

func (c Column) Comment() string {
	if strings.TrimSpace(c.ColumnComment) == "" {
		return ""
	}
	return fmt.Sprintf("// %v", c.ColumnComment)
}

func (c Column) Type() string {
	c.ColumnName = strings.ToLower(c.ColumnName)
	if c.IsPrimaryKey() {
		return "int64"
	}
	switch {
	case strings.Contains(c.DataType, "int"):
		if StringInSlice(c.ColumnName, AutoAssignFields) {
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
	case strings.Contains(c.DataType, "float"):
		return "float64"
	default:
		return "interface{}"
	}
}

func (c Column) IsPrimaryKey() bool {
	return c.ColumnKey == "PRI"
}

var (
	AutoAssignFields []string

	AutoAssignCreateFields = []string{
		"ctime", "Ctime",
		"created", "Created",
		"create_time", "created_time",
		"created_at", "create_at", "createdAt", "createAt",
		"addtime", "AddTime",
	}

	AutoAssignUpdateFields = []string{
		"utime", "Utime",
		"updated", "Updated",
		"update_time", "updated_time",
		"updated_at", "update_at", "updatedAt", "updateAt",
	}

	AutoAssignDeleteFields = []string{
		"dtime", "Dtime",
		"deleted", "Deleted",
		"delete_time", "deleted_time",
		"deletetime", "DeleteTime",
		"deleted_at", "delete_at", "deletedAt", "deleteAt",
	}

	StateFields = []string{
		"state", "status",
	}

	UsernameFields = []string{
		"user", "username",
	}
	PasswordFields = []string{
		"passwd", "password",
	}
)

func init() {
	AutoAssignFields = append(AutoAssignFields, AutoAssignCreateFields...)
	AutoAssignFields = append(AutoAssignFields, AutoAssignUpdateFields...)
	AutoAssignFields = append(AutoAssignFields, AutoAssignDeleteFields...)
}
