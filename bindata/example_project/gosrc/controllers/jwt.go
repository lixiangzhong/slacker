package controllers

import (
	"{{.ProjectPath}}/gosrc/app" 
	"{{.ProjectPath}}/gosrc/errcode" 
	"github.com/lixiangzhong/config"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/lixiangzhong/base64Captcha"
	"net/http"
	"strconv"
	"time"
)

type LoginForm struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	CaptchaID    string `json:"captcha_id"`
	CaptchaValue string `json:"captcha_value"`
}

func GetToken(c *gin.Context) {
	var t LoginForm
	if err := c.ShouldBindJSON(&t); err != nil {
		JSON(c,nil,errcode.BadRequest) 
		return
	}
	if !base64Captcha.VerifyCaptcha(t.CaptchaID, t.CaptchaValue) {
		JSON(c,nil,errcode.InvalidCaptcha)  
		return
	}

	{{if .HasUserTable}}
	user, err := Service.Take{{.UserTable.CamelCaseName}}ByName(t.Username)
	if err != nil {
		JSON(c,nil,errcode.IncorrectUserOrPwd)
		return
	}
	if !Service.ValidPassword(t.Password,user.{{.UserTable.PasswordColumn.CamelCaseName}}) {
		JSON(c,nil,errcode.IncorrectUserOrPwd)
		return
	}  
	token, err := GenerateToken(user.ID, app.JWTExpireIn)
	if err != nil {
		JSON(c,nil,err)
		return
	}
	JSON(c,gin.H{
		"token":    token,
		"expirein": app.JWTExpireIn,
	},nil)
	{{else}}
	username := config.String("auth.user")
	password := config.String("auth.password")
	if username != t.Username {
		JSON(c,nil,errcode.IncorrectUserOrPwd)
		return
	}
	if password != t.Password {
		JSON(c,nil,errcode.IncorrectUserOrPwd)
		return
	} 
	token, err := GenerateToken(0, app.JWTExpireIn)
		if err != nil { 
			JSON(c,nil,err)
			return
		}
		JSON(c,gin.H{
			"token":    token,
			"expirein": app.JWTExpireIn,
		},nil) 
	{{end}}
}

func GenerateToken(uid int64, expirein int64) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = jwt.MapClaims{
		"uid": uid,
		"exp": time.Now().Unix() + expirein,
	}
	s, err := token.SignedString([]byte(app.JWTSecret))
	if err != nil {
		return s, fmt.Errorf("%v", "生成token出错")
	}
	return s, nil
}

func UID(c *gin.Context) int64 {
	uid, err := strconv.ParseInt(fmt.Sprintf("%v", c.MustGet("uid")), 10, 64)
	if err != nil {
		panic(err)
	}
	return uid
}
