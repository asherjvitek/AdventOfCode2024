package main

import (
	"fmt"
	// "time"
)

type point struct {
	y int
	x int
}

type history struct {
	point     point
	direction direction
}

func day6Part2(lines []string) int {
	res := 0
	srcGrid := toRuneSlice(lines)
	printGrid(srcGrid)

	pos, error := getStartingPos(srcGrid)

	if error != nil {
		panic(error)
	}

	for y := 0; y < len(srcGrid); y++ {
		for x := 0; x < len(srcGrid[y]); x++ {

			if srcGrid[y][x] == '#' || srcGrid[y][x] == '^' {
				continue
			}

			//copy the grid
			grid := make([][]rune, len(srcGrid))

			for i, line := range srcGrid {
				grid[i] = make([]rune, len(line))
				copy(grid[i], line)
			}

			//set current location to #
			grid[y][x] = '#'

			// printGrid(grid)

			h := make(map[history]int)

			pos := point{y: pos.y, x: pos.x}
			direction := direction{x: 0, y: -1}

			for isInGrid(pos, grid) {
				if checkDest(grid, pos, direction) {
					move(grid, &pos, direction)
					history := history{point: pos, direction: direction}
					h[history]++
					hist := h[history]

					if hist > 1 {
						res++
						// printGrid(grid)
						break
					}
				} else {
					rotate(&direction)
				}
			}
		}
	}

	return res
}

func day6Part1(lines []string) int {
	grid := toRuneSlice(lines)

	pos, error := getStartingPos(grid)

	if error != nil {
		panic(error)
	}

	direction := direction{x: 0, y: -1}

	printGrid(grid)

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
	fmt.Println("----------------------------------------------------")
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			fmt.Printf("%c", grid[y][x])
		}
		fmt.Printf("\n")
	}
	fmt.Println("----------------------------------------------------")
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

func getStartingPos(grid [][]rune) (point, error) {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			r := grid[y][x]

			if r == '^' {
				return point{x: x, y: y}, nil
			}
		}
	}

	return point{}, fmt.Errorf("No Starting Position")

}
