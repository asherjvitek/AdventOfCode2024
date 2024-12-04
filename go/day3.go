package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func day3Part1(lines []string) int64 {

	result := int64(0)
	r, _ := regexp.Compile(`mul\([0-9]+,[0-9]+\)|do\(\)|don't\(\)`)
	r2, _ := regexp.Compile(`[0-9]+`)
	do := true

	for _, line := range lines {
		res := r.FindAllString(line, -1)

		for _, val := range res {
			fmt.Println(val)

			switch val {
			case "do()":
				do = true
				continue
			case "don't()":
				do = false
				continue
			default:
				// Nothing to see here...
			}

			if strings.Index(val, "mul") < 0 || !do {
				continue
			}

			nums := r2.FindAllString(val, -1)

			num1, _ := strconv.ParseInt(nums[0], 10, 64)
			num2, _ := strconv.ParseInt(nums[1], 10, 64)

			// fmt.Println(num1, num2)

			result += num1 * num2
		}
	}

	return result
}
