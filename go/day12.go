package main

import "fmt"

type plot struct {
	area      int
	perimeter int
}

func day12(lines []string) int {
	res := 0
	tracked := make(map[point]bool)

	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			point := point{y, x}

			if tracked[point] {
				continue
			}

			r := rune(lines[y][x])

			plot := plot{area: 0, perimeter: 0}

            xTracked := make(map[float32]bool)
            yTracked := make(map[float32]bool)

			calculatePlot(lines, y, x, r, &plot, &tracked, &xTracked, &yTracked)

			fmt.Println(string(r), plot.area, len(xTracked), len(yTracked))
			res += plot.area * (len(xTracked) + len(yTracked))

            fmt.Println(xTracked, yTracked)

		}
	}

	fmt.Println("Done?")

	return res
}

func calculatePlot(lines []string, y int, x int, r rune, p *plot, tracked *map[point]bool, xTracked *map[float32]bool, yTracked *map[float32]bool) {

	(*tracked)[point{y, x}] = true

	c := func(lines []string, y int, x int, xy *map[float32]bool, t float32) {

		if y < 0 || y >= len(lines) || x < 0 || x >= len(lines[y]) {
            (*xy)[t] = true
			return
		}

		if r != rune(lines[y][x]) {
            (*xy)[t] = true
			return
		}

		if (*tracked)[point{y, x}] {
			return
		}

		calculatePlot(lines, y, x, r, p, tracked, xTracked, yTracked)
	}

	c(lines, y-1, x, yTracked, float32(y)-0.1)
	c(lines, y, x-1, xTracked, float32(x)-0.1)
	c(lines, y, x+1, xTracked, float32(x)+0.1)
	c(lines, y+1, x, yTracked, float32(y)+0.1)

	p.area++
}

//part 1
// func day12(lines []string) int {
// 	res := 0
// 	tracked := make(map[point]bool)
//
// 	for y := 0; y < len(lines); y++ {
// 		for x := 0; x < len(lines[y]); x++ {
// 			point := point{y, x}
//
// 			if tracked[point] {
// 				continue
// 			}
//
// 			r := rune(lines[y][x])
//
// 			plot := plot{area: 0, perimeter: 0}
//
// 			calculatePlot(lines, y, x, r, &plot, &tracked)
//
// 			// fmt.Println(string(r), plot)
// 			res += plot.area * plot.perimeter
//
// 		}
// 	}
//
// 	fmt.Println("Done?")
//
// 	return res
// }
//
// func calculatePlot(lines []string, y int, x int, r rune, p *plot, tracked *map[point]bool) {
//
// 	(*tracked)[point{y, x}] = true
//
// 	c := func(lines []string, y int, x int) {
//
// 		if y < 0 || y >= len(lines) || x < 0 || x >= len(lines[y]) {
// 			p.perimeter++
// 			return
// 		}
//
// 		if r != rune(lines[y][x]) {
// 			p.perimeter++
// 			return
// 		}
//
// 		if (*tracked)[point{y, x}] {
// 			return
// 		}
//
// 		calculatePlot(lines, y, x, r, p, tracked)
// 	}
//
// 	c(lines, y-1, x)
// 	c(lines, y, x-1)
// 	c(lines, y, x+1)
// 	c(lines, y+1, x)
//
// 	p.area++
// }
