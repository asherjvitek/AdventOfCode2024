package main

import (
	"fmt"
	"strconv"
	"strings"
)

const gridSize = 7
const amount = 12

var end = point{gridSize - 1, gridSize - 1}

func day18(lines []string) int {
	points := getPoints(lines)
	printMem(points)

	for p, _ := range points {
		fmt.Println(p.x, p.y)
	}

	track := make(map[point]int)

	path := make([]point, 0)

	findWay(points, point{0, 0}, 0, &track, path)

	fmt.Println(track[end])

	return 0
}

func findWay(points map[point]bool, current point, count int, track *map[point]int, path []point) {
	// fmt.Println(current)
	if current == end {
		(*track)[current] = count
		fmt.Println(count+1, path)
		return
	}

	if (*track)[current] > 0 && (*track)[current] < count {
		return
	}

	if points[current] {
		// fmt.Println(current)
		return
	}

	if current.x < 0 || current.y < 0 || current.x > gridSize-1 || current.y > gridSize-1 {
		return
	}

	(*track)[current] = count
	path = append(path, current)

	findWay(points, point{x: current.x - 1, y: current.y}, count+1, track, path)
	findWay(points, point{x: current.x + 1, y: current.y}, count+1, track, path)
	findWay(points, point{x: current.x, y: current.y - 1}, count+1, track, path)
	findWay(points, point{x: current.x, y: current.y + 1}, count+1, track, path)
}

func getPoints(lines []string) map[point]bool {
	points := make(map[point]bool)

	for i, line := range lines {
		if i > amount {
			break
		}
		split := strings.Split(line, ",")

		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])

		points[point{x, y}] = true
	}

	return points
}

func printMem(points map[point]bool) {
	for y := 0; y < gridSize; y++ {
		for x := 0; x < gridSize; x++ {
			if points[point{x, y}] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
