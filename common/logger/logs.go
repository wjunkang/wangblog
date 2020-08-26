package logger

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"strings"
	"time"
)

var (
	sugarLogger *zap.SugaredLogger
	zapLogger *zap.Logger
)

func init () {
	encoder := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		MessageKey: "msg",
		LevelKey: "level",
		TimeKey: "time",
		NameKey: "logger",
		CallerKey: "caller",
		StacktraceKey: "stacktrace",
		LineEnding: zapcore.DefaultLineEnding,
		EncodeLevel: zapcore.LowercaseLevelEncoder,
		EncodeTime: zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller: zapcore.FullCallerEncoder,
		EncodeName: zapcore.FullNameEncoder,
	})
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.InfoLevel
	})

	infoWrite := getWriter("./info.log")
	errWrite := getWriter("./error.log")
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(infoWrite)), infoLevel),
		zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(errWrite)), zap.NewAtomicLevelAt(zap.ErrorLevel)))
	zapLogger = zap.New(core)
	sugarLogger = zapLogger.Sugar()
}

func getWriter(filename string) io.Writer {
	hook, err := rotatelogs.New(
		strings.Replace(filename, ".log", "", -1)+"-%Y-%m-%d.log",
		//文件最大寿命30天
		rotatelogs.WithMaxAge(time.Hour*24*30),
		//每隔1天进行文件分割
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	if err != nil {
		panic(err)
	}
	return hook
}

func Info(args ...interface{}) {
	sugarLogger.Info(args)
}

func Infof(template string, args ...interface{}) {
	sugarLogger.Infof(template, args)
}

func Error(args ...interface{}) {
	sugarLogger.Error(args)
}

func Errorf(template string, args ...interface{}) {
	sugarLogger.Errorf(template, args)
}


