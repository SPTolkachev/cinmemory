package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	tests := []struct {
		name           string
		level          Level
		expectedResult bool
	}{
		{
			name:           "Check Non level",
			level:          Non,
			expectedResult: false,
		},
		{
			name:           "Check Debug level",
			level:          Debug,
			expectedResult: true,
		},
		{
			name:           "Check Info level",
			level:          Info,
			expectedResult: true,
		},
		{
			name:           "Check Warn level",
			level:          Warn,
			expectedResult: true,
		},
		{
			name:           "Check Error level",
			level:          Error,
			expectedResult: true,
		},
		{
			name:           "Check unvalid level",
			level:          Level(255),
			expectedResult: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			result := test.level.IsValid()
			assert.EqualValues(tt, test.expectedResult, result)
		})
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		name           string
		level          Level
		expectedResult string
	}{
		{
			name:           "Check Non level",
			level:          Non,
			expectedResult: "",
		},
		{
			name:           "Check Debug level",
			level:          Debug,
			expectedResult: "debug",
		},
		{
			name:           "Check Info level",
			level:          Info,
			expectedResult: "info",
		},
		{
			name:           "Check Warn level",
			level:          Warn,
			expectedResult: "warn",
		},
		{
			name:           "Check Error level",
			level:          Error,
			expectedResult: "error",
		},
		{
			name:           "Check unvalid level",
			level:          Level(10),
			expectedResult: "",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			result := test.level.String()
			assert.EqualValues(tt, test.expectedResult, result)
		})
	}
}
