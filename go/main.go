package main

import "fmt"

func main() {
	path := "./inputs/day18.txt"
	lines, err := readLines(path)

    // for _, line := range lines {
    //     fmt.Println(line)
    // }


	check(err)

    // res, res2 := day17(lines)
    res := day18(lines)

    fmt.Println(res)
}
