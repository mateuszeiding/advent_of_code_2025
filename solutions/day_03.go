package solutions

import (
	"aoc/utils"
	"fmt"
	"slices"
	"strconv"
)

func Day03Part01() {
	input := utils.ReadToMatrix(3, false)

	result := 0
	for _, battery := range input {
		cartesian := getCartesian(battery)
		slices.Sort(cartesian)
		slices.Reverse(cartesian)

		result = result + cartesian[0]
		fmt.Printf("%v \n", cartesian[0])

	}

	fmt.Println(result)
}

func Day03Part02() {
	input := utils.ReadToMatrix(3, true)

	result := 0
	for _, battery := range input {
	}
}

// Part 01
func getCartesian(arr []string) []int {
	cartesain := []int{}

	for ii, i := range arr {
		for _, j := range arr[ii+1:] {
			first, err := strconv.Atoi(i)
			if err != nil {
				panic(err)
			}
			second, err := strconv.Atoi(j)
			if err != nil {
				panic(err)
			}

			cartesain = append(cartesain, first*10+second)
		}
	}

	return cartesain
}
