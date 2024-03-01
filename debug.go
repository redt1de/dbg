package dbg

import (
	"fmt"
	"log"
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
	*log.Logger
	Enabled bool
}

func newD() *debugLogger {
	return &debugLogger{log.New(log.Writer(), "[debug] ", log.LstdFlags), false}
}
func Printf(format string, args ...interface{}) {
	if dbgI.Enabled {
		dbgI.Logger.SetPrefix("[debug] ")
		_, filename, line, _ := runtime.Caller(1)
		fmt.Printf("%s", dBlue)
		fm := fmt.Sprintf("%s:%d\n\t%s", filename, line, format)
		dbgI.Logger.Printf(fm, args...)
		fmt.Printf("%s", dReset)
	}
}

func Errorln(msg string) {
	if dbgI.Enabled {
		_, filename, line, _ := runtime.Caller(1)
		dbgI.Logger.SetPrefix("[ERROR] ")
		fm := fmt.Sprintf("%s:%d\n\t%s", filename, line, msg)
		fmt.Printf("%s", dRed)
		dbgI.Logger.Println(fm)
		fmt.Printf("%s", dReset)
	}
}

func Fatal(msg string) {
	if dbgI.Enabled {
		_, filename, line, _ := runtime.Caller(1)
		dbgI.Logger.SetPrefix("[FATAL] ")
		fm := fmt.Sprintf("%s:%d\n\t%s", filename, line, msg)
		fmt.Printf("%s", dRed)
		dbgI.Logger.Fatalln(fm)
		fmt.Printf("%s", dReset)
	}
}

func Println(msg string) {
	if dbgI.Enabled {
		_, filename, line, _ := runtime.Caller(1)
		dbgI.Logger.SetPrefix("[debug] ")
		fm := fmt.Sprintf("%s:%d\n\t%s", filename, line, msg)
		fmt.Printf("%s", dBlue)
		dbgI.Logger.Println(fm)
		fmt.Printf("%s", dReset)
	}
}

func Dump(a interface{}) {
	if dbgI.Enabled {
		_, filename, line, _ := runtime.Caller(1)
		dbgI.Logger.SetPrefix("[debug-dump] ")
		fmt.Printf("%s", dBlue)
		dbgI.Logger.Printf("%s:%d\n", filename, line)
		fmt.Printf("%s", dReset)
		spew.Dump(a)
	}
}

func Enable() {
	dbgI.Enabled = true
}

func Disable() {
	dbgI.Enabled = false
}

var dbgI = newD()
