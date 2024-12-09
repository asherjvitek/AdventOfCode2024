package main

import (
	"fmt"
	"sort"
)

type signal struct {
	location point
	val      rune
}

func day8Part1(lines []string) int {
	fmt.Println("day 8")

	var signals []signal
	antinodes := make(map[point]bool)

	xBound := len(lines[0])
	yBound := len(lines)

	fmt.Println(xBound, yBound)

	result := make([][]rune, len(lines))

	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			r := rune(lines[y][x])
			result[y] = append(result[y], r)
			if r == '.' {
				continue
			}
			signal := signal{location: point{y: y, x: x}, val: r}
			signals = append(signals, signal)
		}
	}

	for i, sig := range signals {
		fmt.Println("---------------------")
		for j := i + 1; j < len(signals); j++ {
			sig2 := signals[j]
			if sig2.val != sig.val {
				continue
			}

			diffx := sig.location.x - sig2.location.x
			diffy := sig.location.y - sig2.location.y

			added := false
			multiplier := 1

			antinodes[sig.location] = true
			antinodes[sig2.location] = true

			for added || multiplier == 1 {
				added = false

				diffx := diffx * multiplier
				diffy := diffy * multiplier

				antinode1 := point{y: sig.location.y + diffy, x: sig.location.x + diffx}
				antinode2 := point{y: sig2.location.y - diffy, x: sig2.location.x - diffx}

				fmt.Println()
				fmt.Printf("%+v val: %c\n", sig.location, sig.val)
				fmt.Printf("%+v val: %c\n", sig2.location, sig2.val)

				if antinode1.x >= 0 && antinode1.x < xBound && antinode1.y >= 0 && antinode1.y < yBound {
					antinodes[antinode1] = true
					result[antinode1.y][antinode1.x] = '#'
					added = true
				}

				if antinode2.x >= 0 && antinode2.x < xBound && antinode2.y >= 0 && antinode2.y < yBound {
					antinodes[antinode2] = true
					result[antinode2.y][antinode2.x] = '#'
					added = true
				}

				multiplier++
			}
		}
	}

	fmt.Println()

	// sort.Slice(antinodes, func(i, j int) bool {
	// 	return antinodes[i].location.y < antinodes[j].location.y
	// })
	//
	// for _, anti := range antinodes {
	// 	fmt.Printf("%+v val: %c\n", anti.location, anti.val)
	// }

	for y := 0; y < len(result); y++ {
		for x := 0; x < len(result[y]); x++ {
			fmt.Printf("%c", result[y][x])
		}
		fmt.Println()
	}

	return len(antinodes)
}
