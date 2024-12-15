package main

import (
	"fmt"
	"strconv"
	"strings"
)

type robot struct {
	position point
	velocity point
}

func (r *robot) move(maxY int, maxX int) {
	r.position = r.position.add(r.velocity)

	if r.position.y >= maxY {
		r.position.y -= maxY
	}

	if r.position.y < 0 {
		r.position.y += maxY
	}

	if r.position.x >= maxX {
		r.position.x -= maxX
	}

	if r.position.x < 0 {
		r.position.x += maxX
	}
}

func (r robot) withinPoints(p1 point, p2 point) int {
	if r.position.x >= p1.x && r.position.y >= p1.y && r.position.x <= p2.x && r.position.y <= p2.y {
		return 1
	}

	return 0
}

func day14(lines []string) int {
	robots := getRobots(lines)

	maxY := 103
	maxX := 101
	iter := 100

	for i := 0; i < iter; i++ {
		for i, robot := range robots {
			robot.move(maxY, maxX)
			robots[i] = robot
		}
	}

	xMid := maxX / 2
    yMid := maxY / 2

	fmt.Println("mids", xMid, yMid)

	q1 := 0
	q2 := 0
	q3 := 0
	q4 := 0

	for _, robot := range robots {
	    if (robot.position.x < xMid && robot.position.y < yMid) {
            q1++
        } else if (robot.position.x > xMid && robot.position.y < yMid) {
            q2++
        } else if (robot.position.x < xMid && robot.position.y > yMid) {
            q3++
        } else if (robot.position.x > xMid && robot.position.y > yMid) {
            q4++
        }
	}

	fmt.Println("--------------------")
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			count := 0

			for _, r := range robots {
				if r.position.equals(point{y, x}) {
					count++
				}
			}

			if x == xMid || yMid == y {
                fmt.Print(" ")
			} else if count > 0 {
				fmt.Print(count)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println("--------------------")

	fmt.Println(q1, q2, q3, q4)

	return q1 * q2 * q3 * q4
}

func getRobots(lines []string) []robot {
	robots := make([]robot, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, " ")

		points := strings.Split(parts[0][2:], ",")
		x, _ := strconv.Atoi(points[0])
		y, _ := strconv.Atoi(points[1])

		velocities := strings.Split(parts[1][2:], ",")
		v1, _ := strconv.Atoi(velocities[0])
		v2, _ := strconv.Atoi(velocities[1])

		robots[i] = robot{
			position: point{y, x},
			velocity: point{y: v2, x: v1},
		}
	}

	return robots
}
