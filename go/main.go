package main

import (
	"fmt"
	"os"
)

func check(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func main() {
	path := "./inputs/day1.txt"
	dat, err := os.ReadFile(path)

	check(err)

	// day1_part1(dat)
	day1Part2(dat)
}
