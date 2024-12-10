package main

import (
	"fmt"
	"strconv"
)

func day9(lines []string) int {
	res := 0
	input := []rune(lines[0])
	translated := make([]string, 0)

	for i, r := range input {
		index := i / 2
		id := strconv.Itoa(index)
		isId := i%2 == 0
		num, _ := strconv.Atoi(string(r))

		for i := 0; i < num; i++ {
			if isId {
				translated = append(translated, id)
			} else {
				translated = append(translated, ".")
			}
		}
	}

	fmt.Println(translated)

	for i := len(translated) - 1; i >= 0; i-- {
		if translated[i] != "." {
			start, end := getRangeBack(translated, i, translated[i])
			rLen := start - end
			// fmt.Println("i loop", translated[i], rLen, i)

			j := 0
			for ; j < i; j++ {
				if translated[j] == "." {
					s, e := getRangeForward(translated, j, ".")
					// fmt.Println("j loop", translated[i], rLen, s, e, e-s)

					if e-s >= rLen {
						for i := 0; i < rLen; i++ {
							translated[s+i] = translated[start-i]
							translated[start-i] = "."
							// fmt.Println(translated)
						}
						break
					}
				}
			}

            if j == i {
                i -= rLen - 1
            }
		}
	}

	for i, v := range translated {
		if v != "." {
			num, _ := strconv.Atoi(v)
			res += num * i
		}
	}

	fmt.Println(translated)

	return res
}

func getRangeBack(s []string, i int, val string) (int, int) {
	start := i
	end := i - 1
	for end > 0 && s[end] == val {
		end--
	}

	return start, end
}

func getRangeForward(s []string, i int, val string) (int, int) {
	start := i
	end := i + 1
	for end < len(s) && s[end] == val {
		end++
	}

	return start, end
}

func day9Part1(lines []string) int {
	res := 0
	input := []rune(lines[0])
	translated := make([]string, 0)

	for i, r := range input {
		index := i / 2
		id := strconv.Itoa(index)
		isId := i%2 == 0
		num, _ := strconv.Atoi(string(r))

		for i := 0; i < num; i++ {
			if isId {
				translated = append(translated, id)
			} else {
				translated = append(translated, ".")
			}
		}
	}

	fmt.Println(translated)
	j := 0

	for i := len(translated) - 1; i >= 0; i-- {
		if translated[i] != "." {
			for ; j < i; j++ {
				if translated[j] == "." {
					translated[j] = translated[i]
					translated[i] = "."
					break
				}
			}

		}
	}

	for i, v := range translated {
		if v != "." {
			num, _ := strconv.Atoi(v)
			res += num * i
		}
	}

	fmt.Println(translated)

	return res
}
