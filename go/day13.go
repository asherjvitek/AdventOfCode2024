package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type clawGame struct {
	a     point
	b     point
	prize point
}

func day13(lines []string) int {

	res := 0
	game := clawGame{}
	for _, line := range lines {
		split := strings.Split(line, ": ")

		switch split[0] {
		case "Button A":
			game.a = getButton(split[1])
		case "Button B":
			game.b = getButton(split[1])
		case "Prize":
			game.prize = getGame(split[1], 10000000000000)
		case "":
			fmt.Println("------------------------")
			gameRes := runGame(game)
			fmt.Println(gameRes)
			res += gameRes
			fmt.Println("------------------------")
		}
	}

	return res

}

func runGame(game clawGame) int {

	prize := game.prize
	a := game.a
	b := game.b

    //this is not me... I got this from someone else...
    //math is not something that I am strong at
    //This is apparently solving liner equations or something like that.
	det := a.x*b.y - a.y*b.x
    wa := Abs((prize.x*b.y - prize.y*b.x)/det)
    wb := Abs((prize.x*a.y - prize.y*a.x)/det)

    fmt.Println(wa, wb)

    if (prize.x*b.y - prize.y*b.x)%det == 0 && (prize.x*a.y - prize.y*a.x)%det == 0{
        return wa * 3 + wb
    }

    //This was we...
	///old

	// bClicks := max(prize.x/b.x, prize.y/b.y)
	// aClicks := 0
	//
	// fmt.Println(bClicks)
	//
	// for {
	// 	x := bClicks*b.x + aClicks*a.x
	// 	y := bClicks*b.y + aClicks*a.y
	//
	// 	if prize.x == x && prize.y == y {
	// 		return bClicks*1 + aClicks*3
	// 	}
	//
	// 	if x > prize.x || y > prize.y {
	// 		//game over
	// 		if bClicks == 0 {
	// 			return 0
	// 		}
	//
	// 		bClicks--
	// 	} else {
	// 		aClicks++
	// 	}
	// }

	return 0
}

func Abs(i int) int {
    return int(math.Abs(float64(i)))
}

func max(a int, b int) int {

	if a > b {
		return a
	}

	return b
}

func getGame(str string, mod int) point {
	s := strings.Split(str, ", ")

	x, _ := strconv.Atoi(s[0][2:])
	y, _ := strconv.Atoi(s[1][2:])

	// fmt.Println(y, x)

	return point{y+mod, x+mod}
}

func getButton(str string) point {
	s := strings.Split(str, ", ")

	x, _ := strconv.Atoi(s[0][2:])
	y, _ := strconv.Atoi(s[1][2:])

	fmt.Println(y, x)

	return point{y, x}
}
