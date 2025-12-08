package solutions

import (
	"aoc/utils"
	"fmt"
)

func Day07Part01() {
	input := utils.ReadToMatrix(7, false)

	startIndex := utils.IndexOf(input[0], "S")
	beamsXAxis := map[int]bool{}
	beamsXAxis[startIndex] = true
	splits := 0

	for _, v := range input[1:] {
		newAxis := map[int]bool{}

		for x := range beamsXAxis {
			if v[x] == "." {
				newAxis[x] = true
			}
			if v[x] == "^" {
				splits++
				newAxis[x-1] = true
				newAxis[x+1] = true
			}
		}

		beamsXAxis = newAxis
	}

	fmt.Println(splits)
}

func Day07Part02() {
	input := utils.ReadToMatrix(7, false)

	startIndex := utils.IndexOf(input[0], "S")
	beamsXAxis := map[int]int{}
	beamsXAxis[startIndex] = 1
	splits := 1

	for _, v := range input[1:] {
		newAxis := map[int]int{}

		for x, val := range beamsXAxis {
			if v[x] == "." {
				newAxis[x] += val
			}
			if v[x] == "^" {
				splits += val
				newAxis[x-1] += val
				newAxis[x+1] += val
			}
		}

		beamsXAxis = newAxis
	}

	fmt.Println(splits)
}
