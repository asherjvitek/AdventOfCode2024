package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func day1Part2(dat []byte) {
    data := string(dat)

	// fmt.Println(data)

	lines := strings.Split(data, "\n")
	len := len(lines) - 1

	left := make([]int, len)
	right := make([]int, len)

	for i := 0; i < len; i++ {
		line := lines[i]
		parts := strings.Split(line, "   ")

        // fmt.Printf("%s, %s\n", parts[0], parts[1])
		numLeft, err := strconv.Atoi(parts[0])

		check(err)

		numRight, err := strconv.Atoi(parts[1])

		check(err)

		left[i] = numLeft
		right[i] = numRight
	}

	slices.Sort(left)
	slices.Sort(right)

	var res int

	for n := range left {
        count := 0
        for i := range right {
            if right[i] == left[n] {
                count++
            }
        }

        res += left[n] * count
	}

	fmt.Println(res)

}

func day1Part1(dat []byte) {
    data := string(dat)

	// fmt.Println(data)

	lines := strings.Split(data, "\n")
	len := len(lines) - 1

	left := make([]int, len)
	right := make([]int, len)

	for i := 0; i < len; i++ {
		line := lines[i]
		parts := strings.Split(line, "   ")

        // fmt.Printf("%s, %s\n", parts[0], parts[1])
		numLeft, err := strconv.Atoi(parts[0])

		check(err)

		numRight, err := strconv.Atoi(parts[1])

		check(err)

		left[i] = numLeft
		right[i] = numRight
	}

	slices.Sort(left)
	slices.Sort(right)

	var res int

	for n := range left {
		if left[n] > right[n] {
			res += left[n] - right[n]
		} else {
			res += right[n] - left[n]
		}
	}

	fmt.Println(res)

}
