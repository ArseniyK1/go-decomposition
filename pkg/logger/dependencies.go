package logger

type LoggerResponse struct {
	Message string
}

type Logger interface {
	Debug(message string, fields ...any)
	Info(message string, fields ...any)
	Error(message string) error
}

type ConsoleLogger struct{}
