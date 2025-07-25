package main

import (
	"fmt"
	"math"
	"slices"
)

func day20_part2(lines []string) int {
	res := 0
	saved := make(map[int]int)
	path := path(lines)
	fmt.Println(len(path))

	for i, p := range path {
		for i2, p2 := range path {
			distance := math.Abs(float64(p.x-p2.x)) + math.Abs(float64(p.y-p2.y))
			dist := float64(i2 - i)
            result := int(dist-distance)

			if distance < dist && distance <= 20 && result >= 100 {
                res++
                saved[result]++
				// fmt.Println(p, p2, distance, res)
			}
		}

		// break
	}

	fmt.Println(saved)

	return res
}

func day20(lines []string) int {
	res := 0
	saved := make(map[int]int)
	p := path(lines)
	fmt.Println(len(p))

	canCheat := func(poi point, y int, x int, i int) {
		dest := slices.Index(p, point{poi.y + y + y, poi.x + x + x})
		wall := slices.Index(p, point{poi.y + +y, poi.x + x})

		if dest > 0 && wall == -1 && dest > i {
			if dest-i-2 >= 100 {
				res++
			}
			saved[dest-i-2]++
		}
	}

	for i, poi := range p {
		canCheat(poi, -1, 0, i)
		canCheat(poi, 1, 0, i)
		canCheat(poi, 0, -1, i)
		canCheat(poi, 0, 1, i)
	}

	fmt.Println(saved)

	return res
}

// I think that I over complicated this a bit but whatever....
func path(lines []string) []point {
	points := make(map[point]bool, 0)
	var start point
	var end point

	for y, line := range lines {
		for x, r := range line {

			if r == '.' {
				points[point{y, x}] = true
			}

			if r == 'S' {
				start = point{y, x}
			}

			if r == 'E' {
				end = point{y, x}
			}
		}
	}

	ordered := make([]point, len(points)+2)
	ordered[0] = start
	i := 1

	checkPoint := func(p point, x int, y int) bool {
		p = point{p.y + y, p.x + x}

		if points[p] {
			delete(points, p)
			ordered[i] = p
			return true
		}

		return false
	}

	for len(points) > 0 {
		point := ordered[i-1]

		if checkPoint(point, -1, 0) {
			i++
			continue
		}

		if checkPoint(point, 1, 0) {
			i++
			continue
		}

		if checkPoint(point, 0, 1) {
			i++
			continue
		}

		if checkPoint(point, 0, -1) {
			i++
			continue
		}
	}

	ordered[len(ordered)-1] = end

	return ordered
}
