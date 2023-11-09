package preset

import (
	"testing"
)

func TestRunZapExampleLoggerDemo(t *testing.T) {
	if err := runZapExampleLoggerDemo(); err != nil {
		t.Fatal(err)
	}
}

func TestRunZapDevelopmentLoggerDemo(t *testing.T) {
	if err := runZapDevelopmentLoggerDemo(); err != nil {
		t.Fatal(err)
	}
}

func TestRunZapProductionLoggerDemo(t *testing.T) {
	if err := runZapProductionLoggerDemo(); err != nil {
		t.Fatal(err)
	}
}
