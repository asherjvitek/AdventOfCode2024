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
			if tracked[point{y, x}] {
				continue
			}

			r := rune(lines[y][x])

			plot := plot{area: 0, perimeter: 0}
			perimeterPoints := make(map[point]int)

			calculatePlot(lines, y, x, r, &plot, &tracked, &perimeterPoints)

			p := 0
            corner := make(map[point]int)

			// fmt.Println(perimeterPoints)

			for k, v := range perimeterPoints {

				if v > 1 {
                    corner[k] += v
                    continue
				}

				checkSide := func(p point, c point) {
					if perimeterPoints[p] == 1 {
                        corner[c] = 1
					}
				}

                //f me.... this is double finding corners....
                checkSide(point{y: k.y - 1, x: k.x + 1}, point{y: k.y - 1, x: k.x})
				checkSide(point{y: k.y - 1, x: k.x - 1}, point{y: k.y - 1, x: k.x})
				checkSide(point{y: k.y + 1, x: k.x + 1}, point{y: k.y + 1, x: k.x})
				checkSide(point{y: k.y + 1, x: k.x - 1}, point{y: k.y + 1, x: k.x})

				p += v
			}

            sides := 0

            // fmt.Println(corner)


            for k, v := range corner {

                fmt.Println(k)
                sides += v
            }

			fmt.Println(string(r), plot.area, p, sides)
			res += plot.area * p

			// fmt.Println(perimeterPoints)

		}
	}

	fmt.Println("Done?")

	return res
}

func calculatePlot(lines []string, y int, x int, r rune, p *plot, tracked *map[point]bool, perimeterPoints *map[point]int) {

	(*tracked)[point{y, x}] = true

	c := func(lines []string, y int, x int, pp *map[point]int) {

		if y < 0 || y >= len(lines) || x < 0 || x >= len(lines[y]) {
			(*pp)[point{y, x}]++
			return
		}

		if r != rune(lines[y][x]) {
			(*pp)[point{y, x}]++
			return
		}

		if (*tracked)[point{y, x}] {
			return
		}

		calculatePlot(lines, y, x, r, p, tracked, pp)
	}

	c(lines, y-1, x, perimeterPoints)
	c(lines, y, x-1, perimeterPoints)
	c(lines, y, x+1, perimeterPoints)
	c(lines, y+1, x, perimeterPoints)

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
