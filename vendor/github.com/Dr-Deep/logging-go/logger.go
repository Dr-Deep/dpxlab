// Package logging provides a simple thread-safe logging facility.
package logging

import (
	"bufio"
	"os"
	"sync"
)

// Level represents the severity of a log message.
type Level int

const (
	defaultLogLevel = LogDebug

	// LogDebug represents debug messages.
	LogDebug Level = iota

	// LogInfo represents informational messages.
	LogInfo

	// LogError represents errors.
	LogError

	// LogFatal represents fatal errors.
	LogFatal

	unknownFuncName = "???"
)

var (
	// logLevelDisplayNames gives the display name to use for our log levels.
	logLevelDisplayNames = map[Level]string{
		LogDebug: "debug",
		LogInfo:  "info",
		LogError: "error",
		LogFatal: "fatal",
	}
)

// Logger represents a thread-safe logger with level filtering and buffered output.
type Logger struct {
	sync.Mutex
	File   *os.File
	Writer *bufio.Writer
	Level  Level
}

// NewLogger creates a new Logger instance writing to the given file (or stdout if nil).
func NewLogger(file *os.File) *Logger {
	if file == nil {
		file = os.Stdout
	}

	return &Logger{
		File:   file,
		Level:  defaultLogLevel,
		Writer: bufio.NewWriter(file),
	}
}

// Close flushes buffered log output and closes the underlying file if present.
func (logger *Logger) Close() error {
	if err := logger.Writer.Flush(); err != nil {
		return &FlushError{Err: err}
	}

	if err := logger.File.Close(); err != nil {
		return &CloseError{Err: err}
	}

	return nil
}
