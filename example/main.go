package main

import (
	"os"

	"github.com/redt1de/dbg"
)

type test struct {
	Name string
	Age  int
}

var testlog = dbg.Get("test")

func main() {
	_, fakeErr := os.ReadFile("fakefile")
	dbg.Printf("%s\n", "global printf")
	dbg.Println("global println")
	dbg.Debugf("%s\n", "global debugf")
	dbg.Debugln("global  debugln")
	dbg.Warnf("%s\n", "global warnf")
	dbg.Warnln("global warnln")
	dbg.Errorf("%s,%s\n", "global errorf", fakeErr)
	dbg.Errorln(fakeErr)
	////////////////////////////////
	testlog.Printf("%s\n", "global printf")
	testlog.Println("global println")
	testlog.Debugf("%s\n", "global debugf")
	testlog.Debugln("global  debugln")
	testlog.Warnf("%s\n", "global warnf")
	testlog.Warnln("global warnln")
	testlog.Errorf("%s\n", "global errorf")
	testlog.Errorln(fakeErr)

	dbg.SetByName("test", true, dbg.LogError|dbg.LogWarn|dbg.LogSrc)

	dbg.Printf("%s\n", "global printf")
	dbg.Println("global println")
	dbg.Debugf("%s\n", "global debugf")
	dbg.Debugln("global  debugln")
	dbg.Warnf("%s\n", "global warnf")
	dbg.Warnln("global warnln")
	dbg.Errorf("%s\n", "global errorf")
	dbg.Errorln(fakeErr)
	////////////////////////////////

	testlog.Verbose(5)
	testlog.Printf("%s\n", "global printf")
	testlog.Println("global println")
	testlog.Debugf("%s\n", "global debugf")
	testlog.Debugln("global  debugln")
	testlog.Warnf("%s\n", "global warnf")
	testlog.Warnln("global warnln")
	testlog.Errorf("%s\n", "global errorf")
	testlog.Errorln(fakeErr)
}
