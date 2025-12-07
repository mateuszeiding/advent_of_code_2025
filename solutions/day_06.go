package solutions

import (
	"aoc/utils"
	"strconv"
	"strings"
)

func Day06Part01() {
	inputLines := utils.ReadLines(6, false)
	input := [][]string{}

	for i, val := range inputLines {
		input = append(input, []string{})
		input[i] = strings.Fields(val)
	}

	result := 0

	for x := range input[0] {
		answer := 0

		sign := input[len(input)-1][x]

		for y := 0; y < len(input)-1; y++ {
			num, err := strconv.Atoi(input[y][x])
			if err != nil {
				panic(err)
			}

			if y == 0 {
				answer = num
				continue
			}

			switch sign {
			case "*":
				answer = answer * num
			case "+":
				answer = answer + num
			}
		}

		result = result + answer
	}

	println(result)
}
