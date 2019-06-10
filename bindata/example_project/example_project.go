package main

import (
	"{{.ProjectPath}}/gosrc/app"
	"{{.ProjectPath}}/gosrc/controllers"
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
func main() {
	app.Init()
	controllers.Init()
	initroute()
	app.Run()
}
