package main

import "fmt"

func main() {
	path := "./inputs/day19.txt"
	lines, err := readLines(path)

    // for _, line := range lines {
    //     fmt.Println(line)
    // }


	check(err)

    // res, res2 := day17(lines)
    res := day19(lines)

    fmt.Println(res)
}
