package log

import (
	"fmt"
)

var __log *Logger = nil

// INIT initializes the global logger
func INIT(p_log *Logger) {
	if p_log == nil {
		fmt.Println("INIT: cannot initialize with a nil logger")
		return
	}
	__log = p_log
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

// func InfoD(__str string, args ...any) { if testLogger() { return } ; __log.InfoDebug(__str, args...) }
// func WarnD(__str string, args ...any) { if testLogger() { return } ; __log.WarnD(__str, args...) }
// func ErrorD(__str string, args ...any) { if testLogger() { return } ; __log.ErrorD(__str, args...) }
// func FaralD(__str string, args ...any) { if testLogger() { return } ; __log.FatalD(__str, args...) }
