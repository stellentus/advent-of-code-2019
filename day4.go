package main

import (
	"fmt"
)

func init() {
	// Originally implemented at https://play.golang.org/p/SAArjRUQjC-
	codeForDay[4] = day4
}

func day4() {
	min, max := 128392, 643281

	manySum := 0
	soloSum := 0

	for i := min; i <= max; i++ {
		many, solo := check(i)
		if many {
			manySum++
		}
		if solo {
			soloSum++
		}
	}
	fmt.Println("D4-P1:", manySum)
	fmt.Println("D4-P2:", soloSum)
}

func check(num int) (bool, bool) {
	lastDig := 10
	numDub := 0
	soloDub := 0
	seriesLen := 0
	for num > 0 {
		thisDig := num % 10
		num = num / 10
		if thisDig > lastDig {
			return false, false
		}
		thisDouble := thisDig == lastDig
		if thisDouble {
			numDub++
			seriesLen++
		} else {
			if seriesLen == 1 {
				soloDub++
			}
			seriesLen = 0
		}

		lastDig = thisDig
	}
	if seriesLen == 1 {
		soloDub++
	}

	return numDub > 0, soloDub > 0
}
