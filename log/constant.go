package log

import (
	"fmt"
	"io"
)

var __log *Logger = nil

// INIT initializes the global logger
func INIT(

	p_writer io.Writer,
	p_formatter Formatter,

) {
	__log = NewLogger(p_writer, p_formatter, 4)
}

// testLogger checks if the global logger is initialized.
// Returns true if initialized, false otherwise.
func testLogger() bool {
	if __log == nil {
		fmt.Println("The Logger is not initialized yet")
		fmt.Println("Please call the `INIT(p_log *Logger)` function first")
		return true
	}
	return false
}

func Info(__str string, args ...any) {
	if testLogger() {
		return
	}
	__log.Info(__str, args...)
}
func Warn(__str string, args ...any) {
	if testLogger() {
		return
	}
	__log.Warn(__str, args...)
}
func Error(__str string, args ...any) {
	if testLogger() {
		return
	}
	__log.Error(__str, args...)
}
func Fatal(__str string, args ...any) {
	if testLogger() {
		return
	}
	__log.Fatal(__str, args...)
}

func WithObj(__obj any) *LoggerWithFields {
	if testLogger() {
		return nil
	}
	return __log.WithObj(__obj)
}
func WithMap(obj map[string]any) *LoggerWithFields {
	if testLogger() {
		return nil
	}
	return __log.WithMap(obj)
}
func WithErr(__err error) *LoggerWithFields {
	if testLogger() {
		return nil
	}
	return __log.WithErr(__err)
}

// func InfoD(__str string, args ...any) { if testLogger() { return } ; __log.InfoDebug(__str, args...) }
// func WarnD(__str string, args ...any) { if testLogger() { return } ; __log.WarnD(__str, args...) }
// func ErrorD(__str string, args ...any) { if testLogger() { return } ; __log.ErrorD(__str, args...) }
// func FaralD(__str string, args ...any) { if testLogger() { return } ; __log.FatalD(__str, args...) }
