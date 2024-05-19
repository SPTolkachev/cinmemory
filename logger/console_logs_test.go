package logger

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPrintToConsoleDebugFn(t *testing.T) {
	logger := Logger{
		level: Debug,
	}

	PrintToConsoleDebugFn(&logger)
	require.NotNil(t, logger.DebugFn)
	logger.DebugFn(LogInfo{"test debug level", nil})
}

func TestPrintToConsoleInfoFn(t *testing.T) {
	logger := Logger{
		level: Debug,
	}

	PrintToConsoleInfoFn(&logger)
	require.NotNil(t, logger.InfoFn)
	logger.InfoFn(LogInfo{"test info level", nil})
}

func TestPrintToConsoleWarnFn(t *testing.T) {
	logger := Logger{
		level: Debug,
	}

	PrintToConsoleWarnFn(&logger)
	require.NotNil(t, logger.WarnFn)
	logger.WarnFn(LogInfo{"test warn level", nil})
}

func TestPrintToConsoleErrorFn(t *testing.T) {
	logger := Logger{
		level: Debug,
	}

	PrintToConsoleErrorFn(&logger)
	require.NotNil(t, logger.ErrorFn)
	logger.ErrorFn(LogInfo{"test error level", nil})
}
