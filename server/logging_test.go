package server

import (
	"testing"

	"github.com/goburrow/melon/core"
	"github.com/goburrow/melon/logging"
	slogging "github.com/goburrow/melon/server/logging"
)

func TestDefaultRequestLogFactory(t *testing.T) {
	env := core.NewEnvironment()
	factory := DefaultRequestLogFactory{}
	appender := logging.AppenderConfiguration{}
	appender.SetValue(&logging.ConsoleAppenderFactory{})

	factory.Appenders = []logging.AppenderConfiguration{
		appender,
	}

	filter, err := factory.Build(env)
	if err != nil {
		t.Fatal(err)
	}
	switch filter.(type) {
	case *slogging.Filter:
	default:
		t.Fatalf("unexpected filter %#v", filter)
	}
}

func TestNoRequestLogFactory(t *testing.T) {
	env := core.NewEnvironment()
	factory := DefaultRequestLogFactory{}
	filter, err := factory.Build(env)
	if err != nil {
		t.Fatal(err)
	}
	switch filter.(type) {
	case *noRequestLog:
		if filter.Name() != "logging" {
			t.Fatalf("unexpected filter name %#v", filter.Name())
		}
	default:
		t.Fatalf("unexpected filter %#v", filter)
	}
}
