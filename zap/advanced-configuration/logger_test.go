package advanced_configuration

import (
	"go.uber.org/zap"
	"testing"
)

func TestNewLoggerWithLumberjack(t *testing.T) {
	lumberjackOption := LumberJackOptions{
		FileName:   "logs/lumberjack.log",
		MaxSize:    10,
		MaxAge:     1,
		MaxBackups: 0,
	}
	logger, err := NewLoggerWithLumberjack(lumberjackOption)
	if err != nil {
		t.Fatal(err)
	}

	{
		logger.Info("Hello")
		logger.Info("Hello World", zap.String("world", "say hello"))
	}
}

func TestNewLoggerWithMultipleWriters(t *testing.T) {
	lumberjackOption := LumberJackOptions{
		FileName:   "logs/lumberjack.log",
		MaxSize:    10,
		MaxAge:     1,
		MaxBackups: 0,
	}
	logger, err := NewLoggerWithMultipleWriters(lumberjackOption)
	if err != nil {
		t.Fatal(err)
	}

	{
		logger.Info("Hello")
		logger.Info("Hello World", zap.String("world", "say hello"))
	}
}

func TestNewLoggerWithDifferentLevelEnabler(t *testing.T) {
	lumberjackOption := LumberJackOptions{
		FileName:   "logs/lumberjack.log",
		MaxSize:    10,
		MaxAge:     1,
		MaxBackups: 0,
	}
	logger, err := NewLoggerWithDifferentLevelEnabler(lumberjackOption)
	if err != nil {
		t.Fatal(err)
	}

	// test info output, logs below should not be found in file.
	{
		logger.Info("Hello")
		logger.Info("Hello World", zap.String("world", "say hello"))
	}

	// test error output, logs below should only be found in file.
	{
		logger.Error("This is an error message")
		logger.Error("Error occurred", zap.String("err", "Gotcha"))
	}
}
