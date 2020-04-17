package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"

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

const (
	WebsocketWriteWait      = 10 * time.Second
	WebsokcetPongWait       = 60 * time.Second
	WebsocketPingPeriod     = (WebsokcetPongWait * 9) / 10
	WebsocketMaxMessageSize = 512
)

/*
go writews(conn, ch)
	go func() {
		for range tk.C {
			b, err := json.Marshal(struct)
			if err != nil {
				continue
			}
			ch <- b
		}
	}()
	readws(conn)
*/

func readws(c *websocket.Conn) {
	defer func() {
		c.Close()
	}()
	c.SetReadLimit(WebsocketMaxMessageSize)
	c.SetReadDeadline(time.Now().Add(WebsokcetPongWait))
	c.SetPongHandler(func(string) error {
		c.SetReadDeadline(time.Now().Add(WebsokcetPongWait))
		return nil
	})
	for {
		_, _, err := c.ReadMessage()
		if err != nil {
			break
		}
	}
}

func writews(c *websocket.Conn, ch chan []byte) {
	tk := time.NewTicker(WebsocketPingPeriod)
	for {
		select {
		case b, ok := <-ch:
			c.SetWriteDeadline(time.Now().Add(WebsocketWriteWait))
			if !ok {
				c.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			w, err := c.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(b)
			if err := w.Close(); err != nil {
				return
			}
		case <-tk.C:
			c.SetWriteDeadline(time.Now().Add(WebsocketWriteWait))
			if err := c.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}

	}
}
