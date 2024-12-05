package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func day5Part2(lines []string) int {
	res := 0
	rules := make(map[int][]int)

	i := 0

	for ; i < len(lines); i++ {
		line := lines[i]
		if len(strings.TrimSpace(line)) == 0 {
			//This is where we are changing to the pages
			i++
			break
		}

		split := strings.Split(line, "|")

		key, _ := strconv.Atoi(split[0])
		value, _ := strconv.Atoi(split[1])

		values := rules[key]
		values = append(values, value)
		rules[key] = values
	}

	for key, value := range rules {
		fmt.Println(key, value)
	}

	for ; i < len(lines); i++ {
		items := strings.Split(lines[i], ",")
		pages := make([]int, len(items), len(items))

		for j, item := range items {
			pages[j], _ = strconv.Atoi(item)
		}

        pi := 0
		valid := true
        for pi < len(pages) {
            
            page := pages[pi]
			values := rules[page]

			if len(values) > 0 {
				// fmt.Println(page, values)
				//We need to validate the order
				for _, val := range values {
					index := slices.Index(pages, val)
					if index != -1 && index < pi {
						valid = false
                        t := pages[index]
                        pages[index] = pages[pi]
                        pages[pi] = t
                        pi = index
                        continue
					}
				}
			}

            pi++
		}

		if !valid {
			// fmt.Println("page:", pages[len(pages)/2])
            fmt.Println(pages)
			res += pages[len(pages)/2]
		}
	}

	return res
}

func day5Part1(lines []string) int {
	res := 0
	rules := make(map[int][]int)

	i := 0

	for ; i < len(lines); i++ {
		line := lines[i]
		if len(strings.TrimSpace(line)) == 0 {
			//This is where we are changing to the pages
			i++
			break
		}

		split := strings.Split(line, "|")

		key, _ := strconv.Atoi(split[0])
		value, _ := strconv.Atoi(split[1])

		values := rules[key]
		values = append(values, value)
		rules[key] = values
	}

	for key, value := range rules {
		fmt.Println(key, value)
	}

	for ; i < len(lines); i++ {
		items := strings.Split(lines[i], ",")
		pages := make([]int, len(items), len(items))

		for j, item := range items {
			pages[j], _ = strconv.Atoi(item)
		}

		// fmt.Println(pages)

		// fmt.Println()

		valid := true
		for pi, page := range pages {
			if !valid {
				break
			}

			values := rules[page]

			if len(values) > 0 {
				// fmt.Println(page, values)
				//We need to validate the order
				for _, val := range values {
					index := slices.Index(pages, val)
					if index != -1 && index < pi {
						valid = false
						break
					}
				}

			}
		}

		if valid {
			// fmt.Println("page:", pages[len(pages)/2])
			res += pages[len(pages)/2]
		}
	}

	return res
}
