package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	// Log exported
	Log *zap.Logger
)

func init() {
	conf := zap.Config{
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(zapcore.DebugLevel),
		OutputPaths: []string{"./output/logger.log", "stdout"},
	}
	Log, _ = conf.Build()
}
