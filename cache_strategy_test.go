package cinmemory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	tests := []struct {
		name           string
		strategy       Strategy
		expectedResult bool
	}{
		{
			name:           "Check Non strategy",
			strategy:       Non,
			expectedResult: true,
		},
		{
			name:           "Check unvalid strategy",
			strategy:       Strategy(255),
			expectedResult: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			result := test.strategy.IsValid()
			assert.EqualValues(tt, test.expectedResult, result)
		})
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		name           string
		strategy       Strategy
		expectedResult string
	}{
		{
			name:           "Check Non strategy",
			strategy:       Non,
			expectedResult: "Non",
		},
		{
			name:           "Check unvalid strategy",
			strategy:       Strategy(255),
			expectedResult: "",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			result := test.strategy.String()
			assert.EqualValues(tt, test.expectedResult, result)
		})
	}
}
