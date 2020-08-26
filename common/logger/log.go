package logger

import (
	"go.uber.org/zap"
)

var zapLog *zap.Logger

func init () {
	zapLog, _ = zap.NewDevelopment()

}

func Infox(msg string) {
	zapLog.Info(msg)
}
