package main

import (
	"flag"
	"fmt"

	"advent-of-code-2019/intcode"
)

var exampleFlag = flag.Int("example", 0, "use a different example dataset")
var debug = flag.Bool("debug", false, "print extra debug")

func main() {
	flag.Parse()

	program := []int64{} // This day's input

	switch *exampleFlag {
	case 1:
		program = []int64{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99} // copy of self
	case 2:
		program = []int64{1102, 34915192, 34915192, 7, 4, 7, 99, 0} // output 16-digit number
	}

	res := intcode.NewSimpleResult(program, []int64{})

	fmt.Println("D9-P1:", res)
}
