package solutions

import (
	"aoc/utils"
	"fmt"
	"math"
	"strings"
)

func Day02Part01() {
	input := utils.ReadLines(2, false)
	input = strings.Split(input[0], ",")

	result := 0

	for _, v := range input {
		start, end := utils.GetStartAndEndOfRange(v)

		startLen := int(math.Log10(float64(start)))
		endLen := int(math.Log10(float64(end)))

		if startLen%2 == 0 && endLen%2 == 0 {
			continue
		}

		for id := start; id <= end; id++ {
			idLen := int(math.Log10(float64(id))) + 1
			if idLen%2 == 1 {
				continue
			}

			idHalfLen := idLen / 2

			mathThingi := int(math.Pow(10, float64(idHalfLen)))
			firstHalf := id / mathThingi
			secondHalf := id % mathThingi

			if firstHalf == secondHalf {
				result = result + id
			}
		}
	}

	fmt.Println(result)
}

func Day02Part02() {
	input := utils.ReadLines(2, false)
	input = strings.Split(input[0], ",")

	result := 0

	for _, v := range input {
		start, end := utils.GetStartAndEndOfRange(v)

		for id := start; id <= end; id++ {
			idLen := int(math.Log10(float64(id))) + 1

			divisibleBy := getDivisibleBy(idLen)

			for _, num := range divisibleBy {
				parts := splitBy(id, num)

				if isIdInvalid(parts) {
					result = result + id
					break
				}

			}
		}
	}

	fmt.Println(result)
}

// Part 02
func getDivisibleBy(idLen int) []int {
	divisibleBy := []int{}

	for num := 1; num <= idLen/2; num++ {
		if idLen%num == 0 {
			divisibleBy = append(divisibleBy, num)
		}
	}

	return divisibleBy
}

func splitBy(id, by int) []int {
	parts := []int{}
	rest := id

	for rest > 0 {
		var mathThingi int
		if mathThingi == 1 {
			mathThingi = 10
		} else {
			mathThingi = int(math.Pow(10, float64(by)))
		}

		parts = append(parts, rest%mathThingi)

		rest = rest / mathThingi
	}

	return parts
}

func isIdInvalid(parts []int) bool {
	sequence := parts[0]
	for _, p := range parts[1:] {
		if p != sequence {
			return false
		}
	}

	return true
}
