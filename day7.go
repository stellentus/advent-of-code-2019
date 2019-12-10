package main

import "fmt"

func init() {
	// Originally implemented at https://play.golang.org/p/6p-1rDbUYAI
	codeForDay[7] = day7
}

type Intcode struct {
	prog   []int
	input  chan int
	output chan int
	done   chan bool
	label  int
}

func day7() {
	phase := []int{0, 1, 2, 3, 4}
	max, _ := Perm(phase, getThrust)
	fmt.Println("D7-P1:", max)

	phase = []int{5, 6, 7, 8, 9}
	max, _ = Perm(phase, getThrust)
	fmt.Println("D7-P2:", max)
}

func getThrust(phase []int) int {
	program := []int{3, 8, 1001, 8, 10, 8, 105, 1, 0, 0, 21, 38, 63, 72, 81, 106, 187, 268, 349, 430, 99999, 3, 9, 101, 5, 9, 9, 1002, 9, 3, 9, 101, 3, 9, 9, 4, 9, 99, 3, 9, 102, 3, 9, 9, 101, 4, 9, 9, 1002, 9, 2, 9, 1001, 9, 2, 9, 1002, 9, 4, 9, 4, 9, 99, 3, 9, 1001, 9, 3, 9, 4, 9, 99, 3, 9, 102, 5, 9, 9, 4, 9, 99, 3, 9, 102, 4, 9, 9, 1001, 9, 2, 9, 1002, 9, 5, 9, 1001, 9, 2, 9, 102, 3, 9, 9, 4, 9, 99, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 99, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 99, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 99, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 99, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 99}
	//program := []int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0}
	//program:=[]int{3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0}

	done := make([]chan bool, 5)
	io := make([]chan int, 6)
	io[0] = make(chan int)
	for i := 0; i < 5; i++ {
		io[i+1] = make(chan int)
		done[i] = make(chan bool)
		ic := NewIC(program, io[i], io[i+1], done[i], i)
		go func() {
			ic.calculate()
		}()
		io[i] <- phase[i]
	}
	io[0] <- 0 // initial input

	var res int
	isDone := false
	for !isDone {
		select {
		case isDone = <-done[0]:
			res = <-io[5]
		case val := <-io[5]:
			//fmt.Println("WAIT redir")
			//fmt.Println("SEND redir")
			io[0] <- val
			//fmt.Println("DONE redir")
		}
	}

	for i := 1; i < 4; i++ {
		<-done[i]
		//fmt.Println("DONE", i)
		close(done[i])
	}

	close(io[5])

	<-done[4]
	close(done[4])

	return res
}

func NewIC(program []int, input, output chan int, done chan bool, label int) Intcode {
	tmp := make([]int, len(program))
	copy(tmp, program)
	ic := Intcode{tmp, input, output, done, label}
	return ic
}

func (ic *Intcode) calculate() {
	pc := 0
	for pc < len(ic.prog) {
		pc = ic.operate(pc)
		if pc == -1 {
			break
		}
	}
}

func (ic *Intcode) operate(pc int) int {
	switch ic.prog[pc] % 100 {
	case 1:
		ic.prog[ic.prog[pc+3]] = ic.get(pc, 1) + ic.get(pc, 2)
		return pc + 4
	case 2:
		ic.prog[ic.prog[pc+3]] = ic.get(pc, 1) * ic.get(pc, 2)
		return pc + 4
	case 3:
		//fmt.Println("Input", ic.label)
		ic.prog[ic.prog[pc+1]] = <-ic.input
		return pc + 2
	case 4:
		//fmt.Println("Output", ic.label, ic.get(pc, 1))
		ic.output <- ic.get(pc, 1)
		return pc + 2
	case 5:
		if ic.get(pc, 1) != 0 {
			return ic.get(pc, 2)
		} else {
			return pc + 3
		}
	case 6:
		if ic.get(pc, 1) == 0 {
			return ic.get(pc, 2)
		} else {
			return pc + 3
		}
	case 7:
		if ic.get(pc, 1) < ic.get(pc, 2) {
			ic.prog[ic.prog[pc+3]] = 1
		} else {
			ic.prog[ic.prog[pc+3]] = 0
		}
		return pc + 4
	case 8:
		if ic.get(pc, 1) == ic.get(pc, 2) {
			ic.prog[ic.prog[pc+3]] = 1
		} else {
			ic.prog[ic.prog[pc+3]] = 0
		}
		return pc + 4
	case 99:
		close(ic.input)
		ic.done <- true
		return -1
	default:
		fmt.Println("ERROR", pc)
		return -1
	}
}

func (ic Intcode) get(pc, pos int) int {
	code := ic.prog[pc]
	pow := 100
	for i := 1; i < pos; i++ {
		pow *= 10
	}
	code /= pow
	code = code % 10
	if code == 0 {
		return ic.prog[ic.prog[pc+pos]]
	} else {
		return ic.prog[pc+pos]
	}
}

// Perm calls f with each permutation of a.
// modified from https://yourbasic.org/golang/generate-permutation-slice-string/
func Perm(a []int, f func([]int) int) (int, []int) {
	return perm(a, f, 0, 0, make([]int, len(a)))
}

// Permute the values at index i to len(a)-1.
func perm(a []int, f func([]int) int, i, best int, bestA []int) (int, []int) {
	if i > len(a) {
		res := f(a)
		//fmt.Println(a, res, best, bestA)
		if res > best {
			return res, a
		} else {
			return best, bestA
		}
	}
	res, resA := perm(a, f, i+1, best, bestA)
	if res > best {
		best = res
		copy(bestA, resA)
	}
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		res, resA = perm(a, f, i+1, best, bestA)
		if res > best {
			best = res
			copy(bestA, resA)
		}
		a[i], a[j] = a[j], a[i]
	}
	return best, resA
}
