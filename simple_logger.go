package golog

import (
	"fmt"
	"io"
	"strings"
)

type SimpleLogger struct {
	debug      bool
	output     io.Writer
	errOutput  io.Writer
	formatFunc FormatFunc
}

func NewSimpleLogger() *SimpleLogger {
	return &SimpleLogger{
		output:     DefaultOutput,
		errOutput:  DefaultErrorOutput,
		formatFunc: DefaultFormatFunc,
	}
}

func (sl *SimpleLogger) SetFormatFunc(formatFunc FormatFunc) *SimpleLogger {
	if formatFunc != nil {
		sl.formatFunc = formatFunc
	}
	return sl
}

func (sl *SimpleLogger) SetOutput(output io.Writer) *SimpleLogger {
	sl.output = output
	return sl
}

func (sl *SimpleLogger) SetErrorOutput(errOutput io.Writer) *SimpleLogger {
	sl.errOutput = errOutput
	return sl
}

func (sl *SimpleLogger) SetDebug(debug bool) *SimpleLogger {
	sl.debug = debug
	return sl
}

func (sl *SimpleLogger) print(level logLevel, str string) {
	str = strings.TrimSpace(str)
	if str == "" {
		return
	}
	if sl.formatFunc != nil {
		str = sl.formatFunc(string(level), str)
	} else {
		str = DefaultFormatFunc(string(level), str)
	}
	str = strings.TrimSpace(str)
	fmt.Fprintln(sl.output, str)
}

func (sl *SimpleLogger) Info(a ...any) {
	sl.print(Info, fmt.Sprint(a...))
}

func (sl *SimpleLogger) Infof(format string, a ...any) {
	sl.print(Info, fmt.Sprintf(format, a...))
}

func (sl *SimpleLogger) Warn(a ...any) {
	sl.print(Warn, fmt.Sprint(a...))
}

func (sl *SimpleLogger) Warnf(format string, a ...any) {
	sl.print(Warn, fmt.Sprintf(format, a...))
}

func (sl *SimpleLogger) Error(a ...any) {
	sl.print(Error, fmt.Sprint(a...))
}

func (sl *SimpleLogger) Errorf(format string, a ...any) {
	sl.print(Error, fmt.Sprintf(format, a...))
}

func (sl *SimpleLogger) Debug(a ...any) {
	if sl.debug {
		sl.print(Debug, fmt.Sprint(a...))
	}
}

func (sl *SimpleLogger) Debugf(format string, a ...any) {
	if sl.debug {
		sl.print(Debug, fmt.Sprintf(format, a...))
	}
}

func (sl *SimpleLogger) Fatal(a ...any) {
	sl.print(Fatal, fmt.Sprint(a...))
}

func (sl *SimpleLogger) Fatalf(format string, a ...any) {
	sl.print(Fatal, fmt.Sprintf(format, a...))
}

func (sl *SimpleLogger) Panic(a ...any) {
	sl.print(Panic, fmt.Sprint(a...))
}

func (sl *SimpleLogger) Panicf(format string, a ...any) {
	sl.print(Panic, fmt.Sprintf(format, a...))
}

func (sl *SimpleLogger) LevelPrintFunc(level logLevel) func(a ...any) {
	switch level {
	case Info:
		return sl.Info
	case Warn:
		return sl.Warn
	case Error:
		return sl.Error
	case Debug:
		return sl.Debug
	case Fatal:
		return sl.Fatal
	case Panic:
		return sl.Panic
	default:
		return sl.Info
	}
}

func (sl *SimpleLogger) LevelPrintfFunc(level logLevel) func(format string, a ...any) {
	switch level {
	case Info:
		return sl.Infof
	case Warn:
		return sl.Warnf
	case Error:
		return sl.Errorf
	case Debug:
		return sl.Debugf
	case Fatal:
		return sl.Fatalf
	case Panic:
		return sl.Panicf
	default:
		return sl.Infof
	}
}

func (sl *SimpleLogger) LevelPrint(level logLevel, a ...any) {
	sl.LevelPrintFunc(level)(a...)
}

func (sl *SimpleLogger) LevelPrintf(level logLevel, format string, a ...any) {
	sl.LevelPrintfFunc(level)(format, a...)
}
