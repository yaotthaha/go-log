package golog

import (
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
)

type Logger interface {
	Info(a ...any)
	Warn(a ...any)
	Error(a ...any)
	Debug(a ...any)
	Fatal(a ...any)
	Panic(a ...any)
	Infof(format string, a ...any)
	Warnf(format string, a ...any)
	Errorf(format string, a ...any)
	Debugf(format string, a ...any)
	Fatalf(format string, a ...any)
	Panicf(format string, a ...any)
	LevelPrintFunc(level logLevel) func(a ...any)
	LevelPrintfFunc(level logLevel) func(format string, a ...any)
	LevelPrint(level logLevel, a ...any)
	LevelPrintf(level logLevel, format string, a ...any)
}

var (
	DefaultOutput      = os.Stdout
	DefaultErrorOutput = os.Stderr
)

type logLevel string

const (
	Info  logLevel = "Info"
	Warn  logLevel = "Warn"
	Error logLevel = "Error"
	Debug logLevel = "Debug"
	Fatal logLevel = "Fatal"
	Panic logLevel = "Panic"
)

func LevelOutput(level logLevel) string {
	switch level {
	case Info:
		return "[Info] "
	case Warn:
		return "[Warn] "
	case Error:
		return "[Error]"
	case Debug:
		return "[Debug]"
	case Fatal:
		return "[Fatal]"
	case Panic:
		return "[Panic]"
	default:
		return "[Info] "
	}
}

func LevelColorOutput(level logLevel) string {
	switch level {
	case Info:
		return "[" + color.GreenString(string(Info)) + "] "
	case Warn:
		return "[" + color.YellowString(string(Warn)) + "] "
	case Error:
		return "[" + color.RedString(string(Error)) + "]"
	case Debug:
		return "[" + color.CyanString(string(Debug)) + "]"
	case Fatal:
		return "[" + color.RedString(string(Fatal)) + "]"
	case Panic:
		return "[" + color.RedString(string(Panic)) + "]"
	default:
		return "[" + color.GreenString(string(Info)) + "] "
	}
}

type FormatFunc func(levelStr string, str string) string

func DefaultFormatFunc(levelStr string, str string) string {
	return fmt.Sprintf("[%s] %s %s", time.Now().Format("2006-01-02 15:04:05 UTC-07"), LevelOutput(logLevel(levelStr)), str)
}

func DefaultColorFormatFunc(levelStr string, str string) string {
	return fmt.Sprintf("[%s] %s %s", time.Now().Format("2006-01-02 15:04:05 UTC-07"), LevelColorOutput(logLevel(levelStr)), str)
}
