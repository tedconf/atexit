// Test program for atexit, gets output file and data as arguments and writes
// data to output file in atexit handler.
package main

import (
	atexit "./_obj/atexit"
	"flag"
	"fmt"
	"io/ioutil"
)

var outfile = ""
var data = ""

func handler() {
	ioutil.WriteFile(outfile, []byte(data), 0666)
}

func badHandler() {
	n := 0
	fmt.Println(1/n)
}

func main() {
	flag.Parse()
	outfile = flag.Arg(0)
	data = flag.Arg(1)

	atexit.Register(handler)
	atexit.Register(badHandler)
	atexit.Exit(1)
}
