package controllers

import (
	"{{.ProjectPath}}/gosrc/app"
	{{if .HasUserTable}}
	"{{.ProjectPath}}/gosrc/models"
	{{end}}
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
	CaptchaID    string `json:"captcha_id" db:"captcha_id"`
	CaptchaValue string `json:"captcha_value" db:"captcha_value"`
}

func GetToken(c *gin.Context) {
	var t LoginForm
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusOK, JSON.BadRequest())
		return
	}
	if !base64Captcha.VerifyCaptcha(t.CaptchaID, t.CaptchaValue) {
		c.JSON(http.StatusOK, JSON.InValidCaptcha())
		return
	}

	{{if .HasUserTable}}
	user, err := models.{{.UserTable.CamelCaseName}}{}.TakeByName(t.Username)
	if err != nil {
		c.JSON(http.StatusOK, JSON.IncorrectUserOrPwd())
		return
	}
	if user.ValidPassword(t.Password) {
		token, err := GenerateToken(user.ID, app.JWTExpireIn)
		if err != nil {
			c.JSON(http.StatusOK, JSON.Error(err))
			return
		}
		c.JSON(http.StatusOK, JSON.OK(gin.H{
			"token":    token,
			"expirein": app.JWTExpireIn,
		}))
	} else {
		c.JSON(http.StatusOK, JSON.IncorrectUserOrPwd())
		return
	}
	{{else}}
	username := config.String("auth.user")
	password := config.String("auth.password")
	if username != t.Username {
		c.JSON(http.StatusOK, JSON.IncorrectUserOrPwd())
		return
	}
	if password == t.Password {
		token, err := GenerateToken(0, app.JWTExpireIn)
		if err != nil {
			c.JSON(http.StatusOK, JSON.Error(err))
			return
		}
		c.JSON(http.StatusOK, JSON.OK(gin.H{
			"token":    token,
			"expirein": app.JWTExpireIn,
		}))
	} else {
		c.JSON(http.StatusOK, JSON.IncorrectUserOrPwd())
		return
	}
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
