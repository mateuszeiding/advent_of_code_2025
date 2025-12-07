package solutions

import (
	"aoc/utils"
	"fmt"
	"slices"
	"strconv"
)

func Day05Part01() {
	input := utils.ReadLines(5, false)
	emptyLineIndex := slices.Index(input, "")
	ranges := input[:emptyLineIndex]
	ids := input[emptyLineIndex+1:]

	result := 0

	for _, id := range ids {
	out:
		for _, r := range ranges {
			start, end := utils.GetStartAndEndOfRange(r)
			idNum, err := strconv.Atoi(id)
			if err != nil {
				panic(err)
			}

			if idNum >= start && idNum <= end {
				result++

				break out
			}

		}

	}

	fmt.Println(result)
}
