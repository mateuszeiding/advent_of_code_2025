package solutions

import (
	"aoc/utils"
	"fmt"
)

type IGrid interface {
	countAdjacent(this [][]string, x, y int) int
}

type Grid [][]string

func (input Grid) countAdjacent(yAxis, xAxis int) int {
	result := 0
	for y := yAxis - 1; y <= yAxis+1; y++ {
		for x := xAxis - 1; x <= xAxis+1; x++ {
			if y < 0 || y >= len(input) || x < 0 || x >= len(input[0]) {
				continue
			}

			if y == yAxis && x == xAxis {
				continue
			}

			if input[y][x] == "@" {
				result++
			}
		}
	}

	return result
}

func Day04Part01() {
	var input Grid = utils.ReadToMatrix(4, false)
	result := 0

	for y := range input {
		for x := range input[0] {
			if input[y][x] != "@" {
				continue
			}

			if input.countAdjacent(y, x) < 4 {
				result++
			}
		}
	}

	fmt.Println(result)

}

func Day04Part02() {
	var input Grid = utils.ReadToMatrix(4, false)
	result := 0
	repeatAddsResult := true

	for repeatAddsResult {
		repeatAddsResult = false

		for y := range input {
			for x := range input[0] {
				if input[y][x] != "@" {
					continue
				}

				if input.countAdjacent(y, x) < 4 {
					result++
					input[y][x] = "."

					repeatAddsResult = true
				}
			}
		}
	}

	fmt.Println(result)

}
