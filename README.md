```
package main

import "github.com/redt1de/dbg"

type a struct {
	Name string
	blah int
}

func main() {

	dbg.Enable()
	dbg.Println("Hello, World!")
	dbg.Printf("Hello, %s!", "World")
	dbg.Dump(a{"test", 1})
}
```