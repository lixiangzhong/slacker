package app

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/lixiangzhong/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"{{.ProjectName}}/gosrc/log"
	_ "{{.ProjectName}}/gosrc/validator"
)

var (
	Engine         *gin.Engine
	httpserver     *http.Server
	JWTSecret            = "{{.ProjectName}}."
	JWTExpireIn    int64 = 86400 //jwt token 过期时间,秒, 可config.ini配置
	configFileName string
	Version        string = "unknown"
	BuildTime      string
	GitVersion     string
	GoMD5          string
	onShutdownFuncs []func()
)

func Init() {
	initFlag()
	initCfg()
	initEngine()
}

func initFlag() {
	flag.StringVar(&configFileName, "c", "config.yaml", "指定配置文件")
	var printversion bool
	flag.BoolVar(&printversion, "v", false, "查看版本号")
	flag.Parse()
	if printversion {
		fmt.Println("version:", Version)
		fmt.Println("编译时间:", BuildTime)
		fmt.Println("git rev:", GitVersion)
		fmt.Println("gofilemd5:", GoMD5)
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
	if !config.Bool("gin.debug") {
		gin.SetMode(gin.ReleaseMode)
	}
	Engine = gin.New()
	Engine.Use(gin.Recovery())
	Engine.Use(Logger())
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
}

func Run() {
	port := config.String("gin.port") 
	log.Println("listen", port)  
	httpserver = &http.Server{
		Addr:    port,
		Handler: Engine,
	}
	go func() {
		err := httpserver.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()
	waitSignal()
}

func waitSignal() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL)
	<-ch
	log.Println("stop http server...")
	ctx := context.Background()
	ctx,_ = context.WithTimeout(ctx, time.Second*10)
	err := httpserver.Shutdown(ctx)
	if err != nil {
		log.Error(err)
	}
	for _, f := range onShutdownFuncs {
		f()
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

func Logger() gin.HandlerFunc {
	enccfg := zap.NewProductionEncoderConfig()
	enccfg.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05"))
	}
	var enc zapcore.Encoder
	var out zapcore.WriteSyncer
	switch config.String("gin.log") {
	case "file":
		f := &lumberjack.Logger{
			Filename:   "access.log",
			MaxSize:    100, //100m
			MaxBackups: 5, //保留5个文件
			LocalTime:  true,
			Compress:   false,
		}
		out = zapcore.AddSync(f)
		enc = zapcore.NewJSONEncoder(enccfg)
	case "stdout":
		out = zapcore.Lock(os.Stdout)
		enc = zapcore.NewConsoleEncoder(enccfg)
	default:
		log.Fatal("bad log output")
	}
	core := zapcore.NewTee(
		zapcore.NewCore(enc, out, zap.InfoLevel),
	)
	logger := zap.New(core)
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		end := time.Now()
		latency := end.Sub(start)
		if len(c.Errors) > 0 {
			for _, e := range c.Errors.Errors() {
				logger.Error(e)
			}
		} else {
			logger.Info(path,
				zap.String("ip", c.ClientIP()),
				zap.String("method", c.Request.Method),
				zap.String("path", path),
				zap.String("query", query),
				zap.Int("status", c.Writer.Status()),
				zap.String("ua", c.Request.UserAgent()),
				zap.String("latency", latency.Truncate(time.Millisecond).String()),
			)
		}
	}
}

func OnShutdown(f func()) {
	onShutdownFuncs = append(onShutdownFuncs, f)
}