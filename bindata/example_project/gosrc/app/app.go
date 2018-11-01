package app

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/lixiangzhong/config"
	"github.com/lixiangzhong/log"
	"github.com/lixiangzhong/rotatefile"
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
	config.MustLoad(configFileName)
	JWTExpireIn = config.Int64("jwt.expirein")
	JWTSecret += hostSecret()
	config.SetDefault("meta.logo", "{{.ProjectName}}")
	config.SetDefault("meta.title", "{{.ProjectName}}")
}

func initEngine() {
	if config.Bool("gin.debug") {
		Engine = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		Engine = gin.New()
		Engine.Use(gin.Recovery())
		Engine.Use(Logger())
	}
	if config.Bool("gin.gzip") {
		Engine.Use(gzip.Gzip(gzip.DefaultCompression))
	}
	if config.Bool("cors.enable") {
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
	port := config.String("gin.port")
	if !config.Bool("gin.debug") {
		log.Println("listen", port)
	}
	Engine.Run(port)
}

func Logger() gin.HandlerFunc {
	log.Set(func(info, err, debug log.LoggerSetter) {
		info.SetOutput(rotatefile.New("access.log", 3, true))
		err.SetOutput(rotatefile.New("error.log", 7, true))
	})
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
