package dbg

import (
	"fmt"
	"os"
	"runtime"

	"github.com/davecgh/go-spew/spew"
)

var dReset = "\033[0m"
var dRed = "\033[31m"
var dGreen = "\033[32m"
var dYellow = "\033[33m"
var dBlue = "\033[34m"
var dPurple = "\033[35m"
var dCyan = "\033[36m"
var dGray = "\033[37m"
var dWhite = "\033[97m"

type debugLogger struct {
	Debug bool
	Trace bool
}

func newD() *debugLogger {
	return &debugLogger{false, false}
}

func Printf(format string, args ...interface{}) {
	if dbgI.Debug {

		fmt.Printf("%s[DEBUG] ", dBlue)
		if dbgI.Trace {
			_, filename, line, _ := runtime.Caller(1)
			fmt.Printf("%s:%d\n\t", filename, line)
		}

		fmt.Printf(format, args...)
		fmt.Printf("%s", dReset)
	}
}

func Println(v ...any) {
	if dbgI.Debug {
		fmt.Printf("%s[DEBUG] ", dBlue)
		if dbgI.Trace {
			_, filename, line, _ := runtime.Caller(1)
			fmt.Printf("%s:%d\n\t", filename, line)
		}

		fmt.Println(v...)
		fmt.Printf("%s", dReset)
	}
}

func Errorf(format string, args ...interface{}) {
	fmt.Printf("%s[ERROR] ", dRed)
	if dbgI.Trace {
		_, filename, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d\n\t", filename, line)
	}

	fmt.Printf(format, args...)
	fmt.Printf("%s", dReset)
}

func Errorln(v ...any) {
	fmt.Printf("%s[ERROR] ", dRed)
	if dbgI.Trace {
		_, filename, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d\n\t", filename, line)
	}

	fmt.Println(v...)
	fmt.Printf("%s", dReset)
}

func Fatal(v ...any) {
	fmt.Printf("%s[FATAL] ", dRed)
	if dbgI.Trace {
		_, filename, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d\n\t", filename, line)
	}

	fmt.Println(v...)
	fmt.Printf("%s", dReset)
	os.Exit(1)
}

func Dump(a interface{}) {
	if dbgI.Debug {
		fmt.Printf("%s[SPEW-DUMP] ", dYellow)
		_, filename, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d\n", filename, line)
		spew.Dump(a)
		fmt.Printf("%s", dReset)
	}
}

func Debug(v bool) {
	dbgI.Debug = v
}

func Trace(v bool) {
	dbgI.Trace = v
}

var dbgI = newD()
