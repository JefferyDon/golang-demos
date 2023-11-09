package basic_configuration

import (
	"go.uber.org/zap"
	"testing"
)

func TestNewLoggerWithConfig(t *testing.T) {
	logger := NewLoggerWithConfig(InfoLevel())

	logger.Info("Hello World",
		zap.String("msg", "hello"))
}

func ExampleNewLoggerWithJSONConfig() {
	config := `{
	  "level": "info",
	  "encoding": "json",
	  "outputPaths": ["stdout"],
	  "errorOutputPaths": ["stderr"],
	  "initialFields": {"custom_field": "hello"},
	  "encoderConfig": {
	    "messageKey": "message",
	    "levelKey": "level",
	    "levelEncoder": "lowercase"
	  }
	}`

	logger, err := NewLoggerWithJSONConfig(config)
	if err != nil {
		panic(err)
	}

	logger.Info("Hello World",
		zap.String("msg", "hello"))
}

func TestNewLoggerWithJSONConfig(t *testing.T) {
	logger, err := NewLoggerWithJSONConfig(`{
	  "level": "info",
	  "encoding": "json",
	  "outputPaths": ["stdout"],
	  "errorOutputPaths": ["stderr"],
	  "initialFields": {"custom_field": "hello"},
	  "encoderConfig": {
	    "messageKey": "message",
	    "levelKey": "level",
	    "levelEncoder": "lowercase"
	  }
	}`)

	if err != nil {
		t.Fatal(err)
	}

	logger.Info("Hello World",
		zap.String("msg", "hello"))
}
