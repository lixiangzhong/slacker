package controllers

import (
	"github.com/lixiangzhong/config"
	"{{.ProjectPath}}/gosrc/app"
	"github.com/gin-gonic/gin"
)

func Meta(c *gin.Context) {
	logo := config.String("meta.logo")
	title := config.String("meta.title")
	c.JSON(200, JSON.OK(gin.H{
		"logo":       logo,
		"title":      title,
		"version":    app.Version,
		"gitversion": app.GitVersion,
	}))
}
