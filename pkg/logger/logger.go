package logger

import "fmt"

func (l *ConsoleLogger) Debug(message string, fields ...any) {
	fmt.Printf("[DEBUG] %s %v\n", message, fields)
}

func (l *ConsoleLogger) Info(message string, fields ...any) {
	fmt.Printf("INFO(message): %s\n", message)
	fmt.Printf("INFO(fields): %v\n", fields)
}

func (l *ConsoleLogger) Error(message string) error {
	return fmt.Errorf("ERROR(message): %s\n", message)
}
