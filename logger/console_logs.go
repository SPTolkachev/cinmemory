package logger

import (
	"fmt"
	"time"
)

const (
	defaultColor = "\033[0m"
	debugColor   = "\033[3m"
	infoColor    = "\033[37m"
	warnColor    = "\033[33m"
	errorColor   = "\033[31m"
)

func PrintToConsoleDebugFn(c *Config) {
	c.DebugFn = func(msg string, data map[string]any) {
		fmt.Printf("%s%s\t%s\t%v%s\n", debugColor, time.Now().Format(time.DateTime), msg, data, defaultColor)
	}
}

func PrintToConsoleInfoFn(c *Config) {
	c.InfoFn = func(msg string, data map[string]any) {
		fmt.Printf("%s%s\t%s\t%v%s\n", infoColor, time.Now().Format(time.DateTime), msg, data, defaultColor)
	}
}

func PrintToConsoleWarnFn(c *Config) {
	c.WarnFn = func(msg string, data map[string]any) {
		fmt.Printf("%s%s\t%s\t%v%s\n", warnColor, time.Now().Format(time.DateTime), msg, data, defaultColor)
	}
}

func PrintToConsoleErrorFn(c *Config) {
	c.ErrorFn = func(msg string, data map[string]any) {
		fmt.Printf("%s%s\t%s\t%v%s\n", errorColor, time.Now().Format(time.DateTime), msg, data, defaultColor)
	}
}
