package main

import "fmt"

func main() {
	path := "./inputs/day7.txt"
	lines, err := readLines(path)

    // for _, line := range lines {
    //     fmt.Println(line)
    // }

	check(err)

    res := day7Part2(lines)

    fmt.Println(res)
}
