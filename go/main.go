package main

import "fmt"

func main() {
	path := "./inputs/day5.txt"
	lines, err := readLines(path)

    // for _, line := range lines {
    //     fmt.Println(line)
    // }

	check(err)

    res := day5Part2(lines)

    fmt.Println(res)
}
