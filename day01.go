package main

import "fmt"

func init() {
	// Originally implemented at https://play.golang.org/p/eHmrZ0nympS
	codeForDay[1] = day1
}

func day1(example int) {
	vals := []int{99603, 121503, 86996, 72052, 112039, 106616, 123581, 123171, 52480, 68686, 66395, 102661, 110250, 73289, 105725, 123802, 75488, 79426, 98634, 76095, 50852, 141405, 112388, 72180, 103300, 124602, 104531, 94751, 63270, 139027, 145939, 62275, 91812, 74751, 144010, 60221, 62821, 51080, 149802, 53067, 102574, 131339, 78942, 88430, 105314, 72764, 55214, 79095, 97458, 68699, 106974, 141492, 57673, 141866, 139355, 134222, 52145, 83293, 144322, 70741, 107873, 123638, 141011, 133249, 99065, 120480, 100767, 136550, 147323, 146988, 65583, 141287, 53097, 50662, 121124, 94886, 59344, 93981, 112492, 149136, 56647, 96430, 63968, 117987, 138475, 125958, 74967, 64480, 104644, 70273, 50671, 147116, 147101, 89096, 94697, 83282, 74533, 68418, 145578, 59032}

	fuel := 0
	for _, v := range vals {
		fuel += modFuel(v, true)
	}
	fmt.Println("D1-P1: ", fuel)

	fuel = 0
	for _, v := range vals {
		fuel += modFuel(v, false)
	}
	fmt.Println("D1-P2: ", fuel)
}

func modFuel(weight int, simple bool) int {
	fuel := weight/3 - 2

	if simple {
		return fuel
	}

	needsf := fuel
	for needsf > 0 {
		newf := needsf/3 - 2
		if newf > 0 {
			fuel += newf
		}
		needsf = newf
	}

	return fuel
}
