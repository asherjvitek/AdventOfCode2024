package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

func day11(lines []string) int {
	line := lines[0]
	split := strings.Split(line, " ")
	nums := make([]int, len(split))

	for i, v := range split {
		n, _ := strconv.Atoi(v)
		nums[i] = n
	}

	fmt.Println(line)

	for iter := 0; iter < 40; iter++ {
		// fmt.Println("Started", iter)

		for i := 0; i < len(nums); i++ {
			n := nums[i]

			l := 1

			if n != 0 {
				l = int(math.Log10(float64(n)) + 1)
			}

			if n == 0 {
				nums[i] = 1
			} else if l%2 == 0 {
				index := l / 2
				num := strconv.Itoa(n)

				n1, _ := strconv.Atoi(num[0:index])
				n2, _ := strconv.Atoi(num[index : index+index])

				nums[i] = n1
				if i == len(nums)-1 {
					nums = append(nums, n2)
				} else {
					nums = slices.Insert(nums, i+1, n2)
				}
				i++

			} else {
				nums[i] = n * 2024
			}
		}
		fmt.Println(nums[0], len(nums), iter)

		// fmt.Println(strings.Join(updated, " "))

		// fmt.Println("Done", iter)
	}

	return len(nums)
}
