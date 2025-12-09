package solutions

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

func Day09Part01() {
	input := utils.ReadLines(9, false)

	points := []Point{}

	for _, l := range input {
		split := strings.Split(l, ",")

		x, err := strconv.Atoi(split[0])
		if err != nil {
			panic(err)
		}

		y, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}

		points = append(points, Point{
			x,
			y,
		})
	}

	biggest := 0
	for i, curr := range points[:len(points)-2] {
		for _, next := range points[i+1:] {
			a := curr.x - next.x
			b := curr.y - next.y

			a = max(a, -a) + 1
			b = max(b, -b) + 1
			area := a * b

			if area > biggest {
				biggest = area
			}
		}

	}

	fmt.Println(biggest)
}
