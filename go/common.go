package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func readLines(path string) ([]string, error) {
	f, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	var lines []string

	err = nil

	r := bufio.NewReader(f)

	for err == nil {
		var (
			line, ln     []byte
			isPrefix bool = true
		)

		for isPrefix && err == nil {
			ln, isPrefix, err = r.ReadLine()

            if err == nil {
                line = append(line, ln...)
            }
		}


        if err == nil {
            lines = append(lines, string(line))
        }
	}

	return lines, nil
}
