package advanced_configuration

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

type LumberJackOptions struct {
	// FileName is the absolute path of file that stores logs.
	FileName string

	// MaxSize is the maximum size of log file, Megabytes as unit.
	MaxSize int

	// MaxAge is the maximum age of log file, Day as unit.
	MaxAge int

	// MaxBackups is the maximum number of expired files.
	MaxBackups int
}

// NewLoggerWithLumberjack shows how to combine Zap and Lumberjack to
// create a logger that can do both file output and file management.
// More details about lumberjack, please see: https://github.com/natefinch/lumberjack
func NewLoggerWithLumberjack(opt LumberJackOptions) (*zap.Logger, error) {
	lumberjackLogger := &lumberjack.Logger{
		Filename:   opt.FileName,
		MaxSize:    opt.MaxSize,
		MaxAge:     opt.MaxAge,
		MaxBackups: opt.MaxBackups,
	}

	newCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(getEncoderConfig()),
		zapcore.AddSync(lumberjackLogger),
		zap.InfoLevel)

	return zap.New(newCore), nil
}

// NewLoggerWithMultipleWriters shows how to create a logger with multiple writers,
// which means one log will be sent to multiple destinations. Function below uses
// lumberjack and console to build a MultiWriteSyncer.
func NewLoggerWithMultipleWriters(opt LumberJackOptions) (*zap.Logger, error) {
	lumberjackLogger := &lumberjack.Logger{
		Filename:   opt.FileName,
		MaxSize:    opt.MaxSize,
		MaxAge:     opt.MaxAge,
		MaxBackups: opt.MaxBackups,
	}

	multipleWriters := zapcore.NewMultiWriteSyncer(
		zapcore.AddSync(lumberjackLogger),
		zapcore.AddSync(os.Stdout))

	newCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(getEncoderConfig()),
		multipleWriters,
		zap.InfoLevel)

	return zap.New(newCore), nil
}

// NewLoggerWithDifferentLevelEnabler shows how to use zapcore.NewTee to create
// a logger with multiple cores, and it also shows how to create different cores
// by using different levelEnabler. Logger created by function below will:
//
// - output log with level >= zapcore.ErrorLevel only to file, and not to console.
// - output log with level < zapcore.ErrorLevel to console, and not to file.
func NewLoggerWithDifferentLevelEnabler(opt LumberJackOptions) (*zap.Logger, error) {
	lumberjackLogger := &lumberjack.Logger{
		Filename:   opt.FileName,
		MaxSize:    opt.MaxSize,
		MaxAge:     opt.MaxAge,
		MaxBackups: opt.MaxBackups,
	}

	highPriority := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zapcore.ErrorLevel
	})

	lowPriority := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level < zapcore.ErrorLevel
	})

	lumberjackCore := zapcore.NewCore(zapcore.NewJSONEncoder(getEncoderConfig()), zapcore.AddSync(lumberjackLogger), highPriority)
	consoleCore := zapcore.NewCore(zapcore.NewJSONEncoder(getEncoderConfig()), zapcore.AddSync(os.Stdout), lowPriority)

	newCore := zapcore.NewTee(lumberjackCore, consoleCore)
	return zap.New(newCore), nil
}

// getEncoderConfig returns default encoder config
func getEncoderConfig() (config zapcore.EncoderConfig) {
	config = zapcore.EncoderConfig{
		MessageKey:  "message",
		LevelKey:    "level",
		TimeKey:     "time",
		NameKey:     "logger",
		CallerKey:   "caller",
		LineEnding:  zapcore.DefaultLineEnding,
		EncodeLevel: zapcore.LowercaseLevelEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	return config
}
