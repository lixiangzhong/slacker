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

}
func main() {
	initroute()
	app.Run()
}
