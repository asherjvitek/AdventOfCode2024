package main

import "fmt"

func main() {
	path := "./inputs/day10.txt"
	lines, err := readLines(path)

    // for _, line := range lines {
    //     fmt.Println(line)
    // }


	check(err)

    res := day10(lines)

    fmt.Println(res)
}
