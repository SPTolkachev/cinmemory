package logger

import (
	"errors"
)

var (
	ErrLevelInvalid = errors.New("specified logger level is invalid")
)

type LogInfo struct {
	Msg  string
	Data map[string]any
}

type Args func(l *Logger)
type LogFn func(LogInfo)

type Logger struct {
	level   Level
	DebugFn LogFn
	InfoFn  LogFn
	WarnFn  LogFn
	ErrorFn LogFn
}

func New(level Level, args ...Args) (*Logger, error) {
	if !level.IsValid() {
		return nil, ErrLevelInvalid
	}

	logger := Logger{
		level: level,
	}

	for _, arg := range args {
		arg(&logger)
	}

	return &logger, nil
}

func (l *Logger) Debug(msg string, data map[string]any) {
	if l.DebugFn == nil || l.level > Debug {
		return
	}

	l.DebugFn(LogInfo{msg, data})
}

func (l *Logger) Info(msg string, data map[string]any) {
	if l.InfoFn == nil || l.level > Info {
		return
	}

	l.InfoFn(LogInfo{msg, data})
}

func (l *Logger) Warn(msg string, data map[string]any) {
	if l.WarnFn == nil || l.level > Warn {
		return
	}

	l.WarnFn(LogInfo{msg, data})
}

func (l *Logger) Error(msg string, data map[string]any) {
	if l.ErrorFn == nil || l.level > Error {
		return
	}

	l.ErrorFn(LogInfo{msg, data})
}
