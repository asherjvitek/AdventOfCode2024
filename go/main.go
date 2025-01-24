package main

import "fmt"

func main() {
	path := "./inputs/day17.txt"
	lines, err := readLines(path)

    // for _, line := range lines {
    //     fmt.Println(line)
    // }


	check(err)

    // res, res2 := day17(lines)
    res := day17(lines)

    fmt.Println(res)
}
