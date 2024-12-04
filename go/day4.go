package main

import (
	"fmt"
)

type direction struct {
	x int
	y int
}

const xmas = "XMAS"

var directionMap = map[string]direction{
	"right":     {x: 1, y: 0},
	"rightDown": {x: 1, y: -1},
	"down":      {x: 0, y: -1},
	"downLeft":  {x: -1, y: -1},
	"left":      {x: -1, y: 0},
	"leftUp":    {x: -1, y: 1},
	"up":        {x: 0, y: 1},
	"upRight":   {x: 1, y: 1},
}

func day4Part2(lines []string) int {
	// fmt.Println(lines)

	res := 0

    // fmt.Println(len(lines))

	for y := 1; y < len(lines)-1; y++ {
		for x := 1; x < len(lines[y])-1; x++ {
			r := lines[y][x]

			if r != 'A' {
				continue
			}

			masCount := 0

			masCount += checkForMas(lines, x, y, "rightDown")
			masCount += checkForMas(lines, x, y, "downLeft")
			masCount += checkForMas(lines, x, y, "leftUp")
			masCount += checkForMas(lines, x, y, "upRight")

			if masCount == 2 {
				res++
			}
		}

		// fmt.Println()
	}

	return res
}

func day4Part1(lines []string) int {
	fmt.Println(lines)

	res := 0

	for y, line := range lines {
		for x, r := range line {
			// fmt.Printf("%c", lines[y][x])

			if r != 'X' {
				continue
			}

			res += checkForXmas(lines, x, y, "right")
			res += checkForXmas(lines, x, y, "rightDown")
			res += checkForXmas(lines, x, y, "down")
			res += checkForXmas(lines, x, y, "downLeft")
			res += checkForXmas(lines, x, y, "left")
			res += checkForXmas(lines, x, y, "leftUp")
			res += checkForXmas(lines, x, y, "up")
			res += checkForXmas(lines, x, y, "upRight")
		}

		// fmt.Println()
	}

	return res
}

func checkForMas(lines []string, x int, y int, direct string) int {

	direction := directionMap[direct]

	x = x + direction.x*-1
	y = y + direction.y*-1

	if lines[y][x] != xmas[1] {
		return 0
	}

	x = x + direction.x + direction.x
	y = y + direction.y + direction.y

	if lines[y][x] != xmas[3] {
		return 0
	}

	return 1
}

func checkForXmas(lines []string, x int, y int, direct string) int {

	direction := directionMap[direct]

	for i := 0; i < 3; i++ {
		x = x + direction.x
		y = y + direction.y

		if y < 0 || y > len(lines)-1 || x < 0 || x > len(lines[y])-1 {
			return 0
		}

		if lines[y][x] != xmas[i+1] {
			return 0
		}
	}

	return 1
}
