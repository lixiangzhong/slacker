package log

import (
	"io"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	logger *zap.Logger
	sugar  *zap.SugaredLogger
	level  = zap.NewAtomicLevelAt(zap.DebugLevel)
)

func init() {
	infofile := &lumberjack.Logger{
		Filename:   "info.log",
		MaxSize:    100, //单位:MB
		MaxBackups: 3,   //保留n个文件
		LocalTime:  true,
		Compress:   false,
	}
	errfile := &lumberjack.Logger{
		Filename:   "error.log",
		MaxSize:    100, //单位:MB
		MaxBackups: 7,   //保留n个文件
		LocalTime:  true,
		Compress:   false,
	}
	Init(infofile, errfile)
}

func Init(infoWriter, errWriter io.Writer) {
	enccfg := zap.NewProductionEncoderConfig()
	enccfg.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05"))
	}
	enc := zapcore.NewConsoleEncoder(enccfg)
	jsonenc := zapcore.NewJSONEncoder(enccfg)
	core := zapcore.NewTee(
		zapcore.NewCore(enc, zapcore.Lock(os.Stdout), level),
		zapcore.NewCore(jsonenc, zapcore.AddSync(infoWriter), enableLevel(level, zapcore.InfoLevel)),
		zapcore.NewCore(jsonenc, zapcore.AddSync(errWriter), enableLevel(level, zapcore.ErrorLevel, zapcore.FatalLevel, zapcore.PanicLevel)),
	)

	logger = zap.New(core, zap.WithCaller(true), zap.AddCallerSkip(1))
	sugar = logger.Sugar()
}

func enableLevel(lv zap.AtomicLevel, want ...zapcore.Level) zapcore.LevelEnabler {
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

func SetLevel(lv zapcore.Level) {
	level.SetLevel(lv)
}

func Println(args ...interface{}) {
	sugar.Info(args...)
}

func Fatal(args ...interface{}) {
	sugar.Fatal(args...)
}

func Error(args ...interface{}) {
	sugar.Error(args...)
}

func Debug(args ...interface{}) {
	sugar.Debug(args...)
}

func DebugWithField(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

func ErrorWithField(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

func DebugWith(msg interface{}, args ...interface{}) {
	sugar.With(args...).Debug(msg)
}

func ErrorWith(msg interface{}, args ...interface{}) {
	sugar.With(args...).Error(msg)
}

func Named(name string) *zap.Logger {
	return logger.Named(name)
}
