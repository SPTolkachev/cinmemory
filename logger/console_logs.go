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

func PrintToConsoleDebugFn(l *Logger) {
	l.DebugFn = func(info LogInfo) {
		fmt.Printf("%s%s\t%s\t%v%s\n", debugColor, time.Now().Format(time.DateTime), info.Msg, info.Data, defaultColor)
	}
}

func PrintToConsoleInfoFn(l *Logger) {
	l.InfoFn = func(info LogInfo) {
		fmt.Printf("%s%s\t%s\t%v%s\n", infoColor, time.Now().Format(time.DateTime), info.Msg, info.Data, defaultColor)
	}
}

func PrintToConsoleWarnFn(l *Logger) {
	l.WarnFn = func(info LogInfo) {
		fmt.Printf("%s%s\t%s\t%v%s\n", warnColor, time.Now().Format(time.DateTime), info.Msg, info.Data, defaultColor)
	}
}

func PrintToConsoleErrorFn(l *Logger) {
	l.ErrorFn = func(info LogInfo) {
		fmt.Printf("%s%s\t%s\t%v%s\n", errorColor, time.Now().Format(time.DateTime), info.Msg, info.Data, defaultColor)
	}
}
