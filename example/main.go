package main

import (
	"fmt"
	"os"

	"github.com/redt1de/dbg"
)

type test struct {
	Name string
	Age  int
}

var testlog = dbg.Get("test")
var fakeErr error

func main() {
	_, fakeErr = os.ReadFile("fakefile")
	dbg.Printf("%s\n", "global printf")
	dbg.Println("global println")
	dbg.Debugf("%s\n", "global debugf")
	dbg.Debugln("global  debugln")
	dbg.Warnf("%s\n", "global warnf")
	dbg.Warnln("global warnln")
	dbg.Errorf("%s,%s\n", "global errorf", fakeErr)
	dbg.Errorln(fakeErr)
	////////////////////////////////

	fmt.Println("--------------- verbose 1 -----------------")
	testlog.Verbose(1)
	testlog.Printf("%s\n", "global printf")
	testlog.Println("global println")
	testlog.Debugf("%s\n", "global debugf")
	testlog.Debugln("global  debugln")
	testlog.Warnf("%s\n", "global warnf")
	testlog.Warnln("global warnln")
	testlog.Errorf("%s\n", "global errorf")
	testlog.Errorln(fakeErr)
	testlog.Dump("test struct", test{"test", 1})

	fmt.Println("--------------- verbose 2 -----------------")
	testlog.Verbose(2)
	testlog.Printf("%s\n", "global printf")
	testlog.Println("global println")
	testlog.Debugf("%s\n", "global debugf")
	testlog.Debugln("global  debugln")
	testlog.Warnf("%s\n", "global warnf")
	testlog.Warnln("global warnln")
	testlog.Errorf("%s\n", "global errorf")
	testlog.Errorln(fakeErr)
	testlog.Dump("test struct", test{"test", 1})

	fmt.Println("--------------- verbose 3 -----------------")
	testlog.Verbose(3)
	testlog.Printf("%s\n", "global printf")
	testlog.Println("global println")
	testlog.Debugf("%s\n", "global debugf")
	testlog.Debugln("global  debugln")
	testlog.Warnf("%s\n", "global warnf")
	testlog.Warnln("global warnln")
	testlog.Errorf("%s\n", "global errorf")
	testlog.Errorln(fakeErr)
	testlog.Dump("test struct", test{"test", 1})

	fmt.Println("--------------- verbose 4 -----------------")
	testlog.Verbose(4)
	testlog.Printf("%s\n", "global printf")
	testlog.Println("global println")
	testlog.Debugf("%s\n", "global debugf")
	testlog.Debugln("global  debugln")
	testlog.Warnf("%s\n", "global warnf")
	testlog.Warnln("global warnln")
	testlog.Errorf("%s\n", "global errorf")
	testlog.Errorln(fakeErr)
	testlog.Dump("test struct", test{"test", 1})

	fmt.Println("--------------- verbose 5 -----------------")
	testlog.Verbose(5)
	testlog.Printf("%s\n", "global printf")
	testlog.Println("global println")
	testlog.Debugf("%s\n", "global debugf")
	testlog.Debugln("global  debugln")
	testlog.Warnf("%s\n", "global warnf")
	testlog.Warnln("global warnln")
	testlog.Errorf("%s\n", "global errorf")
	testlog.Errorln(fakeErr)
	testlog.Dump("test struct", test{"test", 1})
	testlog.Warnln("global warnln")
	fmt.Println("--------------- trace -----------------")
	FuncA()

}

func FuncA() error {
	testlog.Pause()
	a := FuncB()
	return a
}

func FuncB() error {
	b := FuncC()
	return b
}

func FuncC() error {
	testlog.Trace()

	// testlog.TraceErr(fakeErr)
	return fakeErr
}
