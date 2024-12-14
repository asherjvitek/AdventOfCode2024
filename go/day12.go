package main

import "fmt"

type plot struct {
	area      int
	perimeter int
	locations map[point]bool
}

func day12(lines []string) int {
	res := 0
	tracked := make(map[point]bool)

	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			if tracked[point{y, x}] {
				continue
			}

			r := rune(lines[y][x])

			plot := plot{area: 0, perimeter: 0, locations: make(map[point]bool)}

			calculatePlot(lines, y, x, r, &plot, &tracked)

			sides := 0

			for k := range plot.locations {
				x := k.x
				y := k.y

				isOutsideCorner := func(yAdjust int, xAdjust int) bool {
					return !plot.locations[point{y: y + yAdjust, x: x}] &&
						!plot.locations[point{y: y, x: x + xAdjust}]
				}

				isInsideCorner := func(yAdjust int, xAdjust int) bool {
					return plot.locations[point{y: y + yAdjust, x: x}] &&
						plot.locations[point{y: y, x: x + xAdjust}] &&
						!plot.locations[point{y: y + yAdjust, x: x + xAdjust}]

				}

				if isOutsideCorner(-1, 1) {
					sides++
				}
				if isOutsideCorner(-1, -1) {
					sides++
				}
				if isOutsideCorner(1, -1) {
					sides++
				}
				if isOutsideCorner(1, 1) {
					sides++
				}

				if isInsideCorner(-1, 1) {
					sides++
				}
				if isInsideCorner(-1, -1) {
					sides++
				}
				if isInsideCorner(1, -1) {
					sides++
				}
				if isInsideCorner(1, 1) {
					sides++
				}
			}

			// fmt.Println(string(r), plot.area, sides)
			res += plot.area * sides

			// fmt.Println(perimeterPoints)

		}
	}

	fmt.Println("Done?")

	return res
}

func calculatePlot(lines []string, y int, x int, r rune, p *plot, tracked *map[point]bool) {

	(*tracked)[point{y, x}] = true
	(*p).locations[point{y, x}] = true

	c := func(lines []string, y int, x int) {

		if y < 0 || y >= len(lines) || x < 0 || x >= len(lines[y]) {
			return
		}

		if r != rune(lines[y][x]) {
			return
		}

		if (*tracked)[point{y, x}] {
			return
		}

		calculatePlot(lines, y, x, r, p, tracked)
	}

	c(lines, y-1, x)
	c(lines, y, x-1)
	c(lines, y, x+1)
	c(lines, y+1, x)

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
