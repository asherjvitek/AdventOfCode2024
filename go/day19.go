package main

import (
	"fmt"
	"strings"
)

var memory = make(map[string]int)

func day19(lines []string) int {
	patt, maxLen, towels := initDay19(lines)
	results := 0

	fmt.Println(maxLen)

	for _, towel := range towels {
		fmt.Println(towel)
		// canMake(patt, maxLen, towel, &results)
		results += part2(towel, patt)
	}

	return results
}

func part2(design string, towels map[string]bool) int {
	if len(design) == 0 {
		return 1
	}

	if memory[design] > 0 {
		return memory[design]
	}

	result := 0
	for towel := range towels {
		if strings.Index(design, towel) == 0 {
			remain := design[len(towel):]
			result += part2(remain, towels)
		}
	}

	memory[design] = result

	return result
}

func canMake(patt map[string]bool, maxLen int, towel string, results *int) bool {

	if len(towel) == 0 {
		*results++
		fmt.Println(*results)
		return true
	}

	for i := maxLen; i >= 1; i-- {
		// for i := 1; i <= maxLen; i++ {
		if len(towel) < i {
			continue
		}

		pattern := towel[0:i]
		rest := towel[i:]

		if !patt[pattern] {
			continue
		}

		// rest = strings.Replace(rest, pattern, "", -1)

		if canMake(patt, maxLen, rest, results) {
			continue
		}
	}

	return false
}

func initDay19(lines []string) (map[string]bool, int, []string) {
	patt := make(map[string]bool)
	maxLen := 1
	towels := make([]string, 0)

	for i, line := range lines {
		if i == 0 {
			split := strings.Split(line, ", ")
			for _, item := range split {
				patt[item] = true

				if len(item) > maxLen {
					maxLen = len(item)
				}
			}
			continue
		}

		if len(line) == 0 {
			continue
		}

		towels = append(towels, line)
	}

	return patt, maxLen, towels
}
