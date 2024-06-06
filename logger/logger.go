package logger

import (
	"os"
)

const (
	TraceLevel = 1 + iota
	DebugLevel
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
)

var (
	LogLevel = InfoLevel
)

func ParseLogLevel(level string) int {
	switch level {
	case "trace":
		return TraceLevel
	case "debug":
		return DebugLevel
	case "info":
		return InfoLevel
	case "warn":
		return WarnLevel
	case "error":
		return ErrorLevel
	case "fatal":
		return FatalLevel
	default:
		return InfoLevel
	}
}

func init() {
	level := os.Getenv("LOG_LEVEL")
	LogLevel = ParseLogLevel(level)
}

type LogMessageFunc func() string
type LoggingFunc func(level int, f LogMessageFunc)

func DefaultLogging(level int, f LogMessageFunc) {
	if level < LogLevel {
		return
	}

	message := f()
	os.Stderr.WriteString(message + "\n")
}

var (
	Logging = DefaultLogging
)

func Trace(f LogMessageFunc) {
	Logging(TraceLevel, f)
}

func Debug(f LogMessageFunc) {
	Logging(DebugLevel, f)
}

func Info(f LogMessageFunc) {
	Logging(InfoLevel, f)
}

func Warn(f LogMessageFunc) {
	Logging(WarnLevel, f)
}

func Error(f LogMessageFunc) {
	Logging(ErrorLevel, f)
}

func Fatal(f LogMessageFunc) {
	Logging(FatalLevel, f)
}
