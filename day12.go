package main

import (
	"fmt"
)

func init() {
	codeForDay[12] = day12
}

type pos3 struct {
	x, y, z int
}

func (a pos3) minus(b pos3) pos3 {
	a.x -= b.x
	a.y -= b.y
	a.z -= b.z
	return a
}

func (a pos3) plus(b pos3) pos3 {
	a.x += b.x
	a.y += b.y
	a.z += b.z
	return a
}

func (a pos3) grav(b pos3) pos3 {
	c := b.minus(a)
	c.x = intSign(c.x)
	c.y = intSign(c.y)
	c.z = intSign(c.z)
	return c
}

func intSign(a int) int {
	if a < 0 {
		return -1
	} else if a > 0 {
		return 1
	}
	return 0
}

func day12(example int) {
	moonsPos := []pos3{
		pos3{-17, 9, -5},
		pos3{-1, 7, 13},
		pos3{-19, 12, 5},
		pos3{-6, -6, -4},
	}
	moonsVel := make([]pos3, 4)
	steps := 1000

	switch *exampleFlag {
	case 1:
		moonsPos = []pos3{
			{-1, 0, 2},
			{2, -10, -7},
			{4, -8, 8},
			{3, 5, -1},
		}
		steps = 10
	}

	for i := 0; i < steps; i++ {
		moonsPos, moonsVel = stepMoons(moonsPos, moonsVel)
	}

	fmt.Println("D12-P1:", energy(moonsPos, moonsVel)) //8742

	xPeriod := axisRepeats([]int{-17, -1, -19, -6}) // Prime factor: 2*2*46507
	yPeriod := axisRepeats([]int{9, 7, 12, -6})     // Prime factor: 2*115807
	zPeriod := axisRepeats([]int{-5, 13, 5, -4})    // Prime factor: 2*2*2*7*13*83

	xyGCD := gcd(xPeriod, yPeriod)
	xPeriod /= xyGCD
	zPeriod /= xyGCD
	if zPeriod/xyGCD*xyGCD == zPeriod {
		zPeriod /= xyGCD
	}
	xzGCD := gcd(xPeriod, zPeriod)
	xPeriod /= xzGCD
	zPeriod /= xzGCD
	yzGCD := gcd(yPeriod, zPeriod)
	yPeriod /= yzGCD
	zPeriod /= yzGCD

	fmt.Println("D12-P2:", xyGCD*xzGCD*yzGCD*xPeriod*yPeriod*zPeriod) // Resulting in rank 258 on today's leaderboard
}

func axisRepeats(orig []int) int {
	pos := make([]int, len(orig))
	copy(pos, orig)
	vel := []int{0, 0, 0, 0}
	i := 0
	for {
		pos, vel = stepAxis(pos, vel)
		i++
		if pos[0] == orig[0] && pos[1] == orig[1] && pos[2] == orig[2] && pos[3] == orig[3] {
			if vel[0] == 0 && vel[1] == 0 && vel[2] == 0 && vel[3] == 0 {
				break
			}
		}
	}
	return i
}

func stepAxis(pos, vel []int) ([]int, []int) {
	for i := 0; i < len(pos); i++ {
		for j := i + 1; j < len(pos); j++ {
			grav := intSign(pos[j] - pos[i])
			vel[i] += grav
			vel[j] -= grav
		}
	}
	for i := 0; i < len(pos); i++ {
		pos[i] += vel[i]
	}
	return pos, vel
}

func stepMoons(pos, vel []pos3) ([]pos3, []pos3) {
	for i := 0; i < len(pos); i++ {
		for j := 0; j < len(pos); j++ {
			gravC := pos[i].grav(pos[j])
			vel[i] = vel[i].plus(gravC)
		}
	}
	for i := 0; i < len(pos); i++ {
		pos[i] = pos[i].plus(vel[i])
	}
	return pos, vel
}

func energy(pos, vel []pos3) int {
	en := 0
	for i := 0; i < len(pos); i++ {
		en += potE(vel[i]) * potE(pos[i])
	}
	return en
}

func potE(pos pos3) int {
	return intAbs(pos.x) + intAbs(pos.y) + intAbs(pos.z)
}

func intAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
