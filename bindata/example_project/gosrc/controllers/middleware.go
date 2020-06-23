package controllers

import (
	"bytes"
	"io"
	"net/http"

	"{{.ProjectName}}/gosrc/app"
	"{{.ProjectName}}/gosrc/errcode"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
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


func (middleware) DumpBody(c *gin.Context) {
	reqbody := new(bytes.Buffer)
	dumpreq := false
	if c.Request.Body != nil && c.Request.Method != http.MethodGet {
		dumpreq = true
		c.Request.Body = struct {
			io.Reader
			io.Closer
		}{
			io.TeeReader(c.Request.Body, reqbody),
			c.Request.Body,
		}
	}
	rsprecord := &responseRecord{
		body:           new(bytes.Buffer),
		ResponseWriter: c.Writer,
	}
	c.Writer = rsprecord
	c.Next()
	if dumpreq {
		fmt.Println("请求", c.Request.URL)
		fmt.Println("reqbody", reqbody.String())
		fmt.Println("rspbody", rsprecord.body.String())
	}
}

type responseRecord struct {
	body *bytes.Buffer
	gin.ResponseWriter
}

func (rsp *responseRecord) Write(b []byte) (int, error) {
	rsp.body.Write(b)
	return rsp.ResponseWriter.Write(b)
}
