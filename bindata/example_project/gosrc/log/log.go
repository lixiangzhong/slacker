package log

import (
	"bytes"
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
	Init(infofile, errfile, zap.WithCaller(true), zap.AddCallerSkip(1))
}

func Init(infoWriter, errWriter io.Writer, options ...zap.Option) {
	encoder := zapcore.EncoderConfig{
		TimeKey:       "time",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		MessageKey:    "msg",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.LowercaseColorLevelEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	debugenc := zapcore.NewConsoleEncoder(encoder)
	infoenc := encoder
	infoenc.EncodeLevel = zapcore.LowercaseLevelEncoder
	//infoenc.CallerKey = "" //不打印caller
	errenc := encoder
	errenc.EncodeLevel = zapcore.LowercaseLevelEncoder
	core := zapcore.NewTee(
		zapcore.NewCore(debugenc, zapcore.Lock(os.Stdout), zap.LevelEnablerFunc(func(z zapcore.Level) bool {
			if zapcore.DebugLevel == level.Level() {
				return true
			}
			return false
		})),
		zapcore.NewCore(zapcore.NewJSONEncoder(infoenc), zapcore.AddSync(infoWriter), enableLevel(level, zapcore.InfoLevel)),
		zapcore.NewCore(zapcore.NewJSONEncoder(errenc), zapcore.AddSync(errWriter), enableLevel(level, zapcore.ErrorLevel, zapcore.FatalLevel, zapcore.PanicLevel)),
	)
	logger = zap.New(core, options...)
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

func Println(msg ...interface{}) {
	buf := new(bytes.Buffer)
	for i := range msg {
		if i > 0 {
			buf.WriteByte(' ')
		}
		buf.WriteString("%v")
	}
	sugar.Infof(buf.String(), msg...)
}

func Fatal(args ...interface{}) {
	sugar.Fatal(args...)
}

func Error(msg ...interface{}) {
	buf := new(bytes.Buffer)
	for i := range msg {
		if i > 0 {
			buf.WriteByte(' ')
		}
		buf.WriteString("%v")
	}
	sugar.Errorf(buf.String(), msg...)
}

func Debug(msg ...interface{}) {
	buf := new(bytes.Buffer)
	for i := range msg {
		if i > 0 {
			buf.WriteByte(' ')
		}
		buf.WriteString("%v")
	}
	sugar.Debugf(buf.String(), msg...)
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
