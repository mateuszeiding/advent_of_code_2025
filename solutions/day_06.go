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

func Day06Part02() {
	inputMx := utils.ReadToMatrix(6, false)
	input := flipWorksheet(inputMx)

	result := 0

	sign := ""
	answer := 0
	for i := range input {
		if len(strings.Fields(strings.Join(input[i], ""))) == 0 {
			result = result + answer
			answer = 0
			sign = ""
			continue
		}

		if sign == "" {
			sign = input[i][len(input[i])-1]
		}

		numStr := strings.Join(input[i], "")[:len(input[i])-1]
		num, err := strconv.Atoi(strings.TrimSpace(numStr))
		if err != nil {
			panic(err)
		}

		if answer == 0 {
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

	println(result)
}

func flipWorksheet(inputMx [][]string) [][]string {
	input := make([][]string, len(inputMx[0]))

	for i := range input {
		input[i] = make([]string, len(inputMx))
	}

	for i, val := range inputMx {
		for j, sign := range val {
			input[j][i] = sign
		}
	}

	return input
}
