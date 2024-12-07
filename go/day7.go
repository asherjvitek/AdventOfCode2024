package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day7Part2(lines []string) int {
	res := 0
	fmt.Println("Day 7")

	for _, line := range lines {
		// fmt.Println(line)

		split := strings.Split(line, ": ")

		var possibilities []int
		lookingFor, _ := strconv.Atoi(split[0])
		nums := strings.Split(split[1], " ")

		for i, num := range nums {
			n, _ := strconv.Atoi(num)

			if i == 0 {
				possibilities = append(possibilities, n)
				continue
			}

			pl := len(possibilities)

			for pi := 0; pi < pl; pi++ {
				val := possibilities[pi]

				possibilities[pi] += n
				possibilities = append(possibilities, val*n)
                con, _ := strconv.Atoi(strconv.Itoa(val) + strconv.Itoa(n))
				possibilities = append(possibilities, con)
			}
		}

		for _, pos := range possibilities {
			if pos == lookingFor {
				res += lookingFor
				// fmt.Println(pos)
                break
			}
		}

	}

	return res
}

func day7Part1(lines []string) int {
	res := 0
	fmt.Println("Day 7")

	for _, line := range lines {
		// fmt.Println(line)

		split := strings.Split(line, ": ")

		var possibilities []int
		lookingFor, _ := strconv.Atoi(split[0])
		nums := strings.Split(split[1], " ")

		for i, num := range nums {
			n, _ := strconv.Atoi(num)

			if i == 0 {
				possibilities = append(possibilities, n)
				continue
			}

			pl := len(possibilities)

			for pi := 0; pi < pl; pi++ {
				val := possibilities[pi]

				possibilities[pi] += n
				possibilities = append(possibilities, val*n)
			}
		}

		for _, pos := range possibilities {
			if pos == lookingFor {
				res += lookingFor
				// fmt.Println(pos)
                break
			}
		}

	}

	return res
}
