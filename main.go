package main

import (
	"flag"
	"log"
)

var dayFlag = flag.Int("day", 1, "execute code for a given day")
var exampleFlag = flag.Int("example", 0, "use a different example dataset")

type codeFunc func(example int)

var codeForDay = make(map[int]codeFunc)

func main() {
	flag.Parse()
	cf, ok := codeForDay[*dayFlag]
	if !ok {
		log.Panicf("No code is available for day %d\n", *dayFlag)
	}
	cf(*exampleFlag)
}
