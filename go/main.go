package main

import "fmt"

func main() {
	path := "./inputs/day12.txt"
	lines, err := readLines(path)

    // for _, line := range lines {
    //     fmt.Println(line)
    // }


	check(err)

    res := day12(lines)

    fmt.Println(res)
}
