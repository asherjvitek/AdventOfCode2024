package main

import "fmt"

func main() {
	path := "./inputs/day14.txt"
	lines, err := readLines(path)

    // for _, line := range lines {
    //     fmt.Println(line)
    // }


	check(err)

    res := day14(lines)

    fmt.Println(res)
}
