package solutions

import (
	"aoc/utils"
	"fmt"
	"strconv"
)

func Day_01_part_01() {
	input := utils.Read_lines(1, true)

	pointer := 50
	pass := 0

	for _, line := range input {
		rotation, times := split_rotation(line)

		if times >= 100 {
			times = times % 100
		}

		if rotation == "R" {
			pointer = (pointer + times) % 100
		} else {
			x := pointer - times + 100
			pointer = max(x, -x) % 100

		}

		if pointer == 0 {
			pass++
		}
	}

	fmt.Println(pass)
}

func split_rotation(line string) (string, int) {
	rotation := line[0:1]
	times, err := strconv.Atoi(line[1:])

	if err != nil {
		panic(err)
	}

	return rotation, times
}
