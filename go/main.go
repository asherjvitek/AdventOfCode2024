package main

import "fmt"

func main() {
	path := "./inputs/day16.txt"
	lines, err := readLines(path)

    // for _, line := range lines {
    //     fmt.Println(line)
    // }


	check(err)

    res, res2 := day16(lines)

    fmt.Println(res, res2)
}
