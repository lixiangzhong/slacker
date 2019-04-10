package controllers

import (
	"fmt"

	"github.com/lixiangzhong/log"
)

var JSON ResponseJSON

type ResponseJSON struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func (r ResponseJSON) New(code int, msg string, data interface{}) ResponseJSON {
	return ResponseJSON{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

func (r ResponseJSON) OK(data interface{}) ResponseJSON {
	return r.New(0, "ok", data)
}

func (j ResponseJSON) Error(err interface{}) ResponseJSON {
	return j.New(1, fmt.Sprintf("%v", err), nil)
}

func (j ResponseJSON) Permission() ResponseJSON {
	return j.New(2, "没有操作权限", nil)
}

func (r ResponseJSON) BadRequest() ResponseJSON {
	return r.New(400, "参数错误", nil)
}

func (r ResponseJSON) BadBinding(err interface{}) ResponseJSON {
	return r.New(400, fmt.Sprintf("参数错误:%v", err), nil)
}

func (j ResponseJSON) IncorrectUserOrPwd() ResponseJSON {
	return j.New(402, "账号或密码不正确", nil)
}

func (j ResponseJSON) InValidCaptcha() ResponseJSON {
	return j.New(402, "验证码错误", nil)
}

func (j ResponseJSON) InvalidToken() ResponseJSON {
	return j.New(403, "token无效", nil)
}

func (j ResponseJSON) TimeoutToken() ResponseJSON {
	return j.New(403, "登录信息过期,请重新登录", nil)
}

func (j ResponseJSON) MysqlError(err error) ResponseJSON {
	log.Debug(err)
	return j.New(500, "查询数据库出错", nil)
}
