package solutions

import (
	"aoc/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
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

// 169181023305251
// 169241478792898
// 171052033625939
// 171120577466565
// 171297027032466

func Day03Part02() {
	input := utils.ReadToMatrix(3, true)

	result := 0
	for _, battery := range input {
		batterySlice := battery

		fmt.Println(strings.Join(battery, ""))

		batterySlice = buildJoltage(batterySlice)

		strNum := strings.Join(batterySlice, "")

		num, err := strconv.Atoi(strNum)
		if err != nil {
			panic(err)
		}
		fmt.Println(num)
		fmt.Println()
		result = result + num
	}

	fmt.Println(result)
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

// Part 02
func buildJoltage(arr []string) []string {
	iter := 0

	for len(arr) > 12 {
		biggestIndex := iter
		for i, v := range arr {
			if i <= iter {
				continue
			}
			if v > arr[biggestIndex] {
				biggestIndex = i
			}
		}

		if iter == 0 {
			arr = arr[biggestIndex:]
		} else if biggestIndex == len(arr)-1 {
			arr = arr[:12]
		} else {
			arr = append(arr[0:iter], arr[biggestIndex:]...)
		}
		fmt.Println(strings.Join(arr, ""))
		iter++
	}
	return arr
}
