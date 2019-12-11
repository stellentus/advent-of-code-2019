package main

import "fmt"

type Intcode struct {
	prog         map[int64]int64
	done         chan bool
	label        int
	base         int64
	requestInput *chan bool
	InputProvider
	OutputProvider
}

type InputProvider func() int64
type OutputProvider func(int64)

type sliceInputter struct {
	slice []int64
	idx   int
}

func (si *sliceInputter) get() int64 {
	val := si.slice[si.idx]
	si.idx++
	return val
}

func (ic *Intcode) SendInputSlice(in []int64) {
	si := sliceInputter{slice: in}
	ic.InputProvider = si.get
}

func (ic *Intcode) ExpectOutputArray() *[]int64 {
	ss := &[]int64{}
	ic.OutputProvider = func(i int64) {
		*ss = append(*ss, i)
	}
	return ss
}

func (ic *Intcode) SetInputChan(input chan int64) {
	ic.InputProvider = func() int64 {
		return <-input
	}
}

func (ic *Intcode) SetOutputChan(output chan int64) {
	ic.OutputProvider = func(i int64) {
		output <- i
	}
}

func NewIC(program []int64) Intcode {
	ic := Intcode{
		prog: make(map[int64]int64),
	}
	for i, val := range program {
		ic.prog[int64(i)] = val
	}
	return ic
}

func NewSimpleICResult(program, input []int64) []int64 {
	ic := NewIC(program)
	ic.SendInputSlice(input)
	staticOutput := ic.ExpectOutputArray()
	ic.calculate()
	return *staticOutput
}

func (ic *Intcode) SetLabel(label int) {
	ic.label = label
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
		if ic.requestInput != nil {
			*ic.requestInput <- true
		}
		ic.setMode(pc, 1, ic.InputProvider())
		return pc + 2
	case 4:
		ic.OutputProvider(ic.getMode(pc, 1))
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
		// close(ic.input)
		// ic.done <- true
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
