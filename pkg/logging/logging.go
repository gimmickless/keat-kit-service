package util

import (
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	logMessageKey = "message"
)

func NewLogger() *otelzap.SugaredLogger {
	logcfg := zap.NewProductionConfig()
	logcfg.EncoderConfig.MessageKey = logMessageKey
	logcfg.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	logger, _ := logcfg.Build()
	otellogger := otelzap.New(logger)
	return otellogger.Sugar()
}
