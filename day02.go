package main

import "fmt"

func init() {
	// Originally implemented at https://play.golang.org/p/dmZ-90Ojebw
	codeForDay[2] = day2
}

func day2() {
	program := []int64{1, 0, 0, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 13, 1, 19, 1, 19, 6, 23, 1, 23, 6, 27, 1, 13, 27, 31, 2, 13, 31, 35, 1, 5, 35, 39, 2, 39, 13, 43, 1, 10, 43, 47, 2, 13, 47, 51, 1, 6, 51, 55, 2, 55, 13, 59, 1, 59, 10, 63, 1, 63, 10, 67, 2, 10, 67, 71, 1, 6, 71, 75, 1, 10, 75, 79, 1, 79, 9, 83, 2, 83, 6, 87, 2, 87, 9, 91, 1, 5, 91, 95, 1, 6, 95, 99, 1, 99, 9, 103, 2, 10, 103, 107, 1, 107, 6, 111, 2, 9, 111, 115, 1, 5, 115, 119, 1, 10, 119, 123, 1, 2, 123, 127, 1, 127, 6, 0, 99, 2, 14, 0, 0}
	//ic := Intcode{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50} // demo

	ic := NewSimpleIC(program, []int{})
	ic.set(1, 12)
	ic.set(2, 2)
	fmt.Println("D2-P1:", ic.calculate())

	for i := 0; i < 99; i++ {
		for j := 0; j < 99; j++ {
			ic := NewSimpleIC(program, []int{})
			ic.set(1, i)
			ic.set(2, j)
			if ic.calculate() == 19690720 {
				fmt.Println("D2-P2:", 100*i+j)
				break
			}
		}
	}
}

func (ic Intcode) set(pc, val int) {
	ic.prog[int64(pc)] = int64(val)
}
