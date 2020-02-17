package apilogger

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//InitLogger ...
func InitLogger(logLevel string, logOutput string) (*zap.Logger, error) {

	logConfig := zap.Config{
		OutputPaths: []string{getOutput(logOutput)},
		Level:       zap.NewAtomicLevelAt(getLevel(logLevel)),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "msg",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	return logConfig.Build()

}

func getLevel(logLevel string) zapcore.Level {
	switch strings.ToLower(strings.TrimSpace(os.Getenv(logLevel))) {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "error":
		return zap.ErrorLevel
	default:
		return zap.InfoLevel
	}
}

func getOutput(logOutput string) string {
	switch strings.TrimSpace(os.Getenv(logOutput)) {
	case "stdout":
		return "stdout"
	//TODO
	default:
		return "stdout"
	}
}
