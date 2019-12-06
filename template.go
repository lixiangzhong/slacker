package slacker

import (
	"database/sql"
	"fmt"
	"path"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/lixiangzhong/log"
	"github.com/lixiangzhong/name"
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
	ProjectName    string
	Name           string
	CreateTableSQL string
	DBName         string
	Columns        []Column
}

//首字母
func (t Table) Initials() string {
	return strings.ToLower(string(t.Name[0]))
}

func (t Table) CamelCaseName() string {
	return name.CamelCase(t.Name)
}

func (t Table) CamelCaseNameWithDBName() string {
	return name.CamelCase(t.DBName) + name.CamelCase(t.Name)
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
	return path.Join(t.ProjectName, dir)
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
	var s = fmt.Sprintf("func (d *Dao) Take%v(data *%v.%v) error {\n",
		t.CamelCaseName(), t.LowerName(), t.CamelCaseName(),
	)
	if t.HasStateColumn() {
		s += fmt.Sprintf(`err:=d.gorm.Take(data,"%v=? and %v!=?",data.%v,%v.StateDel).Error`,
			t.PrimaryKeyColumn().ColumnName, t.StateColumn().ColumnName, t.PrimaryKeyColumn().CamelCaseName(), t.LowerName(),
		)
	} else {
		s += fmt.Sprintf("err:=d.gorm.Take(data,data.%v).Error", t.PrimaryKeyColumn().CamelCaseName())
	}
	s += fmt.Sprintf("\n return err")
	s += "}"
	return s
}

func (t Table) MethodDelete() string {
	s := fmt.Sprintf("func (d *Dao) Delete%v(id int64) error {\n", t.CamelCaseName())
	s += fmt.Sprintf("var data %v.%v\n", t.LowerName(), t.CamelCaseName())
	s += fmt.Sprintf("data.%v=id\n", t.PrimaryKeyColumn().CamelCaseName())
	var updatedfields []string
	for _, col := range t.Columns {
		if StringInSlice(col.ColumnName, AutoAssignUpdateFields) {
			updatedfields = append(updatedfields, col.ColumnName)
		}
	}

	switch {
	case t.HasStateColumn() && updatedfields == nil,
		!t.HasStateColumn():
		s += "err:=d.gorm.Delete(&data).Error\n"
	case t.HasStateColumn() && updatedfields != nil:
		s += "var timeNow = time.Now()\n"
		s += "var fields =make(map[string]interface{})\n"
		for _, field := range updatedfields {
			for _, col := range t.Columns {
				if col.ColumnName == field {
					switch col.Type() {
					case "time.Time":
						// s += fmt.Sprintf("data.%v=timeNow\n",   name.CamelCase(field))
						s += fmt.Sprintf(`fields["%v"]=timeNow`, field) + "\n"
					case "int64":
						// s += fmt.Sprintf("data.%v=timeNow.Unix()\n",   name.CamelCase(field))
						s += fmt.Sprintf(`fields["%v"]=timeNow.Unix()`, field) + "\n"
					}
				}
			}
		}
		// s += fmt.Sprintf("data.%v=%v.StateDel\n",   t.StateColumn().CamelCaseName(), t.LowerName())
		s += fmt.Sprintf(`fields["%v"]=%v.StateDel`, t.StateColumn().ColumnName, t.LowerName()) + "\n"
		s += fmt.Sprintf("err:=d.gorm.Model(&data).UpdateColumns(fields).Error\n")
	}
	s += fmt.Sprintf("return err\n}")

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
	ColumnName    string         `json:"column_name" db:"column_name"`
	DataType      string         `json:"data_type" db:"data_type"`
	ColumnType    string         `json:"column_type" db:"column_type"`
	ColumnComment string         `json:"column_comment" db:"column_comment"`
	ColumnKey     string         `json:"column_key" db:"column_key"`
	ColumnDefault sql.NullString `json:"column_default" db:"column_default"`
	// ReferencedDB     sql.NullString `json:"referenced_db" db:"referenced_table_schema"`
	// ReferencedTable  sql.NullString `json:"referenced_table" db:"referenced_table_name"`
	// ReferencedColumn sql.NullString `json:"referenced_column" db:"referenced_column_name"`
	// OnDelete         string         `json:"on_delete" db:"delete_rule"`
	// OnUpdate         string         `json:"on_update" db:"update_rule"`
}

func (c Column) CamelCaseName() string {
	return name.CamelCase(c.ColumnName)
}

func (c Column) Tag() string {
	var index string
	if c.IsPrimaryKey() {
		index = "PRIMARY_KEY"
		c.ColumnType += " AUTO_INCREMENT"
	} else {
		switch c.ColumnKey {
		case "UNI":
			index = "UNIQUE"
		case "MUL":
			index = "index:idx_" + c.ColumnName
		}
	}
	var defaultvalue string
	if c.ColumnDefault.Valid {
		defaultvalue = "DEFAULT:" + c.ColumnDefault.String
	}
	return fmt.Sprintf("`%v`", fmt.Sprintf(`json:"%v" db:"%v" form:"%v" gorm:"column:%v;type:%v;not null%v%v"`,
		c.ColumnName, c.ColumnName, c.ColumnName,
		//gorm
		c.ColumnName, //column
		c.ColumnType, //type
		addSemicolonPrefixIfExist(index),
		addSemicolonPrefixIfExist(defaultvalue),
	))
}

func (c Column) Comment() string {
	if strings.TrimSpace(c.ColumnComment) == "" {
		return ""
	}
	return fmt.Sprintf("// %v", c.ColumnComment)
}

func (c Column) Type() string {
	name := strings.ToLower(c.ColumnName)
	t := c.ColumnType
	switch {
	case strings.Contains(t, "int"):
		if c.IsPrimaryKey() {
			return "int64"
		}
		return guessDataType(name)
	case strings.Contains(t, "char"), strings.Contains(t, "text"):
		return "string"
	case "timestamp" == t:
		return "time.Time"
	case strings.Contains(t, "float"):
		return "float64"
	}
	return "interface{}"
}

func (c Column) ZeroValue() string {
	if strings.Contains(c.Type(), "int") {
		return "0"
	}
	return `""`
}

func guessDataType(name string) string {
	switch {
	case likeTimeUnix(name):
		return "int64"
	case strings.Contains(name, "id"):
		return "int64"
	case strings.Contains(name, "ip"):
		return "uint32"
	case strings.Contains(name, "net"):
		return "uint32"
	}
	return "int"
}

func likeTimeUnix(s string) bool {
	switch {
	case strings.Contains(s, "time"):
	case strings.Contains(s, "create"):
	case strings.Contains(s, "update"):
	default:
		return false
	}
	return true
}

func (c Column) IsPrimaryKey() bool {
	return c.ColumnKey == "PRI"
}

func addSemicolonPrefixIfExist(s string) string {
	if s == "" {
		return s
	}
	return ";" + s
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
