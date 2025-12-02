package solutions

import (
	"aoc/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Day02Part01() {
	input := utils.ReadLines(2, false)
	input = strings.Split(input[0], ",")

	result := 0

	for _, v := range input {
		start, end := getIdStartAndEnd(v)

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
	input := utils.ReadLines(2, true)
	input = strings.Split(input[0], ",")

	result := 0

	for _, v := range input {
		start, end := getIdStartAndEnd(v)

		startLen := int(math.Log10(float64(start)))
		endLen := int(math.Log10(float64(end)))

		if startLen%2 == 0 && endLen%2 == 0 {
			continue
		}

		for id := start; id <= end; id++ {
			idLen := int(math.Log10(float64(id))) + 1
			divisibleBy := []int{}

			for num := 2; num <= idLen; num++ {
				if idLen%num == 0 {
					divisibleBy = append(divisibleBy, num)
				}
			}

			parts := []int{}
			rest := id
			for _, num := range divisibleBy {
				for rest > 0 {
					mathThingi := int(math.Pow(10, float64(num)))
					if mathThingi == 1 {
						mathThingi = 10
					}
					parts = append(parts, rest%mathThingi)

					rest = rest / mathThingi
				}

				rest = id

				sequence := parts[0]
				invalidId := false
				for p := range parts {
					if sequence != p {
						invalidId = true
					}
				}

				parts = []int{}
				if invalidId {
					result = result + id
					break
				}

			}

		}
	}

	fmt.Println(result)
}

func getIdStartAndEnd(val string) (int, int) {
	idRange := strings.Split(val, "-")

	start, err := strconv.Atoi(idRange[0])
	if err != nil {
		panic(err)
	}

	end, err := strconv.Atoi(idRange[1])
	if err != nil {
		panic(err)
	}

	return start, end

}
