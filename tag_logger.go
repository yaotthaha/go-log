package golog

type TagLogger struct {
	tag    string
	logger Logger
}

func NewTagLogger(logger Logger, tag string) *TagLogger {
	return &TagLogger{
		tag:    tag,
		logger: logger,
	}
}

func (tl *TagLogger) Info(a ...any) {
	tl.logger.Info(append([]any{"[", tl.tag, "] "}, a...)...)
}

func (tl *TagLogger) Infof(format string, a ...any) {
	tl.logger.Infof("[%s] "+format, append([]any{tl.tag}, a...)...)
}

func (tl *TagLogger) Warn(a ...any) {
	tl.logger.Warn(append([]any{"[", tl.tag, "] "}, a...)...)
}

func (tl *TagLogger) Warnf(format string, a ...any) {
	tl.logger.Warnf("[%s] "+format, append([]any{tl.tag}, a...)...)
}

func (tl *TagLogger) Error(a ...any) {
	tl.logger.Error(append([]any{"[", tl.tag, "] "}, a...)...)
}

func (tl *TagLogger) Errorf(format string, a ...any) {
	tl.logger.Errorf("[%s] "+format, append([]any{tl.tag}, a...)...)
}

func (tl *TagLogger) Debug(a ...any) {
	tl.logger.Debug(append([]any{"[", tl.tag, "] "}, a...)...)
}

func (tl *TagLogger) Debugf(format string, a ...any) {
	tl.logger.Debugf("[%s] "+format, append([]any{tl.tag}, a...)...)
}

func (tl *TagLogger) Fatal(a ...any) {
	tl.logger.Fatal(append([]any{"[", tl.tag, "] "}, a...)...)
}

func (tl *TagLogger) Fatalf(format string, a ...any) {
	tl.logger.Fatalf("[%s] "+format, append([]any{tl.tag}, a...)...)
}

func (tl *TagLogger) Panic(a ...any) {
	tl.logger.Panic(append([]any{"[", tl.tag, "] "}, a...)...)
}

func (tl *TagLogger) Panicf(format string, a ...any) {
	tl.logger.Panicf("[%s] "+format, append([]any{tl.tag}, a...)...)
}

func (tl *TagLogger) LevelPrintFunc(level logLevel) func(a ...any) {
	switch level {
	case Info:
		return tl.Info
	case Warn:
		return tl.Warn
	case Error:
		return tl.Error
	case Debug:
		return tl.Debug
	case Fatal:
		return tl.Fatal
	case Panic:
		return tl.Panic
	default:
		return tl.Info
	}
}

func (tl *TagLogger) LevelPrintfFunc(level logLevel) func(format string, a ...any) {
	switch level {
	case Info:
		return tl.Infof
	case Warn:
		return tl.Warnf
	case Error:
		return tl.Errorf
	case Debug:
		return tl.Debugf
	case Fatal:
		return tl.Fatalf
	case Panic:
		return tl.Panicf
	default:
		return tl.Infof
	}
}

func (tl *TagLogger) LevelPrint(level logLevel, a ...any) {
	tl.LevelPrintFunc(level)(a...)
}

func (tl *TagLogger) LevelPrintf(level logLevel, format string, a ...any) {
	tl.LevelPrintfFunc(level)(format, a...)
}
