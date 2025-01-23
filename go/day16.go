package main

import (
	"cmp"
	"fmt"
	"maps"
	"slices"
)

const north = '^'
const south = 'v'
const east = '>'
const west = '<'
const maxScore = 94444

type validPath struct {
	score int
	path  map[point]rune
}

func compareValidPaths(a, b validPath) int {
	return cmp.Compare(a.score, b.score)
}

var maze = make([]string, 0)

var trackedPaths = make(map[point]int)

func day16(lines []string) (int, int) {

	maze = lines
	start := point{}
	end := point{}
	path := make(map[point]rune, 0)
	direction := '>'

	for y := 0; y < len(lines); y++ {
		line := lines[y]
		// fmt.Println()
		for x := 0; x < len(line); x++ {
			char := rune(line[x])
			// fmt.Print(string(char))
			switch char {
			case '.':
				path[point{y, x}] = '.'
			case 'S':
				start = point{y, x}
			case 'E':
				end = point{y, x}
			}
		}
	}

	fmt.Println("start:", start, "end:", end)
	scores := make([]validPath, 0)

	path[start] = 'S'
	path[end] = 'E'

	getPaths(path, end, start, 0, &scores, direction)

	fmt.Println("scores", scores)

	low := slices.MinFunc(scores, compareValidPaths).score

	tiles := make(map[point]bool)
	for _, p := range scores {
		if p.score == low {
			printMaze(lines, p.path)
			for point, r := range p.path {
				if r != '.' {
					tiles[point] = true
				}
			}
		}
	}

	return low, len(tiles)
}

func getPaths(path map[point]rune, end point, current point, score int, finished *[]validPath, direction rune) {
    if score > maxScore {
        return
    }

	path[current] = direction
	// printMaze(maze, path)


	if trackedPaths[current] > 0 && trackedPaths[current] < score - 1000 {
		return
	}

	trackedPaths[current] = score


	if len(*finished) > 0 && score > slices.MinFunc(*finished, compareValidPaths).score {
		return
	}

	doThing := func(y int, x int, scoreAdjust int, direction rune) {
		next := point{current.y + y, current.x + x}

		if next == end {
			score += scoreAdjust
			*finished = append(*finished, validPath{score, path})
			fmt.Println("Hit end!", score)
			return
		}

		if path[next] == '.' {
			getPaths(maps.Clone(path), end, next, score+scoreAdjust, finished, direction)
		}
	}

	switch direction {
	case north:
		{
			//forward
			doThing(-1, 0, 1, direction)

			//left
			doThing(0, -1, 1001, west)

			//right
			doThing(0, 1, 1001, east)
		}
	case east:
		{
			//forward
			doThing(0, 1, 1, direction)

			//up
			doThing(-1, 0, 1001, north)

			//down
			doThing(1, 0, 1001, south)
		}
	case south:
		{
			//forward
			doThing(1, 0, 1, direction)

			//left
			doThing(0, -1, 1001, west)

			//right
			doThing(0, 1, 1001, east)
		}
	case west:
		{
			//forward
			doThing(0, -1, 1, direction)

			//up
			doThing(-1, 0, 1001, north)

			//down
			doThing(1, 0, 1001, south)
		}
	}
}

func printMaze(lines []string, path map[point]rune) {
	for y := 0; y < len(lines); y++ {
		line := lines[y]
		fmt.Println()
		for x := 0; x < len(line); x++ {
			char := rune(line[x])
			switch char {
			case '.':
				fmt.Print(string(path[point{y, x}]))
			case 'S':
				fmt.Print(string(path[point{y, x}]))
			case 'E':
				fmt.Print(string(path[point{y, x}]))
			default:
				fmt.Print(string(char))
			}
		}
	}

	fmt.Println()
}
