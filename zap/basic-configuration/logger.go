package basic_configuration

import (
	"encoding/json"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LogLevel interface {
	ToZap() zapcore.Level
}

type logLevelFunc struct {
	level zapcore.Level
}

func (l logLevelFunc) ToZap() zapcore.Level {
	return l.level
}

func InfoLevel() LogLevel {
	return logLevelFunc{
		level: zapcore.InfoLevel,
	}
}

func DebugLevel() LogLevel {
	return logLevelFunc{
		level: zapcore.DebugLevel,
	}
}

func WarnLevel() LogLevel {
	return logLevelFunc{
		level: zapcore.WarnLevel,
	}
}

func ErrorLevel() LogLevel {
	return logLevelFunc{
		level: zapcore.ErrorLevel,
	}
}

// NewLoggerWithConfig gives an example of how to use zap.Config and zap.Must
// function to create a zap.Logger.
func NewLoggerWithConfig(level LogLevel) *zap.Logger {
	cfg := zap.Config{
		Level:             zap.NewAtomicLevelAt(level.ToZap()),
		Development:       false,
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling:          nil,
		Encoding:          "json",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:  "message",
			LevelKey:    "level",
			EncodeLevel: zapcore.LowercaseLevelEncoder,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		InitialFields: map[string]interface{}{
			"custom_field": "hello",
		},
	}

	return zap.Must(cfg.Build())
}

// NewLoggerWithJSONConfig takes parameter config to create a zap.Logger.
// Parameter config is in JSON format, you can create a logger after we parse JSON content to zap.Config.
//
// It's worth noting that some fields in zap.Config implement encoding.TextUnmarshaler interface
// so that we can create a zap.Config by JSON format.
func NewLoggerWithJSONConfig(config string) (*zap.Logger, error) {
	var cfg zap.Config
	if err := json.Unmarshal([]byte(config), &cfg); err != nil {
		return nil, err
	}
	return zap.Must(cfg.Build()), nil
}
