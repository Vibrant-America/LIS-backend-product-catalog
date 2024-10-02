package common

import (
	"log"
	"os"

	zapLogger "github.com/asim/go-micro/plugins/logger/zap/v4"
	"go-micro.dev/v4/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var encoderConfig = zapcore.EncoderConfig{
	MessageKey: "message",

	LevelKey:    "level",
	EncodeLevel: zapcore.CapitalLevelEncoder,

	TimeKey:    "time",
	EncodeTime: zapcore.ISO8601TimeEncoder,

	CallerKey:    "caller",
	EncodeCaller: zapcore.ShortCallerEncoder,
}

// Init Logger
var zapConfig = zap.Config{
	Encoding:      "json",
	Level:         zap.NewAtomicLevelAt(zapcore.InfoLevel),
	OutputPaths:   []string{"stdout"},
	EncoderConfig: encoderConfig,
}

func InitLogger() logger.Logger {

	newLogger, err := zapLogger.NewLogger(zapLogger.WithConfig(zapConfig))
	if err != nil {
		log.Fatal(err)
	}
	if os.Getenv("LOG_LEVEL") == "DEBUG" {
		zapConfig.Level.SetLevel(zapcore.DebugLevel)
	}
	return newLogger

}

var ZapLogger, _ = zapConfig.Build()
