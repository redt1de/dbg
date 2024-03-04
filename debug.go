package dbg

import (
	"encoding/hex"
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/ztrue/tracerr"
)

// TODO:
// use verbose for error stack trace

const (
	LogInfo = 1 << iota
	LogDebug
	LogWarn
	LogError
	LogSuccess
	LogDumps
	LogInfoSrc
	LogDebugSrc
	LogWarnSrc
	LogErrorSrc
	LogSuccessSrc
	LogDumpSrc
	LogTrace
	LogTraceVerbose
	LogErrTrace
)
const (
	LogAll     = LogInfo | LogWarn | LogError | LogDebug | LogTrace | LogDumps | LogErrTrace | LogInfoSrc | LogDebugSrc | LogWarnSrc | LogErrorSrc | LogDumpSrc | LogTraceVerbose | LogSuccess | LogSuccessSrc
	LogDefault = LogInfo | LogWarn | LogError | LogDebug | LogWarn | LogDumps | LogSuccess
	LogWithSrc = LogInfoSrc | LogDebugSrc | LogWarnSrc | LogErrorSrc | LogDumpSrc | LogSuccessSrc
)

var dReset = "\033[0m"
var dRed = "\033[31m"
var dGreen = "\033[32m"
var dOrange = "\033[38;5;208m"
var dYellow = "\033[33m"
var dBlue = "\033[34m"
var dPurple = "\033[35m"
var dCyan = "\033[36m"

// var dGray = "\033[37m"
var dGray = "\033[38;5;242m"

var dWhite = "\033[97m"

