package main

import "fmt"

type Intcode struct {
	prog   map[int64]int64
	input  chan int64
	output chan int64
	done   chan bool
	label  int
	base   int64
}

func NewIC(program []int64, input, output chan int64, done chan bool, label int) Intcode {
	ic := Intcode{
		prog:   make(map[int64]int64),
		input:  input,
		output: output,
		done:   done,
		label:  label,
	}
	for i, val := range program {
		ic.prog[int64(i)] = val
	}
	return ic
}

func (ic *Intcode) calculate() {
	pc := int64(0)
	for pc != -1 {
		pc = ic.operate(pc)
	}
}

func (ic *Intcode) operate(pc int64) int64 {
	switch ic.getAt(pc) % 100 {
	case 1:
		ic.setMode(pc, 3, ic.getMode(pc, 1)+ic.getMode(pc, 2))
		return pc + 4
	case 2:
		ic.setMode(pc, 3, ic.getMode(pc, 1)*ic.getMode(pc, 2))
		return pc + 4
	case 3:
		ic.setMode(pc, 1, <-ic.input)
		return pc + 2
	case 4:
		ic.output <- ic.getMode(pc, 1)
		return pc + 2
	case 5:
		if ic.getMode(pc, 1) != 0 {
			return ic.getMode(pc, 2)
		} else {
			return pc + 3
		}
	case 6:
		if ic.getMode(pc, 1) == 0 {
			return ic.getMode(pc, 2)
		} else {
			return pc + 3
		}
	case 7:
		if ic.getMode(pc, 1) < ic.getMode(pc, 2) {
			ic.setMode(pc, 3, 1)
		} else {
			ic.setMode(pc, 3, 0)
		}
		return pc + 4
	case 8:
		if ic.getMode(pc, 1) == ic.getMode(pc, 2) {
			ic.setMode(pc, 3, 1)
		} else {
			ic.setMode(pc, 3, 0)
		}
		return pc + 4
	case 9:
		ic.base += ic.getMode(pc, 1)
		return pc + 2
	case 99:
		close(ic.input)
		ic.done <- true
		return -1
	default:
		fmt.Println("ERROR", pc)
		return -1
	}
}

func (ic Intcode) posForMode(pc, pos int64) int64 {
	code := ic.getAt(pc)
	pow := int64(100)
	for i := int64(1); i < pos; i++ {
		pow *= 10
	}
	code /= pow
	code = code % 10
	switch code {
	case 0:
		return ic.getAt(pc + pos)
	case 1:
		return pc + pos
	case 2:
		return ic.getAt(pc+pos) + ic.base
	default:
		fmt.Println("param mode NOT IMPLEMENTED", code)
		return 0
	}
}

func (ic Intcode) getMode(pc, pos int64) int64 {
	return ic.getAt(ic.posForMode(pc, pos))
}

func (ic Intcode) setMode(pc, pos, val int64) {
	ic.setAt(ic.posForMode(pc, pos), val)
}

func (ic Intcode) getAt(idx int64) int64 {
	val, ok := ic.prog[idx]
	if !ok {
		return 0
	}
	return val
}

func (ic Intcode) setAt(idx, val int64) {
	ic.prog[idx] = val
}
