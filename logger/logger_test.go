package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	type args struct {
		level         Level
		configurators []Configurators
	}
	type expectedResult struct {
		loggerLevel Level
		err         error
	}

	tests := []struct {
		name           string
		args           args
		expectedResult expectedResult
	}{
		{
			name: "checking standard behavior",
			args: args{
				level: Error,
				configurators: []Configurators{
					PrintToConsoleDebugFn, PrintToConsoleInfoFn, PrintToConsoleWarnFn, PrintToConsoleErrorFn,
				},
			},
			expectedResult: expectedResult{
				loggerLevel: Error,
				err:         nil,
			},
		},
		{
			name: "checking invalid level",
			args: args{
				level: Level(10),
				configurators: []Configurators{
					PrintToConsoleDebugFn, PrintToConsoleInfoFn, PrintToConsoleWarnFn, PrintToConsoleErrorFn,
				},
			},
			expectedResult: expectedResult{
				loggerLevel: Non,
				err:         ErrLevelInvalid,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			logger, err := New(test.args.level, test.args.configurators...)

			if test.expectedResult.err != nil {
				assert.ErrorIs(t, err, test.expectedResult.err, "The received error does not match the expected one")
				assert.Nilf(t, logger, "When an error is returned, the method should not return a value")
			} else {
				assert.NoError(t, err, "The method returned an error")
				assert.EqualValues(
					t, test.expectedResult.loggerLevel, logger.level,
					"Level's value don't equal",
				)
			}
		})
	}
}

func TestDebug(t *testing.T) { //nolint:dupl
	type args struct {
		msg  string
		data map[string]any
	}

	tests := []struct {
		name   string
		args   args
		logger *Logger
	}{
		{
			name: "checking standard behavior",
			args: args{
				msg:  "Test",
				data: map[string]any{"one": 1},
			},
			logger: &Logger{
				level: Debug,
				debugFn: func(msg string, data map[string]any) {
					assert.EqualValues(t, "Test", msg)
					assert.EqualValues(t, map[string]any{"one": 1}, data)
				},
			},
		},
		{
			name: "checking behavior for nil func",
			args: args{
				msg:  "Test",
				data: map[string]any{"one": 1},
			},
			logger: &Logger{
				level: Debug,
			},
		},
		{
			name: "checking behavior for info level",
			args: args{
				msg:  "Test",
				data: map[string]any{"one": 1},
			},
			logger: &Logger{
				level: Info,
				debugFn: func(_ string, _ map[string]any) {
					assert.Fail(t, "Unexpected call to logger.DebugFn")
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			test.logger.Debug(test.args.msg, test.args.data)
		})
	}
}

func TestInfo(t *testing.T) { //nolint:dupl
	type args struct {
		msg  string
		data map[string]any
	}

	tests := []struct {
		name   string
		args   args
		logger *Logger
	}{
		{
			name: "checking standard behavior",
			args: args{
				msg:  "Test",
				data: map[string]any{"one": 1},
			},
			logger: &Logger{
				level: Info,
				infoFn: func(msg string, data map[string]any) {
					assert.EqualValues(t, "Test", msg)
					assert.EqualValues(t, map[string]any{"one": 1}, data)
				},
			},
		},
		{
			name: "checking behavior for nil func",
			args: args{
				msg:  "Test",
				data: map[string]any{"one": 1},
			},
			logger: &Logger{
				level: Info,
			},
		},
		{
			name: "checking behavior for info level",
			args: args{
				msg:  "Test",
				data: map[string]any{"one": 1},
			},
			logger: &Logger{
				level: Warn,
				infoFn: func(_ string, _ map[string]any) {
					assert.Fail(t, "Unexpected call to logger.InfoFn")
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			test.logger.Info(test.args.msg, test.args.data)
		})
	}
}

func TestWarn(t *testing.T) { //nolint:dupl
	type args struct {
		msg  string
		data map[string]any
	}

	tests := []struct {
		name   string
		args   args
		logger *Logger
	}{
		{
			name: "checking standard behavior",
			args: args{
				msg:  "Test",
				data: map[string]any{"one": 1},
			},
			logger: &Logger{
				level: Warn,
				warnFn: func(msg string, data map[string]any) {
					assert.EqualValues(t, "Test", msg)
					assert.EqualValues(t, map[string]any{"one": 1}, data)
				},
			},
		},
		{
			name: "checking behavior for nil func",
			args: args{
				msg:  "Test",
				data: map[string]any{"one": 1},
			},
			logger: &Logger{
				level: Warn,
			},
		},
		{
			name: "checking behavior for warn level",
			args: args{
				msg:  "Test",
				data: map[string]any{"one": 1},
			},
			logger: &Logger{
				level: Error,
				warnFn: func(_ string, _ map[string]any) {
					assert.Fail(t, "Unexpected call to logger.WarnFn")
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			test.logger.Warn(test.args.msg, test.args.data)
		})
	}
}

func TestError(t *testing.T) {
	type args struct {
		msg  string
		data map[string]any
	}

	tests := []struct {
		name   string
		args   args
		logger *Logger
	}{
		{
			name: "checking standard behavior",
			args: args{
				msg:  "Test",
				data: map[string]any{"one": 1},
			},
			logger: &Logger{
				level: Error,
				errorFn: func(msg string, data map[string]any) {
					assert.EqualValues(t, "Test", msg)
					assert.EqualValues(t, map[string]any{"one": 1}, data)
				},
			},
		},
		{
			name: "checking behavior for nil func",
			args: args{
				msg:  "Test",
				data: map[string]any{"one": 1},
			},
			logger: &Logger{
				level: Error,
			},
		},
		{
			name: "checking behavior for error level",
			args: args{
				msg:  "Test",
				data: map[string]any{"one": 1},
			},
			logger: &Logger{
				level: Level(5),
				errorFn: func(_ string, _ map[string]any) {
					assert.Fail(t, "Unexpected call to logger.ErrorFn")
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			test.logger.Error(test.args.msg, test.args.data)
		})
	}
}
