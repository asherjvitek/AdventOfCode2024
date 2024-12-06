package main

import (
	"fmt"
	// "time"
)

type point struct {
	x int
	y int
}

func day6Part1(lines []string) int {
	grid := toRuneSlice(lines)

	pos := point{x: -1, y: -1}
	direction := direction{x: 0, y: -1}

	printGrid(grid)

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			r := lines[y][x]

			if r == '^' {
				pos = point{x: x, y: y}
				break
			}
		}

		if pos.x != -1 {
			break
		}
	}

	for isInGrid(pos, grid) {
		if checkDest(grid, pos, direction) {
			move(grid, &pos, direction)
		} else {
			rotate(&direction)
		}

	}

    fmt.Println()
	printGrid(grid)

	res := 0

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			r := grid[y][x]
			if r == 'X' {
				res++
			}
		}
	}

	return res
}

func printGrid(grid [][]rune) {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			fmt.Printf("%c", grid[y][x])
		}
		fmt.Printf("\n")
	}
}

func isInGrid(pos point, grid [][]rune) bool {
	return pos.x <= len(grid)-1 && pos.x >= 0 && pos.y <= len(grid[0])-1 && pos.y >= 0
}

func checkDest(grid [][]rune, pos point, direction direction) bool {
	if !isInGrid(point{y: pos.y + direction.y, x: pos.x + direction.x}, grid) {
		return true //dumb
	}

	r := grid[pos.y+direction.y][pos.x+direction.x]

	return r != '#'
}

func move(grid [][]rune, pos *point, direction direction) {

	grid[pos.y][pos.x] = 'X'

	pos.x += direction.x
	pos.y += direction.y

	if !isInGrid(*pos, grid) {
		return
	}

	grid[pos.y][pos.x] = '^'

}

func rotate(direction *direction) {
	if direction.x == 0 && direction.y == -1 { //up to right
		direction.x = 1
		direction.y = 0
	} else if direction.x == 1 && direction.y == 0 { //right to down
		direction.x = 0
		direction.y = 1
	} else if direction.x == 0 && direction.y == 1 { //down to left
		direction.x = -1
		direction.y = 0
	} else if direction.x == -1 && direction.y == 0 { //left to up
		direction.x = 0
		direction.y = -1
	}
}
