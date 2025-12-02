package solutions

import (
	"aoc/utils"
	"fmt"
	"math"
	"strconv"
)

func Day01Part01() {
	input := utils.ReadLines(1, false)

	pointer := 50
	pass := 0

	for _, line := range input {
		rotation, times := splitRotation(line)

		var result int
		if rotation == "R" {
			result = pointer + times
		} else {
			result = pointer - times
			if result < 0 {
				result = result + 100
			}
		}

		pointer = result % 100

		fmt.Printf("%s    %d\n", line, pointer)

		if pointer == 0 {
			pass++
		}
	}

	fmt.Println(pass)
}

func Day01Part02() {
	input := utils.ReadLines(1, false)

	pointer := 50
	pass := 0

	for _, line := range input {
		rotation, times := splitRotation(line)

		if times >= 100 {
			pass = pass + int(math.Floor(float64(times)/float64(100)))
			times = times % 100
		}

		var result int
		if rotation == "R" {
			result = pointer + times
			if result > 100 {
				pass++
			}
		} else {
			result = pointer - times
			if result < 0 {
				result = max(result, -result)
				result = 100 - result

				if pointer != 0 {
					pass++
				}
			}

		}

		pointer = result % 100

		if pointer == 0 {
			pass++
		}

	}

	fmt.Println(pass)
}

func splitRotation(line string) (string, int) {
	rotation := line[0:1]
	times, err := strconv.Atoi(line[1:])

	if err != nil {
		panic(err)
	}

	return rotation, times
}
