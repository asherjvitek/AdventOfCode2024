package main

import (
	"fmt"
	"strconv"
)

type trail struct {
    start point
    end point
}

var trails = make(map[trail]bool)

func day10(lines []string) int {
    fmt.Println("Day 10!")
	res := 0

	for y, line := range lines {
		for x, r := range line {
			if r == '0' {
                score := search(lines, point{y, x}, y, x)

                fmt.Println(score, y, x)

				res += score
			}

		}
	}

	return res
}

func search(lines []string, start point, y int, x int) int {
	res := 0

	curr, err := strconv.Atoi(string(lines[y][x]))

    if err != nil {
        return res
    }

    if (curr == 9) {
        // fmt.Println("here")

        return 1
    }

    if y - 1 >= 0  {
        poss, err := strconv.Atoi(string(lines[y-1][x]))
        if err == nil && poss == curr + 1 {
            res += search(lines, start, y-1, x)
        }
    }

    if x - 1 >= 0  {
        poss, err := strconv.Atoi(string(lines[y][x-1]))
        if err == nil && poss == curr + 1 {
            res += search(lines, start, y, x-1)
        }
    }

    if y + 1 < len(lines)  {
        poss, err := strconv.Atoi(string(lines[y+1][x]))
        if err == nil && poss == curr + 1 {
            res += search(lines, start, y+1, x)
        }
    }


    if x + 1 < len(lines[y])  {
        poss, err := strconv.Atoi(string(lines[y][x+1]))
        if err == nil && poss == curr + 1 {
            res += search(lines, start, y, x+1)
        }
    }


	return res
}


func day10Part1(lines []string) int {
	fmt.Println("Day 10!")
	res := 0

	for y, line := range lines {
		for x, r := range line {
			if r == '0' {
				var trails = make(map[trail]bool)
				search1(trails, lines, point{y, x}, y, x)

                score := len(trails)

                res += score

				fmt.Println(score, y, x)
			}

		}
	}

	return res
}

func search1(trails map[trail]bool, lines []string, start point, y int, x int) {
	curr, err := strconv.Atoi(string(lines[y][x]))

	if err != nil {
		return
	}

	if curr == 9 {
		trails[trail{start: start, end: point{y, x}}] = true
	}

	if y-1 >= 0 {
		poss, err := strconv.Atoi(string(lines[y-1][x]))
		if err == nil && poss == curr+1 {
			search1(trails, lines, start, y-1, x)
		}
	}

	if x-1 >= 0 {
		poss, err := strconv.Atoi(string(lines[y][x-1]))
		if err == nil && poss == curr+1 {
			search1(trails, lines, start, y, x-1)
		}
	}

	if y+1 < len(lines) {
		poss, err := strconv.Atoi(string(lines[y+1][x]))
		if err == nil && poss == curr+1 {
			search1(trails, lines, start, y+1, x)
		}
	}

	if x+1 < len(lines[y]) {
		poss, err := strconv.Atoi(string(lines[y][x+1]))
		if err == nil && poss == curr+1 {
			search1(trails, lines, start, y, x+1)
		}
	}
}
