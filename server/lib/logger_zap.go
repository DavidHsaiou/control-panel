package lib

import (
	"fmt"

	"go.uber.org/zap"

	"control-panel/server/lib/inf"
)

type loggerZap struct {
	loggerName string
	logger     *zap.Logger
	sugar      *zap.SugaredLogger
}

func newLoggerZap(name string) inf.ILogger {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(fmt.Sprintf("can not init logger, %+v", err))
	}

	return &loggerZap{
		loggerName: name,
		logger:     logger,
		sugar:      logger.Sugar(),
	}
}

func (l *loggerZap) Debug(msg string) {
	l.sugar.Debug(msg)
}

func (l *loggerZap) Info(msg string) {
	l.sugar.Info(msg)
}

func (l *loggerZap) Warn(msg string) {
	l.sugar.Warn(msg)
}

func (l *loggerZap) Error(msg string) {
	l.sugar.Error(msg)
}

func (l *loggerZap) FlushLogs() {
	err := l.logger.Sync()
	if err != nil {
		panic(fmt.Sprintf("can not flush logs from buffer, err: %+v", err))
	}
}
