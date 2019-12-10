package main

import "fmt"

func init() {
	codeForDay[10] = func() {
		day10_1()
		day10_2()
	}
}

func day10_1() {
	// locations := [][]bool{
	// 	[]bool{false, true, false, false, true},
	// 	[]bool{false, false, false, false, false},
	// 	[]bool{true, true, true, true, true},
	// 	[]bool{false, false, false, false, true},
	// 	[]bool{false, false, false, true, true},
	// }
	// locations := [][]bool{
	// 	[]bool{false, false, false, false, false, false, true, false, true, false},
	// 	[]bool{true, false, false, true, false, true, false, false, false, false},
	// 	[]bool{false, false, true, true, true, true, true, true, true, false},
	// 	[]bool{false, true, false, true, false, true, true, true, false, false},
	// 	[]bool{false, true, false, false, true, false, false, false, false, false},
	// 	[]bool{false, false, true, false, false, false, false, true, false, true},
	// 	[]bool{true, false, false, true, false, false, false, false, true, false},
	// 	[]bool{false, true, true, false, true, false, false, true, true, true},
	// 	[]bool{true, true, false, false, false, true, false, false, true, false},
	// 	[]bool{false, true, false, false, false, false, true, true, true, true},
	// }
	locations := [][]bool{
		[]bool{false, false, false, false, true, false, false, false, true, true, true, true, false, true, false, true, false, false, false, false, false, false, false, false, false, false, false, true, false, false, false, false, false, false, false, false},
		[]bool{true, true, true, true, true, false, false, true, false, true, false, true, false, false, false, false, false, false, true, true, true, true, true, false, false, false, true, false, true, false, false, false, true, false, false, false},
		[]bool{true, true, false, true, true, false, false, true, false, true, false, true, false, false, false, false, false, true, false, false, false, false, false, true, true, false, true, false, true, false, false, true, false, false, false, false},
		[]bool{false, false, false, true, false, false, true, false, false, false, true, false, true, true, false, false, false, false, false, false, false, false, true, false, false, true, false, false, false, false, false, false, false, true, false, true},
		[]bool{true, false, false, false, true, true, false, false, false, true, true, true, false, false, false, true, true, true, false, false, true, false, false, false, true, false, false, false, false, false, true, false, false, false, false, false},
		[]bool{true, true, false, false, false, false, false, false, false, true, false, false, false, false, false, true, false, false, false, false, false, false, false, false, false, true, false, true, false, false, false, false, true, false, true, false},
		[]bool{false, false, true, false, false, false, true, false, true, true, false, true, true, false, false, false, false, false, true, false, false, false, false, true, true, false, false, true, false, false, false, false, false, false, true, false},
		[]bool{false, false, true, true, true, false, false, true, true, false, false, true, false, false, true, false, false, false, true, false, false, false, false, false, false, true, true, false, false, false, true, false, false, false, false, true},
		[]bool{true, true, false, false, true, true, false, false, false, false, false, true, false, false, false, true, false, true, false, false, false, true, false, false, false, false, false, false, true, false, true, false, true, false, false, true},
		[]bool{false, false, false, true, true, true, false, false, false, false, true, false, false, true, false, true, false, false, false, false, false, false, true, false, false, false, true, false, false, false, false, false, false, false, true, false},
		[]bool{true, false, false, false, false, true, false, false, false, true, true, false, false, false, false, false, false, false, true, false, false, true, false, false, false, false, false, false, false, true, false, false, true, false, false, false},
		[]bool{true, false, false, false, false, false, false, false, false, false, false, false, true, false, false, false, false, false, true, false, false, false, false, false, true, false, true, false, false, false, true, false, true, true, false, true},
		[]bool{true, true, true, false, false, true, false, false, false, false, true, true, true, true, false, false, true, false, true, true, true, false, false, false, true, false, false, false, false, true, false, false, true, false, false, false},
		[]bool{true, true, false, false, false, false, true, false, true, false, false, true, false, true, false, false, false, false, false, false, true, true, false, false, false, false, false, false, false, true, false, false, false, false, true, false},
		[]bool{false, false, true, false, true, false, false, false, false, true, false, true, false, true, false, false, true, false, false, false, true, false, true, true, false, true, true, false, false, true, false, false, false, false, false, false},
		[]bool{false, false, false, true, false, false, false, false, false, true, false, false, false, false, false, false, true, false, true, false, true, false, true, true, false, false, false, false, false, true, false, false, true, true, true, false},
		[]bool{false, false, true, false, true, false, true, true, true, false, false, false, false, false, false, false, true, false, false, true, false, true, false, false, false, false, true, true, false, false, false, false, false, true, false, false},
		[]bool{false, true, false, true, false, true, false, false, false, true, false, false, true, false, true, false, false, true, true, false, true, false, false, false, false, false, false, false, false, false, false, true, false, false, false, true},
		[]bool{false, false, false, false, false, true, false, true, false, true, false, false, false, true, false, false, true, false, false, true, false, false, false, true, true, true, false, true, false, false, false, true, false, true, false, false},
		[]bool{true, false, false, true, false, false, true, false, false, false, false, false, true, false, true, true, false, false, true, true, false, false, false, true, true, false, true, false, false, false, false, false, true, false, false, false},
		[]bool{false, false, false, false, true, true, false, false, false, false, true, false, true, true, false, false, false, true, false, false, false, false, false, false, false, false, false, false, true, false, true, true, false, false, false, false},
		[]bool{false, false, false, true, false, false, false, false, true, true, true, false, true, false, false, false, true, true, false, false, false, false, false, false, false, false, true, true, false, true, true, false, false, true, true, false},
		[]bool{true, false, false, true, false, false, false, false, true, false, false, false, false, false, false, true, false, false, false, false, false, false, true, true, true, false, false, false, false, false, false, false, false, false, false, false},
		[]bool{true, true, false, false, false, true, false, false, true, false, true, true, false, true, true, false, false, true, true, false, false, false, false, true, false, false, true, false, false, true, true, false, false, true, false, true},
		[]bool{false, true, false, false, false, false, true, false, false, true, true, false, false, false, false, false, true, false, true, false, false, false, false, false, false, false, false, false, false, false, false, true, true, false, false, false},
		[]bool{false, true, true, true, false, false, false, false, false, false, false, false, false, true, false, false, false, false, true, false, true, true, false, true, false, false, true, false, true, false, false, true, false, true, false, false},
		[]bool{true, false, false, false, true, false, false, true, false, false, false, true, false, true, false, true, false, false, false, false, false, true, false, false, false, false, true, false, false, false, false, false, false, true, true, true},
		[]bool{true, false, false, false, false, false, false, false, false, false, false, false, true, true, false, true, false, false, false, false, true, false, true, true, false, false, false, false, false, false, true, false, true, false, false, true},
		[]bool{false, false, false, false, true, false, false, false, true, false, false, true, false, false, false, true, false, true, true, true, true, false, false, false, true, false, true, false, false, true, false, true, true, false, false, false},
		[]bool{false, false, false, false, false, false, true, true, true, true, false, false, false, false, false, true, false, false, true, false, false, false, false, true, false, false, false, false, true, false, false, false, false, true, false, true},
		[]bool{false, true, true, false, true, false, false, true, true, true, false, false, true, true, true, true, false, false, false, true, false, false, false, false, false, false, false, true, false, true, false, false, false, false, true, false},
		[]bool{true, false, true, true, true, false, false, false, false, true, false, false, false, false, true, false, false, false, false, false, false, false, false, false, false, true, false, false, false, false, false, true, true, true, false, true},
		[]bool{false, false, false, true, false, false, false, false, false, false, true, false, false, false, false, true, true, false, false, false, true, true, false, false, true, false, false, true, false, false, false, true, true, true, false, false},
		[]bool{false, false, true, false, false, false, true, true, true, false, true, true, true, false, false, false, false, false, false, false, false, false, true, false, true, false, false, true, false, true, false, false, true, false, false, false},
		[]bool{false, true, false, true, false, false, false, false, false, false, false, false, false, false, false, false, false, true, false, true, false, false, false, false, true, false, false, false, false, false, false, false, false, false, false, false},
		[]bool{false, false, true, false, false, false, true, false, true, true, true, false, false, false, true, true, false, false, false, false, true, true, false, true, false, true, false, true, false, false, false, false, true, false, true, false},
	}

	maxsighted := 0
	maxi := -1
	maxj := -1

	for thisi, row := range locations {
		for thisj, val := range row {
			if !val {
				continue
			}
			sighted := numSeen(locations, thisi, thisj)
			if sighted > maxsighted {
				maxsighted = sighted
				maxi = thisi
				maxj = thisj
			}
		}
	}

	fmt.Println("D10-P1:", maxsighted, maxi, maxj)
}

