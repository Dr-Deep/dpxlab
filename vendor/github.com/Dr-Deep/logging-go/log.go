package logging

import (
	"bytes"
	"fmt"
	"os"
	"runtime"
	"time"
)

// Log writes a message at the given log level with the given type and message parts.
func (logger *Logger) Log(level Level, logType string, messageParts ...string) {
	// check if we log on the right level
	if level < logger.Level {
		return
	}

	var rawBuf bytes.Buffer

	// assemble our log line
	fmt.Fprintf(
		&rawBuf,
		"[%s]-[%s]-[%s] ", time.Now().Format("2006-01-02T15:04:05.000Z"), logLevelDisplayNames[level], logType,
	)

	for i, p := range messageParts {
		rawBuf.WriteString(p)

		if i != len(messageParts)-1 {
			rawBuf.WriteString(" : ")
		}
	}

	rawBuf.WriteRune('\n')
	logger.Lock()

	// output line
	if _, err := logger.Writer.Write(rawBuf.Bytes()); err != nil {
		panic(err)
	}

	if err := logger.Writer.Flush(); err != nil {
		panic(err)
	}

	logger.Unlock()
}

// Debug logs a debug-level message including the caller function name.
func (logger *Logger) Debug(messageParts ...string) {
	pc, _, _, ok := runtime.Caller(1)
	funcname := runtime.FuncForPC(pc).Name()

	if !ok {
		funcname = unknownFuncName
	}

	logger.Log(LogDebug, funcname, messageParts...)
}

// Info logs an info-level message including the caller function name.
func (logger *Logger) Info(messageParts ...string) {
	pc, _, _, ok := runtime.Caller(1)
	funcname := runtime.FuncForPC(pc).Name()

	if !ok {
		funcname = unknownFuncName
	}

	logger.Log(LogInfo, funcname, messageParts...)
}

// Error logs an error-level message including the caller function name.
func (logger *Logger) Error(messageParts ...string) {
	pc, _, _, ok := runtime.Caller(1)
	funcname := runtime.FuncForPC(pc).Name()

	if !ok {
		funcname = unknownFuncName
	}

	logger.Log(LogError, funcname, messageParts...)
}

// Fatal logs a fatal error message including the caller function name and exits the program.
func (logger *Logger) Fatal(messageParts ...string) {
	pc, _, _, ok := runtime.Caller(1)
	funcname := runtime.FuncForPC(pc).Name()

	if !ok {
		funcname = unknownFuncName
	}

	logger.Log(LogFatal, funcname, messageParts...)
	os.Exit(1)
}
