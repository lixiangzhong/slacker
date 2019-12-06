package controllers

import (
	"github.com/lixiangzhong/config"
	"github.com/jmoiron/sqlx"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket" 
	"{{.ProjectName}}/gosrc/service"
)

var Service *service.Service

func InitService(db *sqlx.DB) { 
	Service=service.New(db)
}

var WebSocket = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func DefaultQueryInt(c *gin.Context, key string, defaultval int) int {
	i, err := strconv.Atoi(c.Query(key))
	if err != nil {
		i = defaultval
	}
	return i
}

func QueryUint32(c *gin.Context, key string) uint32 {
	i, _ := strconv.ParseUint(c.Query(key), 10, 32)
	return uint32(i)
}

func QueryUint16(c *gin.Context, key string) uint16 {
	i, _ := strconv.ParseUint(c.Query(key), 10, 16)
	return uint16(i)
}

func strings2ints(s []string) []int {
	var result = make([]int, 0, len(s))
	for _, v := range s {
		i, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		result = append(result, i)
	}
	return result
}

func strings2int64s(s []string) []int64 {
	var result = make([]int64, 0, len(s))
	for _, v := range s {
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			continue
		}
		result = append(result, i)
	}
	return result
}
