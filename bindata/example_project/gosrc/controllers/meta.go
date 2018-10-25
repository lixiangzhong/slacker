package controllers

import (
	"dns.com/ini"
	"{{.ProjectPath}}/gosrc/app"
	"github.com/gin-gonic/gin"
)

func Meta(c *gin.Context) {
	logo := ini.DefaultString("meta", "logo", "{{.ProjectName}}")
	title := ini.DefaultString("meta", "title", "{{.ProjectName}}")
	c.JSON(200, JSON.OK(gin.H{
		"logo":       logo,
		"title":      title,
		"version":    app.Version,
		"gitversion": app.GitVersion,
	}))
}
