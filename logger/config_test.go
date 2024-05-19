package logger

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigure(t *testing.T) {
	debugFn := func(_ string, _ map[string]any) {}
	infoFn := func(_ string, _ map[string]any) {}
	warnFn := func(_ string, _ map[string]any) {}
	errorFn := func(_ string, _ map[string]any) {}

	logger := Logger{}
	configure(&logger, func(c *Config) {
		c.DebugFn = debugFn
		c.InfoFn = infoFn
		c.WarnFn = warnFn
		c.ErrorFn = errorFn
	})

	assert.EqualValues(
		t, fmt.Sprintf("%p", debugFn), fmt.Sprintf("%p", logger.debugFn),
		"DebugFn's value don't equal",
	)
	assert.EqualValues(
		t, fmt.Sprintf("%p", infoFn), fmt.Sprintf("%p", logger.infoFn),
		"InfoFn's value don't equal",
	)
	assert.EqualValues(
		t, fmt.Sprintf("%p", warnFn), fmt.Sprintf("%p", logger.warnFn),
		"WarnFn's value don't equal",
	)
	assert.EqualValues(
		t, fmt.Sprintf("%p", errorFn), fmt.Sprintf("%p", logger.errorFn),
		"ErrorFn's value don't equal",
	)
}
