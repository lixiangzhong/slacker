//Generated by slacker
package dao

import (
	"github.com/jmoiron/sqlx"
	"github.com/jinzhu/gorm"
	"github.com/lixiangzhong/log" 
	{{range $i,$table:=.Tables}}{{$table.LowerName}} "{{$.ProjectPath}}/gosrc/models/{{$table.Name}}"
	{{end}}
)

type Dao struct {
	db *sqlx.DB
	gorm *gorm.DB
}

func New(db *sqlx.DB) *Dao {
	d:= &Dao{
		db: db,
	}
	d.initGorm()
	return d
}

func (d *Dao) Init() {
	d.autoMigrate()
}

func (d *Dao) initGorm()  {
	db, err := gorm.Open("mysql", d.db.DB)
	if err != nil {
		log.Error(err)
	}
	d.gorm = db.Set("gorm:save_associations", false)
}

func (d *Dao) autoMigrate() {
	d.gorm.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		{{range $i,$table:=.Tables}}{{$table.LowerName}}.{{$table.CamelCaseName}}{},
		{{end}}
	)
}


func (d *Dao) Create(data interface{})(error) { 
	return d.gorm.Create(&data).Error
}

func (d *Dao) Update(data interface{})error { 
	return d.gorm.Save(&data).Error
}

func (d *Dao) Patch(model interface{},u map[string]interface{})error { 
	return d.gorm.Model(&model).UpdateColumns(u).Error
}

func (d *Dao) Delete(data interface{})error { 
	return d.gorm.Delete(&data).Error
}

func (d *Dao) Take(data interface{})error { 
	return d.gorm.Take(&data).Error
}

func (d *Dao) List(data interface{})error { 
	return d.gorm.Find(&data).Error
}


func OffsetLimitScope(offset, limit uint64) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if offset > 0 {
			db = db.Offset(offset)
		}
		if limit > 0 {
			db = db.Limit(limit)
		}
		return db
	}
}
