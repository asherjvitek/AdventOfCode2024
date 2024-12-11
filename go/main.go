package main

import "fmt"

func main() {
	path := "./inputs/day11.txt"
	lines, err := readLines(path)

    // for _, line := range lines {
    //     fmt.Println(line)
    // }


	check(err)

    res := day11(lines)

    fmt.Println(res)
}
