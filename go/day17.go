package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Registers struct {
	a int
	b int
	c int
}

func day17(lines []string) string {
	program := make([]int, 0)
	registers := Registers{}
	res := make([]int, 0)
	ip := 0

	for _, line := range lines {
		split := strings.Split(line, ": ")

		switch split[0] {
		case "Register A":
			n, _ := strconv.Atoi(split[1])
			registers.a = n
			break
		case "Register B":
			n, _ := strconv.Atoi(split[1])
			registers.b = n
			break
		case "Register C":
			n, _ := strconv.Atoi(split[1])
			registers.c = n
			break
		case "Program":
			nums := strings.Split(split[1], ",")
			for _, num := range nums {
				n, _ := strconv.Atoi(num)
				program = append(program, n)
			}
			break
		}
	}

	// a := 0
	// for {
	// 	a++
		// if a > 117440 {
		// 	break
		// }

 //        if (a % 100000 == 0) {
 //            fmt.Println(a)
 //        }
	//
 //        ip = 0
	// 	registers.a = a
	// 	registers.b = 0
	// 	registers.c = 0
	// 	res = nil
	//
	// outer:
		for ip < len(program) {
			advance := true
			oper1 := program[ip]
			oper2 := program[ip+1]

			fmt.Println("ip:", ip, "oper1", getOper(oper1), "oper2", oper2, registers)

			switch oper1 {
			case 0:
				adv(oper2, &registers)
				break
			case 1:
				bxl(oper2, &registers)
				break
			case 2:
				bst(oper2, &registers)
				break
			case 3:
				advance = jnz(oper2, &registers, &ip)
				break
			case 4:
				bxc(&registers)
				break
			case 5:
				res = append(res, out(oper2, registers))
				// for i, r := range res {
				// 	if r != program[i] {
				// 		break outer
				// 	}
				// }
				break
			case 6:
				bdv(oper2, &registers)
				break
			case 7:
				cdv(oper2, &registers)
				break

			}

			if advance {
				ip += 2
			}
		}

	// 	if len(res) != len(program) {
	// 		continue
	// 	}
	//
	// 	for i, r := range res {
	// 		if r != program[i] {
	// 			continue
	// 		}
	// 	}
	//
 //        break
	//
	// }

	// fmt.Println("ip:", ip, registers, a)
	fmt.Println("ip:", ip, registers)
	results := make([]string, len(res))

	for i, n := range res {
		results[i] = strconv.Itoa(n)
	}
	return strings.Join(results, ",")
}

func getOper(oper1 int) string {
	switch oper1 {
	case 0:
		return "adv"
	case 1:
		return "bxl"
	case 2:
		return "bst"
	case 3:
		return "jnz"
	case 4:
		return "bxc"
	case 5:
		return "out"
	case 6:
		return "bdv"
	case 7:
		return "cdv"
	}

	return "WHAT?"
}

func adv(oper int, registers *Registers) {
	combo := float64(getComboValue(oper, *registers))
	denominator := math.Pow(2, combo)
	registers.a = registers.a / int(denominator)
}

func bxl(oper int, registers *Registers) {
	registers.b ^= oper
}

func bst(oper int, registers *Registers) {
	combo := getComboValue(oper, *registers)
	registers.b = combo % 8
}

func jnz(oper int, registers *Registers, ip *int) bool {
	if registers.a == 0 {
		return true
	}

	*ip = oper

	return false
}

func bxc(registers *Registers) {
	registers.b ^= registers.c
}

func out(oper int, registers Registers) int {
	combo := getComboValue(oper, registers)
	return combo % 8
}

func bdv(oper int, registers *Registers) {
	combo := float64(getComboValue(oper, *registers))
	denominator := math.Pow(2, combo)
	registers.b = registers.a / int(denominator)
}

func cdv(oper int, registers *Registers) {
	combo := float64(getComboValue(oper, *registers))
	denominator := math.Pow(2, combo)
	registers.c = registers.a / int(denominator)
}

func getComboValue(operand int, registers Registers) int {
	switch operand {
	case 0:
		return operand
	case 1:
		return operand
	case 2:
		return operand
	case 3:
		return operand
	case 4:
		return registers.a
	case 5:
		return registers.b
	case 6:
		return registers.c
	default:
		fmt.Println("This should not be valid")
		panic("WTF!")
	}
}
