package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var lines []string

	file, err := os.Open("../inputs/3.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	file.Close()

	var zeroes, ones []int

	zeroes = make([]int, 12)
	ones = make([]int, 12)

	for _, line := range lines {
		for i, char := range line {
			if char == '0' {
				zeroes[i]++
			} else {
				if char == '1' {
					ones[i]++
				}
			}
		}
	}

	var final []string
	final = make([]string, 12)

	for i, z := range zeroes {
		if z > ones[i] {
			final[i] = "0"
		} else {
			final[i] = "1"
		}
	}

	rEpsilon := make([]string, 12)

	for i, v := range final {
		if v == "1" {
			rEpsilon[i] = "0"
		} else {
			rEpsilon[i] = "1"
		}
	}

	fmt.Println(strings.Join(rEpsilon, ""))
	fmt.Println(strings.Join(final, ""))

	epsilonInt, err := strconv.ParseInt(strings.Join(rEpsilon, ""), 2, 16)
	if err != nil {
		panic(err)
	}
	gammaInt, err := strconv.ParseInt(strings.Join(final, ""), 2, 16)
	if err != nil {
		panic(err)
	}

	fmt.Println(epsilonInt * gammaInt)

	// finished at 23:19
}
