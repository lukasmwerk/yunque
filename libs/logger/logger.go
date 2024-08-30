package logger

import (
	"log"
)

// Using tabs to visually differentiate message categories
const (
	logInfoPrefix        = " (-) "
	logConfigPrefix      = " \t(#) "
	logStatePrefix       = " \t\t(@) "
	logLockMutexPrefix   = " \t\t\t([) "
	logUnlockMutexPrefix = " \t\t\t(]) "
	logErrorPrefix       = " (*) "
	logWarningPrefix     = " (!) "
	logFatalPrefix       = " (X) "
)

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

	// Lock Mutex is used to log a locking of a mutex
	LockMutex(format string, a ...any)
}

type MainLogger struct {
	logLevel int
}

func NewLogger(level int) Logger {
	switch level {
	case 0:
		return &MainLogger{logLevel: 0}
	default:
		return &MainLogger{}
	}
}

func (l *MainLogger) Println(a ...interface{}) {
	log.Println(a...)
}

// Printf adds a space to the format
func (l *MainLogger) Printf(format string, a ...any) {
	log.Printf(format+" ", a...)
}

func (l *MainLogger) UpdateConfig(format string, a ...any) {
	l.Printf(logConfigPrefix+format, a...)
}

func (l *MainLogger) UpdateStatus(format string, a ...any) {
	l.Printf(logStatePrefix+format, a...)
}

func (l *MainLogger) LockMutex(format string, a ...any) {
	l.Printf(logLockMutexPrefix+format, a...)
}

func (l *MainLogger) UnlockMutex(format string, a ...any) {
	l.Printf(logUnlockMutexPrefix+format, a...)
}
