package main

import (
	"strings"
)

type clawGame struct {
	a     point
	b     point
	prize point
}

func day13(lines []string) int {

    res := 0
    game :=  clawGame{}
	for _, line := range lines {
		split := strings.Split(line, ": ")

		switch split[0] {
		case "Button A":
			game.a = getButton(split[1])
		case "Button B":
			game.b = getButton(split[1])
		case "Prize":
			game.prize = getPrize(split[1])
		case "":
            res += runGame(game)
		}
	}

	return res

}

func runGame(game clawGame) int {
    return 0
}

func getButton(str string) point {

	return point{}
}

func getPrize(str string) point {

	return point{}
}
