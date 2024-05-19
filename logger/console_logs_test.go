package logger

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPrintToConsoleDebugFn(t *testing.T) {
	config := Config{}

	PrintToConsoleDebugFn(&config)
	require.NotNil(t, config.DebugFn)
	config.DebugFn("test debug level", nil)
}

func TestPrintToConsoleInfoFn(t *testing.T) {
	config := Config{}

	PrintToConsoleInfoFn(&config)
	require.NotNil(t, config.InfoFn)
	config.InfoFn("test info level", nil)
}

func TestPrintToConsoleWarnFn(t *testing.T) {
	config := Config{}

	PrintToConsoleWarnFn(&config)
	require.NotNil(t, config.WarnFn)
	config.WarnFn("test warn level", nil)
}

func TestPrintToConsoleErrorFn(t *testing.T) {
	config := Config{}

	PrintToConsoleErrorFn(&config)
	require.NotNil(t, config.ErrorFn)
	config.ErrorFn("test error level", nil)
}
