package main

import (
	"fmt"

	"advent-of-code-2019/intcode"
)

func init() {
	// Originally implemented at https://play.golang.org/p/6p-1rDbUYAI
	codeForDay[7] = day7
}

func day7(example int) {
	phase := []int{0, 1, 2, 3, 4}
	max, _ := Perm(example, phase, getThrust)
	fmt.Println("D7-P1:", max)

	phase = []int{5, 6, 7, 8, 9}
	max, _ = Perm(example, phase, getThrust)
	fmt.Println("D7-P2:", max)
}

func getThrust(ex int, phase []int) int {
	program := []int64{3, 8, 1001, 8, 10, 8, 105, 1, 0, 0, 21, 38, 63, 72, 81, 106, 187, 268, 349, 430, 99999, 3, 9, 101, 5, 9, 9, 1002, 9, 3, 9, 101, 3, 9, 9, 4, 9, 99, 3, 9, 102, 3, 9, 9, 101, 4, 9, 9, 1002, 9, 2, 9, 1001, 9, 2, 9, 1002, 9, 4, 9, 4, 9, 99, 3, 9, 1001, 9, 3, 9, 4, 9, 99, 3, 9, 102, 5, 9, 9, 4, 9, 99, 3, 9, 102, 4, 9, 9, 1001, 9, 2, 9, 1002, 9, 5, 9, 1001, 9, 2, 9, 102, 3, 9, 9, 4, 9, 99, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 99, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 99, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 99, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 99, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 99}

	switch ex {
	case 1:
		program = []int64{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0}
	case 2:
		program = []int64{3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23, 101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0}
	}

	// Make all channels and load the initial phases
	io := make([]chan int64, 5)
	for i := 0; i < 5; i++ {
		io[i] = make(chan int64, 1)
		io[i] <- int64(phase[i]) // won't block because channel has buffer of 1
	}

	// Make the first 4 channels
	for i := 0; i < 4; i++ {
		ic := newChanIC(program, io[i], io[i+1], i)
		go func(i int) {
			// this go func will naturally end when this Intcode is done
			ic.Calculate()
		}(i)
	}

	// Make the last channel and don't launch a thread
	ic := newChanIC(program, io[4], io[0], 4)
	io[0] <- 0     // initial input
	ic.Calculate() // Once this ends, the buffered output has one last read

	res := int(<-io[0])

	for i := 0; i < 5; i++ {
		close(io[i])
	}

	return res
}

func newChanIC(prog []int64, input, output chan int64, label int) intcode.Intcode {
	ic := intcode.New(prog)
	ic.SetInputChan(input)
	ic.SetOutputChan(output)
	ic.SetLabel(label)
	return ic
}

// Perm calls f with each permutation of a.
// modified from https://yourbasic.org/golang/generate-permutation-slice-string/
func Perm(ex int, a []int, f func(int, []int) int) (int, []int) {
	return perm(ex, a, f, 0, 0, make([]int, len(a)))
}

// Permute the values at index i to len(a)-1.
func perm(ex int, a []int, f func(int, []int) int, i, best int, bestA []int) (int, []int) {
	if i > len(a) {
		res := f(ex, a)
		if res > best {
			return res, a
		} else {
			return best, bestA
		}
	}
	res, resA := perm(ex, a, f, i+1, best, bestA)
	if res > best {
		best = res
		copy(bestA, resA)
	}
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		res, resA = perm(ex, a, f, i+1, best, bestA)
		if res > best {
			best = res
			copy(bestA, resA)
		}
		a[i], a[j] = a[j], a[i]
	}
	return best, resA
}
