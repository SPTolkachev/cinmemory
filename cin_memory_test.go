package cinmemory

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SPTolkachev/cinmemory/mock"
)

func TestNew(t *testing.T) { //nolint:funlen
	logger := mock.NewLoggerMock(t)

	type args struct {
		strategy Strategy
		logger   Logger
	}
	type expectedResult struct {
		cinMemory *CinMemory
		err       error
	}

	tests := []struct {
		name           string
		args           args
		expectedResult expectedResult
	}{
		{
			name: "checking standard behavior",
			args: args{
				strategy: Non,
				logger:   logger,
			},
			expectedResult: expectedResult{
				cinMemory: &CinMemory{
					strategy: Non,
					logger:   logger,
				},
				err: nil,
			},
		},
		{
			name: "checking invalid strategy",
			args: args{
				strategy: Strategy(255),
				logger:   logger,
			},
			expectedResult: expectedResult{
				cinMemory: nil,
				err:       ErrStrategyInvalid,
			},
		},
		{
			name: "checking nil logger",
			args: args{
				strategy: Non,
				logger:   nil,
			},
			expectedResult: expectedResult{
				cinMemory: nil,
				err:       ErrLoggerIsNil,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			result, err := New(test.args.strategy, test.args.logger)

			if test.expectedResult.err != nil {
				assert.ErrorIs(t, err, test.expectedResult.err, "The received error does not match the expected one")
				assert.Nilf(t, result, "When an error is returned, the method should not return a value")
			} else {
				assert.NoError(t, err, "The method returned an error")
				assert.EqualValues(
					t, test.expectedResult.cinMemory, result,
					"Level's value don't equal",
				)
			}
		})
	}
}
