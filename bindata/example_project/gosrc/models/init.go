package models

import (
	"github.com/lixiangzhong/config"
	"github.com/lixiangzhong/log" 
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"net"
	"strings"
	"time"
	"{{.ProjectPath}}/gosrc/bindata"
)

const (
	StateOK  = 0
	StateDel = 1
)

var (
	db *sqlx.DB
)

func init() {
	initdb()
	{{if .HasUserTable}}
	initDefaultUser()
	{{end}}
}

func initdb() {
	mysqlconfig:=config.MySQLConfig("mysql")
	dbname:=mysqlconfig.DBName
	mysqlconfig.DBName=""
	mysql,err:=	sqlx.Connect("mysql",mysqlconfig.FormatDSN())
	if err != nil {
		log.Error(err)
		return 
	}
	defer mysql.Close()
	_,err=	mysql.Exec("CREATE DATABASE IF NOT EXISTS "+dbname+" default charset utf8mb4 COLLATE utf8mb4_general_ci")
	if err != nil {
		log.Error(err)
			return 
	}
	db, err = sqlx.Connect("mysql", config.MySQLConfig("mysql").FormatDSN())
	if err != nil {
		log.Error(err)
		return 
	}else{
		CreateTable()
	}
	if config.Bool("mysql.unsafe") {
		db = db.Unsafe()
	}
	go func() {
		tk := time.NewTicker(time.Second * 15)
		for range tk.C {
			db.Ping()
		}
	}()

}
 

func CreateTable()  {
	var tables=[]string{
		{{range $i,$v:=.Tables}}"{{$v.Name}}.sql",{{end}}
	}
 for _, table := range tables {
	 sql:=string(bindata.MustAsset(table))
	 if strings.TrimSpace(sql)!=""{
		 _,err:=db.Exec(sql)
		 if err != nil {
			 log.Error(err)
		 }
	 }
 }
}

{{if .HasUserTable}}
func initDefaultUser() {
	var defaultuser {{.UserTable.CamelCaseName}}
	defaultuser.{{.UserTable.UsernameColumn.CamelCaseName}} = "admin"
	defaultuser.{{.UserTable.PasswordColumn.CamelCaseName}} = "admin"
	defaultuser.Create()
}
{{end}}



func Tx(f func(*sqlx.Tx) error) error {
	tx, err := db.Beginx()
	if err != nil {
		return err
	}
	defer func() {
		if e := recover(); e != nil {
			log.Error(e)
			tx.Rollback()
		}
	}()
	err = f(tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}