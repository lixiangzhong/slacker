//Generated by slacker
package dao

import (
//sq	"github.com/Masterminds/squirrel"
"{{"gosrc/models"| .ImportLibrary}}/{{.Name}}"
)

func (d *Dao) List{{.CamelCaseName}}(offset,limit uint64,search {{.LowerName}}.{{.CamelCaseName}}) ([]{{.LowerName}}.{{.CamelCaseName}},int,error) {
    var data =make([]{{.LowerName}}.{{.CamelCaseName}},0)
    var total int
	s:=d.gorm.Model({{.LowerName}}.{{.CamelCaseName}}{})
	{{if Contains .SwitchCase "state"}} 
		s=s.Where("state!=?",{{.LowerName}}.StateDel)
	{{end}}
	err := s.Count(&total).Order("{{.PrimaryKeyColumn.ColumnName}} desc").Offset(offset).Limit(limit).Find(&data).Error
    return data,total,err
}

{{.MethodTake}}

func (d *Dao) Create{{.CamelCaseName}}({{.LowerName}} {{.LowerName}}.{{.CamelCaseName}})({{.LowerName}}.{{.CamelCaseName}}, error ){
 	{{.LowerName | .AutomaticCreateUpdateExpression}}

	err:=d.gorm.Create(&{{.LowerName}}).Error
    return {{.LowerName}}, err
}

func (d *Dao) Update{{.CamelCaseName}}({{.LowerName}} {{.LowerName}}.{{.CamelCaseName}}) error {
	 {{.LowerName | .AutomaticUpdateExpression}}
   // _,err := d.db.NamedExec("update {{.Name}} set {{.NamedSQL}} where {{.PrimaryKeyColumn.ColumnName}}=:{{.PrimaryKeyColumn.ColumnName}}",{{.LowerName}})
   err:=d.gorm.Save({{.LowerName}}).Error
    return err
}

func (d *Dao) Patch{{.CamelCaseName}}(id int64,update map[string]interface{}) error {
  	{{.AutomaticUpdateMapExpression}}
  	var data {{.LowerName}}.{{.CamelCaseName}}
  	data.{{.PrimaryKeyColumn.CamelCaseName}}=id
	err:=d.gorm.Model(data).Updates(update).Error
	return err
}


{{.MethodDelete}}



{{if .IsUserTable}} 

func (d *Dao) Take{{.CamelCaseName}}ByName(username string) ({{.LowerName}}.{{.CamelCaseName}}, error) {
    var {{.LowerName}} {{.LowerName}}.{{.CamelCaseName}}
	username = strings.ToLower(username)
	SQL := "select * from {{.Name}} where {{.UsernameColumn.ColumnName }}=? limit 1"
	err := d.db.Get(&{{.LowerName}}, SQL, username)
	return {{.LowerName}}, err
}

func (d *Dao) {{.CamelCaseName}}IsExist(username string) bool {
	SQL := "select 1 from {{.Name}} where {{.UsernameColumn.ColumnName}}=? limit 1"
	var exist bool
	err := d.db.QueryRow(SQL,username).Scan(&exist)
	if err != nil && err != sql.ErrNoRows {
		return true
	}
	return exist
}

{{end}}