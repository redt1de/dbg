package dbg

import (
	"encoding/hex"
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/ztrue/tracerr"
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

type dbgLogger struct {
	name    string
	enabled bool
	verbose bool
}

type lm map[string]*dbgLogger

var loggerMap = make(lm)

// get a named logger or create a new one
func Get(name string) *dbgLogger {
	found, ok := loggerMap[name]
	if !ok {
		found = newD(name)
	}
	return found
}

func Set(namedlogger string, enabled, verbose bool) error {
	if _, ok := loggerMap[namedlogger]; !ok {
		return fmt.Errorf("logger %s not found", namedlogger)
	}
	loggerMap[namedlogger].enabled = enabled
	loggerMap[namedlogger].verbose = verbose
	return nil
}

// SetAll sets all loggers to the same state
func SetAll(enabled, verbose bool) {
	for k, v := range loggerMap {
		v.enabled = enabled
		v.verbose = verbose
		loggerMap[k] = v
	}
	dbgI.enabled = enabled
	dbgI.verbose = verbose
}

func newD(name string) *dbgLogger {
	n := &dbgLogger{name, false, false}
	loggerMap[name] = n
	return n
}

// Enable debug output
func (d *dbgLogger) Enable(v bool) {
	d.enabled = v
}

// Enable verbose debug output, includes file and line number
func (d *dbgLogger) Verbose(v bool) {
	d.verbose = v
}

func (d *dbgLogger) Printf(format string, args ...interface{}) {
	if d.enabled {
		var ver string
		if d.verbose {
			_, filename, line, _ := runtime.Caller(1)
			ver = fmt.Sprintf("[%s:%d] ", filename, line)
		}
		var modnme string
		if d.name != "" {
			modnme = fmt.Sprintf("[%s] ", strings.ToUpper(d.name))
		}
		fmt.Printf("%s[DEBUG] %s%s%s", dBlue, modnme, ver, dReset)
		fmt.Printf(format, args...)
	}
}

func (d *dbgLogger) Println(v ...any) {
	if d.enabled {
		var ver string
		if d.verbose {
			_, filename, line, _ := runtime.Caller(1)
			ver = fmt.Sprintf("[%s:%d] ", filename, line)
		}
		var modnme string
		if d.name != "" {
			modnme = fmt.Sprintf("[%s] ", strings.ToUpper(d.name))
		}
		fmt.Printf("%s[DEBUG] %s%s%s", dBlue, modnme, ver, dReset)

		fmt.Println(v...)

	}
}
func (d *dbgLogger) Errorf(format string, args ...interface{}) {
	var ver string
	if d.verbose {
		_, filename, line, _ := runtime.Caller(1)
		ver = fmt.Sprintf("[%s:%d] ", filename, line)
	}
	var modnme string
	if d.name != "" {
		modnme = fmt.Sprintf("[%s] ", strings.ToUpper(d.name))
	}
	fmt.Printf("%s[ERROR] %s%s%s", dRed, modnme, ver, dReset)
	fmt.Printf(format, args...)
}

func (d *dbgLogger) Errorln(v ...any) {
	var ver string
	if d.verbose {
		_, filename, line, _ := runtime.Caller(1)
		ver = fmt.Sprintf("[%s:%d] ", filename, line)
	}
	var modnme string
	if d.name != "" {
		modnme = fmt.Sprintf("[%s] ", strings.ToUpper(d.name))
	}
	fmt.Printf("%s[ERROR] %s%s%s", dRed, modnme, ver, dReset)

	fmt.Println(v...)
}

func (d *dbgLogger) Fatal(v ...any) {
	var ver string
	if d.verbose {
		_, filename, line, _ := runtime.Caller(1)
		ver = fmt.Sprintf("[%s:%d] ", filename, line)
	}
	var modnme string
	if d.name != "" {
		modnme = fmt.Sprintf("[%s] ", strings.ToUpper(d.name))
	}
	fmt.Printf("%s[FATAL] %s%s%s", dRed, modnme, ver, dReset)

	fmt.Println(v...)
	os.Exit(1)
}

func (d *dbgLogger) Dump(a interface{}) {
	if d.enabled {
		var modnme string
		if d.name != "" {
			modnme = fmt.Sprintf("[%s] ", strings.ToUpper(d.name))
		}
		fmt.Printf("%s[DUMP] %s", dYellow, modnme)
		_, filename, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d\n", filename, line)
		spew.Dump(a)
		fmt.Printf("%s", dReset)
	}
}

func (d *dbgLogger) TraceErr(err error) {
	var ver string
	if d.verbose {
		_, filename, line, _ := runtime.Caller(1)
		ver = fmt.Sprintf("[%s:%d] ", filename, line)
	}
	var modnme string
	if d.name != "" {
		modnme = fmt.Sprintf("[%s] ", strings.ToUpper(d.name))
	}
	fmt.Printf("%s[ERROR-TRACE] %s%s%s%s\n", dRed, modnme, ver, dReset, err.Error())

	///////////////////////////////////////////////
	err = tracerr.Wrap(err)
	a := tracerr.SprintSourceColor(err)

	lns := strings.Split(a, "\n")
	start := false // this is to skip TraceErr() itself, and stop at runtime.main() since we are not really interested in those
	for _, l := range lns {
		if strings.Contains(l, "runtime.main()") {
			fmt.Printf("%s", dReset)
			break
		}
		tmp, _ := hex.DecodeString("1b5b316d2f")
		if strings.HasPrefix(l, string(tmp)) && !strings.Contains(l, "TraceErr") {
			start = true
		}
		if start {
			fmt.Println(l)
		}
	}

}

func (d *dbgLogger) Trace() {
	nilError := tracerr.Errorf("%s", "")
	err := tracerr.Wrap(nilError)
	tracerr.PrintSourceColor(err)

}

// ///////////////////////////////////////////////////////// Global instance ///////////////////////////////////////////////////////////////////
func Printf(format string, args ...interface{}) {
	dbgI.Printf(format, args...)
}

func Println(v ...any) {
	dbgI.Println(v...)
}

func Errorf(format string, args ...interface{}) {
	dbgI.Errorf(format, args...)
}

func Errorln(v ...any) {
	dbgI.Errorln(v...)
}

func Fatal(v ...any) {
	dbgI.Fatal(v...)
}

func Dump(a interface{}) {
	dbgI.Dump(a)
}

func TraceErr(err error) {
	dbgI.TraceErr(err)
}

func Trace() {
	dbgI.Trace()
}

// Enable debug output
func Enable(v bool) {
	dbgI.enabled = v
}

// Enable verbose debug output, includes file and line number
func Verbose(v bool) {
	dbgI.verbose = v
}

// builtin global instance
var dbgI = newD("")
