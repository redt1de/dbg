package main

import (
	"github.com/redt1de/dbg"
)

type test struct {
	Name string
	Age  int
}

func main() {

	dbg.Enable(true)
	dbg.Printf("%s\n", "Printf")
	dbg.Println("Println")
	dbg.Errorf("%s\n", "Errorf")
	dbg.Errorln("Errorln")
	dbg.Dump(test{Name: "John", Age: 25})

	dbg.Verbose(true)
	dbg.Printf("%s\n", "Printf verbose")
	dbg.Println("Println verbose")
	dbg.Errorf("%s\n", "Errorf verbose")
	dbg.Errorln("Errorln verbose")
	// dbg.Fatal("Fatal Goodbye,Cruel World!")

	// named logger
	tl := dbg.Get("test")
	tl.Enable(true)
	tl.Verbose(true)
	tl.Printf("%s\n", "Printf")
	tl.Println("Println")
	tl.Errorf("%s\n", "Errorf")
	tl.Errorln("Errorln")
	tl.Dump(test{Name: "John", Age: 25})
	tl.Trace()
	tl.TraceErr(nil)
}
