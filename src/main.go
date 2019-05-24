package main

import (
	l "github.com/BNPrashanth/goexample-uber-zap/src/logger"
	"go.uber.org/zap"
)

func main() {

	l.Log.Debug("message param", zap.String("Message", "This is a DEBUG message"))
	l.Log.Info("This is an INFO message")
	l.Log.Info("This is an INFO message with fields", zap.String("region", "us-west"), zap.Int("id", 2))

}
