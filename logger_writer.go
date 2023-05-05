package golog

type LogWriter struct {
	logger Logger
	level  logLevel
}

func NewLogWriter(logger Logger, level logLevel) *LogWriter {
	return &LogWriter{
		logger: logger,
		level:  level,
	}
}

func (lw *LogWriter) Write(p []byte) (n int, err error) {
	lw.logger.LevelPrint(lw.level, string(p))
	return len(p), nil
}
