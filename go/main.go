package main

import "fmt"

func main() {
	path := "./inputs/day2.txt"
	lines, err := readLines(path)

	check(err)

    res := day2Part2(lines)

    fmt.Println(res)
}
