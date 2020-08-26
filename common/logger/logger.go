package logger


import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

const (
	defaultMaxSize = 512
)

func NewLogger(filePath string, level zapcore.Level, maxSize, maxAge, maxBackups int, compress bool, serviceName string) *zap.Logger {
	core := newCore(filePath, level, maxSize, maxBackups, maxAge, compress)
	return zap.New(core, zap.AddCaller(), zap.Development(), zap.Fields(zap.String("serviceName", serviceName)))
}

func newCore(filePath string, level zapcore.Level, maxSize, maxAge, maxBackups int, compress bool) zapcore.Core {
	hook := lumberjack.Logger{
		Filename: filePath,
		MaxSize: maxSize, // 单位 M
		MaxAge: maxAge, // 单位 days
		MaxBackups: maxBackups,
		Compress: compress,
	}
	atomicLevel := zap.NewAtomicLevelAt(level)
	encoderConfig := zapcore.EncoderConfig{
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
	}
	return zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), atomicLevel)
}
