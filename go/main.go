package main

import "fmt"

func main() {
	path := "./inputs/day6.txt"
	lines, err := readLines(path)

    // for _, line := range lines {
    //     fmt.Println(line)
    // }

	check(err)

    res := day6Part1(lines)

    fmt.Println(res)
}
