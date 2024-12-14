package main

import "fmt"

func main() {
	path := "./inputs/day13.txt"
	lines, err := readLines(path)

    // for _, line := range lines {
    //     fmt.Println(line)
    // }


	check(err)

    res := day13(lines)

    fmt.Println(res)
}
