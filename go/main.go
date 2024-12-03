package main

import "fmt"

func main() {
	path := "./inputs/day3.txt"
	lines, err := readLines(path)

    // for _, line := range lines {
    //     fmt.Println(line)
    // }

	check(err)

    res := day3Part1(lines)

    fmt.Println(res)
}