func numSeen(loc [][]bool, thisi, thisj int) int {
	maxi := len(loc)
	hidden := make([][]bool, maxi)
	maxj := len(loc[0])
	for i := range loc {
		hidden[i] = make([]bool, maxj)
	}

	seen := 0
	for i := thisi; i < maxi; i++ {
		for j := thisj; j < maxj; j++ {
			if !loc[i][j] || hidden[i][j] || (i == thisi && j == thisj) {
				continue
			}
			// fmt.Println("\tHiding at", i, j, seen)
			seen++
			hide(hidden, thisi, thisj, i, j, maxi, maxj)
		}
	}
	for i := thisi; i < maxi; i++ {
		for j := thisj - 1; j >= 0; j-- {
			if !loc[i][j] || hidden[i][j] || (i == thisi && j == thisj) {
				continue
			}
			// fmt.Println("\tHiding at", i, j, seen)
			seen++
			hide(hidden, thisi, thisj, i, j, maxi, maxj)
		}
	}
	for i := thisi - 1; i >= 0; i-- {
		for j := thisj - 1; j >= 0; j-- {
			if !loc[i][j] || hidden[i][j] || (i == thisi && j == thisj) {
				continue
			}
			// fmt.Println("\tHiding at", i, j, seen)
			seen++
			hide(hidden, thisi, thisj, i, j, maxi, maxj)
		}
	}
	for i := thisi - 1; i >= 0; i-- {
		for j := thisj; j < maxj; j++ {
			if !loc[i][j] || hidden[i][j] || (i == thisi && j == thisj) {
				continue
			}
			// fmt.Println("\tHiding at", i, j)
			seen++
			hide(hidden, thisi, thisj, i, j, maxi, maxj)
		}
	}
	// fmt.Printf("Visible at %d,%d (%d)\n\t", thisi, thisj, seen)
	// seen = 0
	for i, row := range loc {
		for j, val := range row {
			if i == thisi && j == thisj {
				// fmt.Printf("%c", 'O')
			} else if !val {
				// fmt.Printf("%c", '.')
			} else if hidden[i][j] {
				// fmt.Printf("%c", '-')
			} else {
				// fmt.Printf("%c", '#')
				// seen++
			}
		}
		// fmt.Printf("\n\t")
	}
	// fmt.Println("or just ", seen)
	return seen
}

func hide(hidden [][]bool, thisi, thisj, i, j, maxi, maxj int) {
	diffi, diffj := minDiff(i-thisi, j-thisj)
	// fmt.Printf("%d,%d to %d,%d scaled by %d,%d (diff %d,%d)\n", thisi, thisj, i, j, diffi, diffj, i-thisi, j-thisj)
	i += diffi
	j += diffj
	for i >= 0 && i < maxi && j >= 0 && j < maxj {
		// fmt.Println("\t\tHiding", i, j)
		if thisi != i || thisj != j {
			hidden[i][j] = true
			// fmt.Printf("\thiding%d,%d\n", i, j)
		}
		i += diffi
		j += diffj
	}
}

func minDiff(diffi, diffj int) (int, int) {
	gc := gcd(diffi, diffj)
	if gc < 0 {
		gc *= -1
	}
	return diffi / gc, diffj / gc
}

func gcd(x, y int) int {
	if x == 0 {
		return y
	}
	if y == 0 {
		return x
	}
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func modAccess(loc [][]bool, i, j, maxi, maxj int) bool {
	return loc[i%maxi][j%maxj]
}
