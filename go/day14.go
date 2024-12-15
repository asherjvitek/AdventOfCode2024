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

	maxY := 7
	maxX := 11
	iter := 5

	for i := 0; i < iter; i++ {
		for i, robot := range robots {
			robot.move(maxY, maxX)
			robots[i] = robot
		}
	}

	for _, robot := range robots {
		fmt.Println(robot.position)
	}

	q1Low := point{0, 0}
	q1High := point{y: (maxY - 1) / 2, x: (maxX-1)/2 - 1}

	q2Low := point{y: 0, x: maxX - ((maxX - 1) / 2)}
	q2High := point{y: (maxY - 1) / 2, x: maxX - 1}

	q3Low := point{maxY/2 + 2, 0}
	q3High := point{maxY - 1, maxX / 2}

	q4Low := point{y: maxY/2 + 2, x: maxX/2 + 1}
	q4High := point{maxY - 1, maxX - 1}

	fmt.Println(q1Low, q1High)
	fmt.Println(q2Low, q2High)
	fmt.Println(q3Low, q3High)
	fmt.Println(q4Low, q4High)

	q1 := 0
	q2 := 0
	q3 := 0
	q4 := 0

	for _, robot := range robots {
		q1 += robot.withinPoints(q1Low, q1High)
		q2 += robot.withinPoints(q2Low, q2High)
		q3 += robot.withinPoints(q3Low, q3High)
		q4 += robot.withinPoints(q4Low, q4High)
	}

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
