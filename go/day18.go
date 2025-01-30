package main

import (
	"fmt"
	"strconv"
	"strings"
)

const gridSize = 71
const amount = 3000

var end = point{gridSize - 1, gridSize - 1}

func day18(lines []string) int {
	points := getPoints(lines, amount)
	printMem(points)

	track := make(map[point]int)

	findWay(points, point{0, 0}, 0, &track)

	fmt.Println("here", track[end])

	amount2 := amount
	for {
		amount2++
		points := getPoints(lines, amount2)
		// printMem(points)

		track := make(map[point]int)

		findWay(points, point{0, 0}, 0, &track)

		if track[end] == 0 {
			fmt.Println(amount2, lines[amount2])
			return 0
		}
	}

	// return 0
}

func findWay(points map[point]bool, current point, count int, track *map[point]int) {
	if current == end && (*track)[current] > count {
		(*track)[current] = count
		return
	}

	if (*track)[current] > 0 && (*track)[current] <= count {
		return
	}

	if points[current] {
		return
	}

	if current.x < 0 || current.y < 0 || current.x > gridSize-1 || current.y > gridSize-1 {
		return
	}

	(*track)[current] = count

	findWay(points, point{x: current.x - 1, y: current.y}, count+1, track)
	findWay(points, point{x: current.x + 1, y: current.y}, count+1, track)
	findWay(points, point{x: current.x, y: current.y - 1}, count+1, track)
	findWay(points, point{x: current.x, y: current.y + 1}, count+1, track)
}

func getPoints(lines []string, amount int) map[point]bool {
	points := make(map[point]bool)

	for i, line := range lines {
		if i > amount {
			break
		}
		split := strings.Split(line, ",")

		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])

		points[point{y, x}] = true
	}

	return points
}

func printMem(points map[point]bool) {
	for y := 0; y < gridSize; y++ {
		for x := 0; x < gridSize; x++ {
			if points[point{y, x}] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
