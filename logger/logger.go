package logger

import (
	"errors"
)

var (
	ErrLevelInvalid = errors.New("specified logger level is invalid")
)

type Configurators func(c *Config)
type LogFn func(msg string, data map[string]any)

type Logger struct {
	level   Level
	debugFn LogFn
	infoFn  LogFn
	warnFn  LogFn
	errorFn LogFn
}

func New(level Level, args ...Configurators) (*Logger, error) {
	if !level.IsValid() {
		return nil, ErrLevelInvalid
	}

	logger := Logger{
		level: level,
	}

	configure(&logger, args...)

	return &logger, nil
}

func (l *Logger) Debug(msg string, data map[string]any) {
	if l.debugFn == nil || l.level > Debug {
		return
	}

	l.debugFn(msg, data)
}

func (l *Logger) Info(msg string, data map[string]any) {
	if l.infoFn == nil || l.level > Info {
		return
	}

	l.infoFn(msg, data)
}

func (l *Logger) Warn(msg string, data map[string]any) {
	if l.warnFn == nil || l.level > Warn {
		return
	}

	l.warnFn(msg, data)
}

func (l *Logger) Error(msg string, data map[string]any) {
	if l.errorFn == nil || l.level > Error {
		return
	}

	l.errorFn(msg, data)
}
