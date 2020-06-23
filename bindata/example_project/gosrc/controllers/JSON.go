package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/lixiangzhong/config"
	"{{$.ProjectName}}/gosrc/log"
	"{{.ProjectName}}/gosrc/errcode"
)

type ResponseJSON struct {
	Code int         `json:"code"`
	Message  string      `json:"message"`
	Data interface{} `json:"data,omitempty"`
}

func JSON(c *gin.Context, data interface{}, err error) {
	var e errcode.CodeError

	switch err {
	case nil:
		e=errcode.OK
	case sql.ErrNoRows, gorm.ErrRecordNotFound:
		e=errcode.NotFound
		data=nil
	default:
		data=nil
		var ok bool
		e, ok = err.(errcode.CodeError)
		if !ok {
			if config.Bool("debug") {
				e = errcode.New(errcode.Error, err)
			} else {
				log.Debug(err)
				e = errcode.Error
			}
		}
	}
	c.JSON(http.StatusOK, ResponseJSON{
		Code: e.Code(),
		Message:  e.Message(),
		Data: data,
	})
}


