package main

import "fmt"

func main() {
	path := "./inputs/day15.txt"
	lines, err := readLines(path)

    // for _, line := range lines {
    //     fmt.Println(line)
    // }


	check(err)

    res := day15(lines)

    fmt.Println(res)
}
