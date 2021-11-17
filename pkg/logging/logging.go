package util

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	logMessageKey = "message"
)

func NewLogger() *zap.SugaredLogger {
	logcfg := zap.NewProductionConfig()
	logcfg.EncoderConfig.MessageKey = logMessageKey
	logcfg.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	logger, _ := logcfg.Build()
	return logger.Sugar()
}
