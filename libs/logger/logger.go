package logger

import "fmt"

// Log Levels:
// -1, no logs
// 0, default: necessary logs
// 1, verbose
// 2, all info

// Logger is a wrapper for logging actions
type Logger interface {
	Println(a ...interface{})
	Printf(format string, a ...any)

	// UpdateStatus handles information relating to system configuration
	UpdateConfig(format string, a ...any)

	// UpdateStatus handles information relating to system status
	UpdateStatus(format string, a ...any)
}

type MainLogger struct {
	placeholder int
}

func NewLogger(level int) Logger {
	switch level {
	case 0:
		return &MainLogger{}
	default:
		return &MainLogger{}
	}
}

func (l *MainLogger) Println(a ...interface{}) {
	fmt.Println(a...)
}

// Printf adds a space to the format
func (l *MainLogger) Printf(format string, a ...any) {
	fmt.Printf(format+" ", a...)
}

func (l *MainLogger) UpdateConfig(format string, a ...any) {
	l.Printf(format, a...)
}

func (l *MainLogger) UpdateStatus(format string, a ...any) {
	l.Printf(format, a...)
}
