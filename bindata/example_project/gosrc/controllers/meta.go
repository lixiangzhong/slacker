package controllers

import (
	"github.com/lixiangzhong/config"
	"{{.ProjectName}}/gosrc/app"
	"github.com/gin-gonic/gin"
)

func Meta(c *gin.Context) {
	logo := config.String("meta.logo")
	title := config.String("meta.title") 
	JSON(c, gin.H{
		"logo":       logo,
		"title":      title,
		"version":    app.Version,
		"gitversion": app.GitVersion,
	}, nil)
}
