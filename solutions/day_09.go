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

func createPointFromLine(l string) Point {
	split := strings.Split(l, ",")

	x, err := strconv.Atoi(split[0])
	if err != nil {
		panic(err)
	}

	y, err := strconv.Atoi(split[1])
	if err != nil {
		panic(err)
	}

	return Point{
		x,
		y,
	}
}

func Day09Part01() {
	input := utils.ReadLines(9, false)

	points := []Point{}

	for _, l := range input {
		points = append(points, createPointFromLine(l))
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

func Day09Part02() {
	input := utils.ReadLines(9, true)

	points := []Point{}

	for _, l := range input {
		points = append(points, createPointFromLine(l))
	}

	biggest := 0
	for i, curr := range points[:len(points)-2] {
		for _, next := range points[i+1:] {
			printPoints(points)
			if isAnyPointInsideSquare(points, curr, next) {
				continue
			}

			a := curr.x - next.x
			b := curr.y - next.y

			a = max(a, -a) + 1
			b = max(b, -b) + 1
			area := a * b

			if area > biggest {
				biggest = area
			}

			fmt.Println("false")
			fmt.Println(biggest)

		}

	}

	fmt.Println(biggest)
}

func isAnyPointInsideSquare(points []Point, curr, next Point) bool {
	minX, maxX := min(curr.x, next.x), max(curr.x, next.x)
	minY, maxY := min(curr.y, next.y), max(curr.y, next.y)
	fmt.Println(minX, maxX, minY, maxY, curr, next)

	for _, p := range points {
		fmt.Println(p.x, p.y)
		if p.x > minX && p.x < maxX && p.y > minY && p.y < maxY {
			return true
		}
	}

	return false
}

func printPoints(points []Point) {
	hiY, hiX := 0, 0

	for _, p := range points {
		if p.x > hiX {
			hiX = p.x
		}
		if p.y > hiY {
			hiY = p.y
		}
	}

	hiY += 3
	hiX += 3

	print := make([][]string, hiY)
	printLine := make([]string, hiX)

	for i := range printLine {
		printLine[i] = "."
	}

	for i := range print {
		print[i] = make([]string, hiX)
		copy(print[i], printLine)
	}

	for _, p := range points {
		print[p.y][p.x] = "#"
	}

	utils.PrintMatrix(print)

}
