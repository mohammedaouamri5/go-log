package log

import (
	"fmt"
	"os"
)

var __log *Logger

// INIT initializes the global logger
func INIT(p_log *Logger) {
	if p_log == nil {
		fmt.Fprintln(os.Stderr, "INIT: cannot initialize with a nil logger")
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
		return false
	}
	return true
}





func  INFO(__str string, args ...any) {
	if testLogger() {return}
	__log.Info(__str,args...)
}

func  INFODEBUG(__str string, args ...any) {
	if testLogger() {return}
	__log.InfoDebug(__str,args...)
}

func  WARN(__str string, args ...any) {
	if testLogger() {return}
	__log.Warn(__str,args...)
}

func  WARND(__str string, args ...any) {
	if testLogger() {return}
	__log.WarnD(__str,args...)
}

func  ERROR(__str string, args ...any) {
	if testLogger() {return}
	__log.Error(__str,args...)
}

func  ERRORD(__str string, args ...any) {
	if testLogger() {return}
	__log.ErrorD(__str,args...)
}

func  FATAL(__str string, args ...any) {
	if testLogger() {return}
	__log.Fatal(__str,args...)
}

func  FATALD(__str string, args ...any) {
	if testLogger() {return}
	__log.FatalD(__str,args...)
}







