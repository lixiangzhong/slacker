package app

import (
	"crypto/md5"
	"dns.com/ini"
	"dns.com/log"
	"encoding/hex"
	"flag"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"net"
	"os"
	"time"
)

var (
	Engine         *gin.Engine
	JWTSecret            = "{{.ProjectName}}."
	JWTExpireIn    int64 = 86400 //jwt token 过期时间,秒, 可config.ini配置
	configFileName string
	Version        string = "unknown"
	BuildTime      string
	GitVersion     string
)

func init() {
	initFlag()
	initCfg()
	initEngine()
}

func initFlag() {
	flag.StringVar(&configFileName, "c", "config.ini", "指定配置文件")
	var printverion bool
	flag.BoolVar(&printverion, "v", false, "查看版本号")
	flag.Parse()
	if printverion {
		fmt.Println("version:", Version)
		fmt.Println("编译时间:", BuildTime)
		fmt.Println("git rev:", GitVersion)
		os.Exit(0)
	}
}

func initCfg() {
	ini.Load(configFileName)
	JWTExpireIn = ini.Int64("jwt", "expirein")
	JWTSecret += hostSecret()
}

func initEngine() {
	if ini.Bool("gin", "debug") {
		Engine = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		Engine = gin.New()
		Engine.Use(gin.Recovery())
		Engine.Use(Logger())
	}
	if ini.Bool("gin", "gzip") {
		Engine.Use(gzip.Gzip(gzip.DefaultCompression))
	}
	if ini.Bool("cors", "enable") {
		Engine.Use(cors.New(cors.Config{
			AllowOriginFunc:  func(origin string) bool { return true },
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
			AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}))
	}
	Engine.Static("/static", "dist/static")
	Engine.LoadHTMLFiles("dist/index.html")
	// Engine.NoRoute(func(c *gin.Context) {
	// 	c.HTML(200, "index.html", nil)
	// })

	Engine.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
}

func Run() {
	port := ini.String("gin", "port")
	if !ini.Bool("gin", "debug") {
		log.Info("listen", port)
	}
	Engine.Run(port)
}

func Logger() gin.HandlerFunc {
	log.Config.FileName = "access.log"
	log.Config.Daily = true
	log.Config.Rotate = true
	log.Config.MaxDays = 3
	log.Config.Apply()
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		c.Next()

		end := time.Now()
		latency := end.Sub(start).Truncate(time.Millisecond)
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		if raw != "" {
			path = path + "?" + raw
		}
		uid, _ := c.Get("uid")
		log.Printf("[GIN] |%v|%8v|%6v:%15s|%-7s|%s", statusCode, latency, uid, clientIP, method, path)

	}
}

func hostSecret() string {
	itfs, err := net.Interfaces()
	if err != nil {
		return ""
	}
	m := md5.New()
	for _, itf := range itfs {
		m.Write([]byte(itf.HardwareAddr.String()))
	}
	return hex.EncodeToString(m.Sum(nil))
}
