package dao

import (
	"github.com/jmoiron/sqlx"
	"github.com/lixiangzhong/gorm"
	"github.com/lixiangzhong/log"
)

type Dao struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Dao {
	return &Dao{
		db: db,
	}
}

func (d *Dao) Init() {
	d.AutoMigrate()
}

func (d *Dao) gorm() *gorm.DB {
	db, err := gorm.Open("mysql", d.db.DB)
	if err != nil {
		log.Error(err)
	}
	return db
}

func (d *Dao) AutoMigrate() {
	d.gorm().Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		{{range $i,$table:=.Tables}}{{$table.LowerName}}.{{$table.CamelCaseName}}{},
		{{end}}
	)
}