type dbgLogger struct {
	name    string
	enabled bool
	Flags   int
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

func SetByName(namedlogger string, enabled bool, flags int) error {
	if _, ok := loggerMap[namedlogger]; !ok {
		return fmt.Errorf("logger %s not found", namedlogger)
	}
	loggerMap[namedlogger].enabled = enabled
	loggerMap[namedlogger].Flags = flags
	return nil
}

// SetAll sets all loggers to the same state
func SetAll(enabled bool, flags int) {
	for k, v := range loggerMap {
		v.enabled = enabled
		v.Flags = flags
		loggerMap[k] = v
	}
	dbgI.enabled = enabled
	dbgI.Flags = flags
}

func newD(name string) *dbgLogger {
	n := &dbgLogger{name, true, LogDefault}
	loggerMap[name] = n
	return n
}

// Enable debug output
func (d *dbgLogger) Enabled(v bool) {
	d.enabled = v
}

// Enable verbose output
func (d *dbgLogger) MaxVerbose(v bool) {
	d.Flags = LogAll
}

func (d *dbgLogger) SetFlags(flags int) {
	d.Flags = flags
}

func (d *dbgLogger) Verbose(level int) {
	if level > 5 {
		level = 5
	}
	if level <= 0 {
		d.enabled = false
	}
	switch level {
	case 1:
		d.Flags = LogError | LogWarn | LogSuccess
	case 2:
		d.Flags = LogError | LogWarn | LogDebug | LogInfo | LogSuccess
	case 3:
		d.Flags = LogError | LogWarn | LogDebug | LogInfo | LogDumps | LogSuccess
	case 4:
		d.Flags = LogError | LogWarn | LogDebug | LogInfo | LogDumps | LogErrorSrc | LogWarnSrc | LogDebugSrc | LogInfoSrc | LogDumpSrc | LogSuccess | LogSuccessSrc
	case 5:
		d.Flags = LogAll
	}
}

func (d *dbgLogger) Printf(format string, args ...interface{}) {
	if d.enabled && d.Flags&LogInfo != 0 {
		var ver string
		if d.Flags&LogInfoSrc != 0 {
			_, filename, line, _ := runtime.Caller(1)
			ver = fmt.Sprintf("[%s:%d] ", filename, line)
		}
		var modnme string
		if d.name != "" {
			modnme = fmt.Sprintf("[%s] ", strings.ToUpper(d.name))
		}
		fmt.Printf("%s[INFO] %s%s%s", dWhite, modnme, ver, dReset)
		fmt.Printf(format, args...)
	}
}

func (d *dbgLogger) Println(v ...any) {
	if d.enabled && d.Flags&LogInfo != 0 {
		var ver string
		if d.Flags&LogInfoSrc != 0 {
			_, filename, line, _ := runtime.Caller(1)
			ver = fmt.Sprintf("[%s:%d] ", filename, line)
		}
		var modnme string
		if d.name != "" {
			modnme = fmt.Sprintf("[%s] ", strings.ToUpper(d.name))
		}
		fmt.Printf("%s[INFO] %s%s%s", dWhite, modnme, ver, dReset)

		fmt.Println(v...)

	}
}
func (d *dbgLogger) Successf(format string, args ...interface{}) {
	if d.enabled && d.Flags&LogSuccess != 0 {
		var ver string
		if d.Flags&LogSuccessSrc != 0 {
			_, filename, line, _ := runtime.Caller(1)
			ver = fmt.Sprintf("[%s:%d] ", filename, line)
		}
		var modnme string
		if d.name != "" {
			modnme = fmt.Sprintf("[%s] ", strings.ToUpper(d.name))
		}
		fmt.Printf("%s[SUCCESS] %s%s%s", dGreen, modnme, ver, dReset)
		fmt.Printf(format, args...)
	}
}

func (d *dbgLogger) Successln(v ...any) {
	if d.enabled && d.Flags&LogSuccess != 0 {
		var ver string
		if d.Flags&LogSuccessSrc != 0 {
			_, filename, line, _ := runtime.Caller(1)
			ver = fmt.Sprintf("[%s:%d] ", filename, line)
		}
		var modnme string
		if d.name != "" {
			modnme = fmt.Sprintf("[%s] ", strings.ToUpper(d.name))
		}
		fmt.Printf("%s[SUCCESS] %s%s%s", dGreen, modnme, ver, dReset)

		fmt.Println(v...)

	}
}

func (d *dbgLogger) Debugf(format string, args ...interface{}) {
	if d.enabled && d.Flags&LogDebug != 0 {
		var ver string
		if d.Flags&LogDebugSrc != 0 {
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

func (d *dbgLogger) Debugln(v ...any) {
	if d.enabled && d.Flags&LogDebug != 0 {
		var ver string
		if d.Flags&LogDebugSrc != 0 {
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

func (d *dbgLogger) Warnf(format string, args ...interface{}) {
	if d.enabled && d.Flags&LogWarn != 0 {
		var ver string
		if d.Flags&LogWarnSrc != 0 {
			_, filename, line, _ := runtime.Caller(1)
			ver = fmt.Sprintf("[%s:%d] ", filename, line)
		}
		var modnme string
		if d.name != "" {
			modnme = fmt.Sprintf("[%s] ", strings.ToUpper(d.name))
		}
		fmt.Printf("%s[WARN] %s%s%s", dOrange, modnme, ver, dReset)
		fmt.Printf(format, args...)
	}
}

func (d *dbgLogger) Warnln(v ...any) {
	if d.enabled && d.Flags&LogWarn != 0 {
		var ver string
		if d.Flags&LogWarnSrc != 0 {
			_, filename, line, _ := runtime.Caller(1)
			ver = fmt.Sprintf("[%s:%d] ", filename, line)
		}
		var modnme string
		if d.name != "" {
			modnme = fmt.Sprintf("[%s] ", strings.ToUpper(d.name))
		}
		fmt.Printf("%s[WARN] %s%s%s", dOrange, modnme, ver, dReset)

		fmt.Println(v...)

	}
}

func (d *dbgLogger) Errorf(format string, args ...interface{}) {
	if d.enabled && d.Flags&LogError != 0 {
		var ver string
		if d.Flags&LogErrorSrc != 0 {
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
}

func (d *dbgLogger) Errorln(err any) {
	if d.enabled && d.Flags&LogError != 0 {
		var ver string
		if d.Flags&LogErrorSrc != 0 {
			_, filename, line, _ := runtime.Caller(1)
			ver = fmt.Sprintf("[%s:%d] ", filename, line)
		}
		var modnme string
		if d.name != "" {
			modnme = fmt.Sprintf("[%s] ", strings.ToUpper(d.name))
		}
		fmt.Printf("%s[ERROR] %s%s%s", dRed, modnme, ver, dReset)

		fmt.Println(err)

	}
}

func (d *dbgLogger) Pause() {
	if d.enabled {
		var ver string
		_, filename, line, _ := runtime.Caller(1)
		ver = fmt.Sprintf("[%s:%d] ", filename, line)
		var modnme string
		if d.name != "" {
			modnme = fmt.Sprintf("[%s] ", strings.ToUpper(d.name))
		}
		fmt.Printf("%s[PAUSE] %s%s%s", dCyan, modnme, ver, dReset)
		fmt.Println("Press enter to continue...")
		var input string
		fmt.Scanln(&input)
	}
}

func (d *dbgLogger) Fatal(err error) {
	if d.enabled && d.Flags&LogError != 0 {
		var ver string
		if d.Flags&LogErrorSrc != 0 {
			_, filename, line, _ := runtime.Caller(1)
			ver = fmt.Sprintf("[%s:%d] ", filename, line)
		}
		var modnme string
		if d.name != "" {
			modnme = fmt.Sprintf("[%s] ", strings.ToUpper(d.name))
		}
		fmt.Printf("%s[Fatal] %s%s%s", dRed, modnme, ver, dReset)
		fmt.Println(err)
		d.TraceErr(err)
		os.Exit(1)

	}
}

func (d *dbgLogger) Dump(a ...interface{}) {
	if d.enabled && d.Flags&LogDumps != 0 {
		var note string
		if len(a) > 1 {
			switch a[0].(type) {
			case string:
				note = a[0].(string)
				a = a[1:]
			}
		}

		var ver string
		if d.Flags&LogDumpSrc != 0 {
			_, filename, line, _ := runtime.Caller(1)
			ver = fmt.Sprintf("[%s:%d] ", filename, line)
		}
		var modnme string
		if d.name != "" {
			modnme = fmt.Sprintf("[%s] ", strings.ToUpper(d.name))
		}
		fmt.Printf("%s[DUMP] %s%s%s\n", dYellow, modnme, ver, note)

		spew.Dump(a...)
		fmt.Printf("%s", dReset)
	}
}

func (d *dbgLogger) TraceErr(err error) {

	if d.Flags&LogErrTrace != 0 {
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

}

func (d *dbgLogger) Trace() {
	var ver string
	if d.enabled && d.Flags&LogTrace != 0 {
		_, filename, line, _ := runtime.Caller(1)
		ver = fmt.Sprintf("[%s:%d] ", filename, line)

		var modnme string
		if d.name != "" {
			modnme = fmt.Sprintf("[%s] ", strings.ToUpper(d.name))
		}
		fmt.Printf("%s[TRACE] %s%s%s\n", dYellow, modnme, ver, dReset)

		stackTrace := debug.Stack()
		s := string(stackTrace)
		o := strings.ReplaceAll(s, "\n\t", "%")
		tmp := strings.Split(o, "\n")
		indent := ""
		tmp2 := reverse(tmp)
		cnt := 0
		symb := ""
		for _, l := range tmp2 {
			if strings.HasPrefix(l, "goroutine") || strings.HasPrefix(l, "runtime/debug.Stack") || strings.HasPrefix(l, "github.com/redt1de/dbg") {
				continue
			}
			prts := strings.Split(l, "%")
			if len(prts) < 2 {
				continue
			}

			if cnt > 0 {
				symb = "└──> "
			}

			fn := prts[0]
			pth := prts[1]
			if d.Flags&LogTraceVerbose != 0 {
				fmt.Printf("%s%s%s%s%s  %s[%s]%s\n", dYellow, indent, symb, fn, dReset, dGray, pth, dReset)
			} else {
				fmt.Printf("%s%s%s%s%s\n", dYellow, indent, symb, fn, dReset)
				// if cnt > 0 {
				// 	fmt.Printf("%s%s%s%s%s\n", dGray, indent, "    ", pth, dReset)
				// } else {
				// 	fmt.Printf("%s%s%s%s%s\n", dGray, indent, "", pth, dReset)
				// }
			}

			indent += "    "
			cnt++
		}
	}

}
func reverse(s []string) []string {
	a := make([]string, len(s))
	copy(a, s)

	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}

	return a
}

// ///////////////////////////////////////////////////////// Global instance ///////////////////////////////////////////////////////////////////
func Printf(format string, args ...interface{}) {
	dbgI.Printf(format, args...)
}

func Println(v ...any) {
	dbgI.Println(v...)
}

func Debugf(format string, args ...interface{}) {
	dbgI.Debugf(format, args...)
}

func Debugln(v ...any) {
	dbgI.Debugln(v...)
}
func Warnf(format string, args ...interface{}) {
	dbgI.Warnf(format, args...)
}

func Warnln(v ...any) {
	dbgI.Warnln(v...)
}
func Errorf(format string, args ...interface{}) {
	dbgI.Errorf(format, args...)
}

func Errorln(err error) {
	dbgI.Errorln(err)
}

func Fatal(err error) {
	dbgI.Fatal(err)
}

func Dump(a ...interface{}) {
	dbgI.Dump(a...)
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
	dbgI.Flags = LogAll
}

// builtin global instance
var dbgI = newD("")
