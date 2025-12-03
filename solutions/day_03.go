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

func Day03Part02() {
	input := utils.ReadToMatrix(3, false)

	result := 0
	for _, battery := range input {
		batterySlice := battery

		fmt.Println(strings.Join(battery, ""))
		iter := 0
		for len(batterySlice) > 12 {
			batterySlice = buildJoltage(batterySlice, iter)
			iter++
			// fmt.Printf("%v \n", batterySlice)
			fmt.Println(strings.Join(batterySlice, ""))
		}

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
func buildJoltage(arr []string, iter int) []string {
	biggestIndex := iter
	fmt.Printf("%d %d %d \n", len(arr), iter, 12+iter)
	if len(arr)-(12+iter) > iter {

		for i, v := range arr[iter : len(arr)-(12+iter)] {
			if v > arr[biggestIndex] {
				biggestIndex = i + iter
			}
		}
		fmt.Printf("biggest i: %d %s \n", biggestIndex, arr[biggestIndex])

		if iter == 0 {
			return arr[biggestIndex:]
		}
		if biggestIndex != iter {
			return append(arr[0:iter], arr[biggestIndex:]...)
		}
	}

	smallestIndex := 0

	for i, digit := range arr {
		if digit < arr[smallestIndex] {
			smallestIndex = i
		}
	}

	return append(arr[:smallestIndex], arr[smallestIndex+1:]...)
}
