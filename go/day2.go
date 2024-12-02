package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day2Part2(lines []string) int {
	res := 0

	for _, line := range lines {

		strNums := strings.Split(line, " ")

		if isValid(strNums) {
			res++
			continue
		}

		for i := 0; i < len(strNums); i++ {
			modified := removeItem(strNums, i)

			if isValid(modified) {
				res++
				break
			}
		}
	}

	return res
}

func removeItem[S ~[]E, E any](items S, index int) S {
	res := make(S, len(items)-1)

	offset := 0
	for i := 0; i < len(items); i++ {

		if i == index {
			offset = 1
			continue
		}

		res[i-offset] = items[i]
	}

	return res
}

func isValid(strNums []string) bool {
	asc := false
	desc := false

	for i, num := range strNums {
		if i == 0 || len(strNums) == i-1 {
			continue
		}

		num2, err := strconv.Atoi(num)
		num, err2 := strconv.Atoi(strNums[i-1])

		// fmt.Println(num, num2)

		check(err)
		check(err2)

		diff := 0

		if num < num2 {
			desc = true
			diff = num2 - num
		}

		if num > num2 {
			asc = true
			diff = num - num2
		}

		if asc && desc || diff > 3 || diff == 0 {
			return false
		}
	}

	return true
}

func day2Part1(lines []string) int {
	// fmt.Println("Day 2 go!\n", lines)
	res := 0

	for _, line := range lines {

		strNums := strings.Split(line, " ")
		asc := false
		desc := false
		valid := true

		for i, num := range strNums {
			if i == 0 || len(strNums) == i-1 {
				continue
			}

			num2, err := strconv.Atoi(num)
			num, err2 := strconv.Atoi(strNums[i-1])

			// fmt.Println(num, num2)

			check(err)
			check(err2)

			diff := 0

			if num < num2 {
				desc = true
				diff = num2 - num
			}

			if num > num2 {
				asc = true
				diff = num - num2
			}

			if asc && desc || diff > 3 || diff == 0 {
				valid = false
				break
			}
		}

		fmt.Println(valid)
		fmt.Println()

		if valid {
			res++
		}
	}

	return res
}
