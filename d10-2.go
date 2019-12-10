package main

import (
	"fmt"
	"math"
	"sort"
)

func day10_2() {
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
	// locations := [][]bool{
	// 	[]bool{false, true, false, false, true, true, false, true, true, true, false, false, false, true, true, true, true, true, true, true},
	// 	[]bool{true, true, false, true, true, true, true, true, true, true, true, true, true, true, true, false, false, true, true, false},
	// 	[]bool{false, true, false, true, true, true, true, true, true, false, true, true, true, true, true, true, true, true, false, true},
	// 	[]bool{false, true, true, true, false, true, true, true, true, true, true, true, false, true, true, true, true, false, true, false},
	// 	[]bool{true, true, true, true, true, false, true, true, false, true, false, true, true, false, true, true, true, false, true, true},
	// 	[]bool{false, false, true, true, true, true, true, false, false, true, false, true, true, true, true, true, true, true, true, true},
	// 	[]bool{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true},
	// 	[]bool{true, false, true, true, true, true, false, false, false, false, true, true, true, false, true, false, true, false, true, true},
	// 	[]bool{true, true, false, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true},
	// 	[]bool{true, true, true, true, true, false, true, true, false, true, true, true, false, false, true, true, true, true, false, false},
	// 	[]bool{false, false, true, true, true, true, true, true, false, false, true, true, false, true, true, true, true, true, true, true},
	// 	[]bool{true, true, true, true, false, true, true, false, true, true, true, true, false, false, false, true, true, false, false, true},
	// 	[]bool{false, true, true, true, true, true, false, false, true, false, true, true, true, true, true, true, false, true, true, true},
	// 	[]bool{true, true, false, false, false, true, false, true, true, true, true, true, true, true, true, true, true, false, false, false},
	// 	[]bool{true, false, true, true, true, true, true, true, true, true, true, true, false, true, true, true, true, true, true, true},
	// 	[]bool{false, true, true, true, true, false, true, false, true, true, true, false, true, true, true, false, true, false, true, true},
	// 	[]bool{false, false, false, false, true, true, false, true, true, false, true, true, true, false, false, true, true, true, true, true},
	// 	[]bool{false, true, false, true, false, true, true, true, true, true, true, true, true, true, true, true, false, true, true, true},
	// 	[]bool{true, false, true, false, true, false, true, true, true, true, true, false, true, true, true, true, false, true, true, true},
	// 	[]bool{true, true, true, false, true, true, false, true, true, true, true, false, true, true, false, true, false, false, true, true},
	// }

	maxi := len(locations)
	tans := make([][]float64, maxi)
	maxj := len(locations[0])
	for i := range locations {
		tans[i] = make([]float64, maxj)
	}

	tari := 31
	tarj := 25
	// tari := 13
	// tarj := 11
	for thisi, row := range locations {
		for thisj, val := range row {
			if !val {
				continue
			}
			tans[thisi][thisj] = calcTan(thisi, thisj, tari, tarj)
		}
	}
	findShadows(locations, tans, tari, tarj)

	trips := make(map[float64]point, maxi)
	keys := make([]float64, 0, maxi*maxj)

	for thisi, row := range tans {
		for thisj, tan := range row {
			if !locations[thisi][thisj] || (thisi == tari && thisj == tarj) {
				continue
			}
			if _, ok := trips[tan]; ok {
				fmt.Println("OHDHNO", thisi, thisj, tan)
			}
			trips[tan] = point{thisi, thisj}
			keys = append(keys, tan)
		}
	}
	sort.Float64s(keys)

	fmt.Println("D10-D2:", trips[keys[199]])
}

func calcTan(thisi, thisj, tari, tarj int) float64 {
	tan := math.Atan2(float64(thisj-tarj), -float64(thisi-tari))
	if tan < 0 {
		tan += 2 * math.Pi
	}
	// fmt.Printf("Pos %d,%d gets %f\n", thisi, thisj, tan)
	return tan
}

func findShadows(loc [][]bool, tans [][]float64, thisi, thisj int) {
	maxi := len(loc)
	maxj := len(loc[0])

	for i := thisi; i < maxi; i++ {
		for j := thisj; j < maxj; j++ {
			if !loc[i][j] || (i == thisi && j == thisj) {
				continue
			}
			increaseHidden(tans, thisi, thisj, i, j, maxi, maxj)
		}
	}
	for i := thisi; i < maxi; i++ {
		for j := thisj - 1; j >= 0; j-- {
			if !loc[i][j] || (i == thisi && j == thisj) {
				continue
			}
			increaseHidden(tans, thisi, thisj, i, j, maxi, maxj)
		}
	}
	for i := thisi - 1; i >= 0; i-- {
		for j := thisj - 1; j >= 0; j-- {
			if !loc[i][j] || (i == thisi && j == thisj) {
				continue
			}
			increaseHidden(tans, thisi, thisj, i, j, maxi, maxj)
		}
	}
	for i := thisi - 1; i >= 0; i-- {
		for j := thisj; j < maxj; j++ {
			if !loc[i][j] || (i == thisi && j == thisj) {
				continue
			}
			increaseHidden(tans, thisi, thisj, i, j, maxi, maxj)
		}
	}
}

func increaseHidden(tans [][]float64, thisi, thisj, i, j, maxi, maxj int) {
	diffi, diffj := minDiff(i-thisi, j-thisj)
	i += diffi
	j += diffj
	for i >= 0 && i < maxi && j >= 0 && j < maxj {
		if thisi != i || thisj != j {
			tans[i][j] += 10.0
		}
		i += diffi
		j += diffj
	}
}
