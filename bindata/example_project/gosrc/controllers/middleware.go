package controllers

import (
	"{{.ProjectName}}/gosrc/app"
	"{{.ProjectName}}/gosrc/errcode"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
	"net/http"
)

var MiddleWare middleware

type middleware struct {
}

//检验token
func (m middleware) JWTAuth(c *gin.Context) {
	token, err := request.ParseFromRequest(c.Request, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
		return []byte(app.JWTSecret), nil
	})
	if err != nil {
		//token 过期或不正确 
		c.Abort()
		JSON(c, nil,errcode.InvalidToken  )
		return
	}
	if m, ok := token.Claims.(jwt.MapClaims); ok {
		uid, ok := m["uid"]
		if !ok {
			c.Abort()
			JSON(c, nil,errcode.InvalidToken )
			return
		}
		c.Set("uid", uid)
	}
}
