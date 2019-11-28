package main

import (
	"{{.ProjectPath}}/gosrc/app"
	"{{.ProjectPath}}/gosrc/controllers"
	"github.com/lixiangzhong/log"
	"github.com/lixiangzhong/config"
	"github.com/jmoiron/sqlx"
)

func initroute() {
	app.Engine.GET("/captcha", controllers.Captcha)
	app.Engine.POST("/token", controllers.GetToken)
	app.Engine.GET("/meta", controllers.Meta)

	api := app.Engine.Group("/api", controllers.MiddleWare.JWTAuth)
	// importdata:=api.Group("/import")
	// exportdata:=api.Group("/export")
	//websocket:=api.Group("/websocket")

	{{range $i,$table:=.Tables}}
	//{{$table.Name}}
	api.GET("/{{$table.LowerName}}/:id",controllers.Take{{$table.CamelCaseName}})
	api.GET("/{{$table.LowerName}}",controllers.List{{$table.CamelCaseName}})
	api.POST("/{{$table.LowerName}}",controllers.Create{{$table.CamelCaseName}})
	api.PUT("/{{$table.LowerName}}/:id",controllers.Update{{$table.CamelCaseName}})
	api.PATCH("/{{$table.LowerName}}/:id",controllers.Patch{{$table.CamelCaseName}})
	api.DELETE("/{{$table.LowerName}}/:id",controllers.Delete{{$table.CamelCaseName}})
	{{end}}

}
func initDB() *sqlx.DB {
	var dbname string
	var cfg = config.MySQLConfig("mysql")
	cfg.DBName, dbname = dbname, cfg.DBName
	mysql := sqlx.MustConnect("mysql", cfg.FormatDSN())
	_, err := mysql.Exec("CREATE DATABASE IF NOT EXISTS `" + dbname + "` default charset utf8mb4 COLLATE utf8mb4_general_ci")
	if err != nil {
		log.Error(err)
	}
	defer mysql.Close()
	db := sqlx.MustConnect("mysql", config.MySQLConfig("mysql").FormatDSN())
	return db
}

func init(){
	app.Init()
	db:=initDB()
	controllers.InitService(db)
	initroute()
}

func main() {

	app.Run()
}
