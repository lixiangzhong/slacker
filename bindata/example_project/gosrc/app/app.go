package app

import (
	"github.com/jmoiron/sqlx"
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
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	_ "{{.ProjectPath}}/gosrc/validator"
	"net"
	"os"
	"time"
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
)

func Init() {
	initFlag()
	initCfg()
	initEngine()
}

func initFlag() {
	flag.StringVar(&configFileName, "c", "config.yaml", "指定配置文件")
	var printverion bool
	flag.BoolVar(&printverion, "v", false, "查看版本号")
	flag.Parse()
	if printverion {
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
	if config.Bool("gin.debug") {
		Engine = gin.Default()
	} else {
		log.Set(func(info, err, debug log.LoggerSetter) {
			info.SetOutput(rotatefile.New("info.log", 3, true))
			err.SetOutput(rotatefile.New("error.log", 7, true))
		})
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
	Engine.Static("/static", "dist/static") 
	Engine.LoadHTMLFiles("dist/index.html")
	Engine.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
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
}
func Logger() gin.HandlerFunc {
	var out zapcore.WriteSyncer
	loglevel := zap.NewAtomicLevelAt(zap.InfoLevel)
	switch config.String("gin.log") {
	case "file":
		f := &lumberjack.Logger{
			Filename:   "access.log",
			MaxSize:    10, //10m
			MaxBackups: 10, //保留10个文件
			LocalTime:  true,
			Compress:   false,
		}
		out = zapcore.AddSync(f)
	case "stdout":
		loglevel.SetLevel(zap.DebugLevel)
		out = zapcore.Lock(os.Stdout)
	default:
		log.Fatal("bad log output")
	}
	enccfg := zap.NewProductionEncoderConfig()
	enccfg.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05"))
	}
	enc := zapcore.NewConsoleEncoder(enccfg)
	core := zapcore.NewTee(
		// zapcore.NewCore(enc, zapcore.Lock(os.Stdout), DebugLevelEnalbe(loglevel, zap.DebugLevel, zap.InfoLevel)),
		zapcore.NewCore(enc, out, NewLevelEnable(loglevel, zap.InfoLevel)),
	)
	accesslog := zap.New(core).Sugar()
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
		scheme := c.Request.URL.Scheme
		if scheme == "" {
			scheme = "http"
		}
		if raw != "" {
			path = path + "?" + raw
		}
		uid, _ := c.Get("uid") 
		accesslog.Infof("%v|%8v|%6v:%15s|%v|%-7s|%s", statusCode, latency, uid, clientIP, scheme, method, path)
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

func NewLevelEnable(lv zap.AtomicLevel, want ...zapcore.Level) zapcore.LevelEnabler {
	return zap.LevelEnablerFunc(func(l zapcore.Level) (ok bool) {
		for _, v := range want {
			if v == l && lv.Enabled(l) {
				ok = true
				break
			}
		}
		return
	})
}

func DebugLevelEnalbe(lv zap.AtomicLevel, want ...zapcore.Level) zapcore.LevelEnabler {
	return zap.LevelEnablerFunc(func(l zapcore.Level) (ok bool) {
		for _, v := range want {
			if v == l && lv.Enabled(zap.DebugLevel) {
				ok = true
				break
			}
		}
		return
	})
}
