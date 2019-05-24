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
		EncoderConfig: zapcore.EncoderConfig{
			CallerKey: "caller",
			EncodeCaller:   zapcore.ShortCallerEncoder,
			MessageKey: "msg",
		},
	}
	Log, _ = conf.Build()
	Log.Info("Logger Initialized..")
}
