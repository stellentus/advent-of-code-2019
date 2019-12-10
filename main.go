package main

import (
	"flag"
	"log"
)

var day = flag.Int("day", 1, "execute code for a given day")

type codeFunc func()

var codeForDay = make(map[int]codeFunc)

func main() {
	flag.Parse()
	cf, ok := codeForDay[*day]
	if !ok {
		log.Panicf("No code is available for day %d\n", *day)
	}
	cf()
}
