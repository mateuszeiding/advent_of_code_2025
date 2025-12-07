package solutions

import (
	"aoc/utils"
	"fmt"
	"slices"
	"sort"
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

type Range struct {
	start, end int
}

func Day05Part02() {
	input := utils.ReadLines(5, false)
	emptyLineIndex := slices.Index(input, "")
	ranges := []Range{}
	for _, r := range input[:emptyLineIndex] {
		start, end := utils.GetStartAndEndOfRange(r)

		ranges = append(ranges, Range{start: start, end: end})
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].start <= ranges[j].start
	})

	rangesMerged := []Range{ranges[0]}

	for _, r := range ranges {
		last := rangesMerged[len(rangesMerged)-1]

		if r.start > last.end {
			rangesMerged = append(rangesMerged, r)
		}

		if last.end < r.end {
			rangesMerged[len(rangesMerged)-1].end = r.end
		}
	}

	result := 0
	for _, r := range rangesMerged {
		diff := r.start - r.end

		diffAbs := max(diff, -diff)

		result = result + diffAbs + 1
	}

	fmt.Println(result)
}
