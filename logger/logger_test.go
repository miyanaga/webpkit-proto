package logger

import (
	"io"
	"os"
	"testing"
)

func TestLogLevel(t *testing.T) {
	tc := []struct {
		str   string
		level int
	}{
		{"debug", DebugLevel},
		{"info", InfoLevel},
		{"warn", WarnLevel},
		{"error", ErrorLevel},
		{"fatal", FatalLevel},
	}

	for _, c := range tc {
		level := ParseLogLevel(c.str)
		if level != c.level {
			t.Errorf("ParseLogLevel() failed")
		}
	}
}

func captureStderr(cb func()) string {
	old := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w
	defer func() {
		os.Stderr = old
	}()

	cb()
	w.Close()

	out, _ := io.ReadAll(r)
	return string(out)
}

func TestLogger(t *testing.T) {
	var message string

	LogLevel = InfoLevel
	message = captureStderr(func() {
		Debug(func() string { return "debug-log" })
	})
	if message != "" {
		t.Errorf("Debug() failed")
	}

	LogLevel = DebugLevel
	message = captureStderr(func() {
		Debug(func() string { return "debug-log" })
	})
	if message != "debug-log\n" {
		t.Errorf("Debug() failed")
	}
}
