package main

import (
	"{{.ProjectPath}}/gosrc/app"
	"{{.ProjectPath}}/gosrc/controllers"
)

func initroute() {
	app.Engine.GET("/captcha", controllers.Captcha{}.Get)
	app.Engine.POST("/token", controllers.GetToken)
	app.Engine.GET("/meta", controllers.Meta)

	api := app.Engine.Group("/api", controllers.MiddleWare.JWTAuth)
	importdata:=api.Group("/import")
	exportdata:=api.Group("/export")

	{{range $i,$table:=.Tables}}
	//{{$table.Name}}
	api.GET("/{{$table.LowerName}}/:id",controllers.{{$table.CamelCaseName}}{}.Take)
	api.GET("/{{$table.LowerName}}",controllers.{{$table.CamelCaseName}}{}.List)
	api.POST("/{{$table.LowerName}}",controllers.{{$table.CamelCaseName}}{}.Create)
	api.PUT("/{{$table.LowerName}}/:id",controllers.{{$table.CamelCaseName}}{}.Update)
	api.PATCH("/{{$table.LowerName}}/:id",controllers.{{$table.CamelCaseName}}{}.Patch)
	api.DELETE("/{{$table.LowerName}}/:id",controllers.{{$table.CamelCaseName}}{}.Delete) 
	api.PATCH("/{{$table.LowerName}}",controllers.{{$table.CamelCaseName}}{}.BatchPatch)
	api.DELETE("/{{$table.LowerName}}",controllers.{{$table.CamelCaseName}}{}.BatchDelete)
	api.PUT("/{{$table.LowerName}}",controllers.{{$table.CamelCaseName}}{}.BatchUpdate)
	importdata.POST("/{{$table.LowerName}}", controllers.{{$table.CamelCaseName}}{}.Import)
	exportdata.GET("/{{$table.LowerName}}", controllers.{{$table.CamelCaseName}}{}.Export)
	{{end}}

}
func main() {
	initroute()
	app.Run()
}
