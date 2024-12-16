package main

import "fmt"

const up = '^'
const right = '>'
const left = '<'
const down = 'v'
const wall = '#'
const box = 'O'
const robit = '@'
const empty = '.'

func day15(lines []string) int {
	warehouse := make([][]rune, 0)
	directions := make([]rune, 0)
	robot := point{0, 0}

	buildWarehouse := true

	for y, line := range lines {

		if len(line) == 0 {
			buildWarehouse = false
			continue
		}

		if buildWarehouse {
			l := make([]rune, len(line))
			for i, r := range line {
				l[i] = r
			}

			warehouse = append(warehouse, l)

			for x, r := range line {
				if r == robit {
					robot = point{y, x}
				}
			}
		} else {
			for _, r := range line {
				directions = append(directions, r)
			}
		}
	}

	printWarehouse(warehouse)

	fmt.Println(string(directions))

	for _, direction := range directions {
		printWarehouse(warehouse)
		fmt.Println("Move", string(direction))
		x := 0
		y := 0

		switch direction {
		case up:
			y = -1
		case right:
			x = 1
		case left:
			x = -1
		case down:
			y = 1
		}

		adacent1 := warehouse[robot.y+y][robot.x+x]

		if adacent1 == wall {
			continue
		}

		if adacent1 == empty {
			warehouse[robot.y][robot.x] = empty
			warehouse[robot.y+y][robot.x+x] = robit
            robot.y += y
            robot.x += x
            continue
		}

		adacent2 := warehouse[robot.y+y+y][robot.x+x+x]

		if adacent2 == wall {
			continue
		}

		if adacent1 == box && adacent2 == empty {
			warehouse[robot.y][robot.x] = empty
			warehouse[robot.y+y][robot.x+x] = robit
			warehouse[robot.y+y+y][robot.x+x+x] = box
            robot.y += y
            robot.x += x
		}

	}

    printWarehouse(warehouse)

	return 0
}

func printWarehouse(warehouse [][]rune) {
	fmt.Println("--------------------")
	for y := 0; y < len(warehouse); y++ {
		for x := 0; x < len(warehouse[y]); x++ {
			fmt.Print(string(warehouse[y][x]))
		}
		fmt.Println()
	}
	fmt.Println("--------------------")
}
