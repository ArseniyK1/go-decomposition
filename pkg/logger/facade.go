package logger

var globalLogger Logger = &ConsoleLogger{}

func SetLogger(l Logger) {
	globalLogger = l
}

func Debug(msg string, fields ...any) {
	globalLogger.Debug(msg, fields...)
}

func Info(msg string, fields ...any) {
	globalLogger.Info(msg, fields...)
}

func Error(msg string) error {
	return globalLogger.Error(msg)
}
