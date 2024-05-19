package cinmemory

import (
	"errors"
)

var (
	ErrStrategyInvalid = errors.New("specified Strategy is invalid")
	ErrLoggerIsNil     = errors.New("logger is nil")
)

//go:generate mkdir -p mock
//go:generate minimock -o ./mock/ -s .go -g -i Logger
type Logger interface {
	Debug(msg string, data map[string]any)
	Error(msg string, data map[string]any)
	Info(msg string, data map[string]any)
	Warn(msg string, data map[string]any)
}

type CinMemory struct {
	strategy Strategy
	logger   Logger
}

func New(strategy Strategy, logger Logger) (*CinMemory, error) {
	if !strategy.IsValid() {
		return nil, ErrStrategyInvalid
	}
	if logger == nil {
		return nil, ErrLoggerIsNil
	}

	return &CinMemory{
		strategy: strategy,
		logger:   logger,
	}, nil
}
