package logger

import "go.uber.org/zap"

type Logger struct {
}

var internalLogger *zap.SugaredLogger

func init() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	internalLogger = logger.Sugar()
}

func (log *Logger) Info(msg string, fields interface{}) {
	internalLogger.Info(msg, fields)
}

func (log *Logger) Warn(msg string, fields interface{}) {
	internalLogger.Warn(msg, fields)
}
func (log *Logger) Error(msg string, fields interface{}) {
	internalLogger.Error(msg, fields)
}
