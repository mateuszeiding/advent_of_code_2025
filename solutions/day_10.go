package solutions

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

type Manual struct {
	expected   int16
	joltageReq string
	buttons    []int16
}

func (m *Manual) readInput(input []string) {
	m.joltageReq = input[len(input)-1:][0]

	for i, part := range input[0] {
		if part == '#' {
			m.expected = m.expected | 1<<i
		}
	}

	for _, val := range input[1 : len(input)-1] {
		var newNum int16 = 0
		splt := strings.Split(val, "")

		for i := 1; i < len(splt)-1; i = i + 2 {
			conv, err := strconv.Atoi(splt[i])
			if err != nil {
				panic(err)
			}

			newNum = newNum | 1<<conv
		}

		m.buttons = append(m.buttons, newNum)
	}
}

func (m *Manual) findShortest() int {
	iter := 0

	currentSettings := []int16{0}

	for {
		iter++
		newSettings := []int16{}
		for _, sett := range currentSettings {

			for _, button := range m.buttons {
				result := sett ^ button

				if result == m.expected {
					return iter
				}

				newSettings = append(newSettings, result)

			}
		}

		currentSettings = newSettings
	}

}

func Day10Part01() {
	input := utils.ReadLines(10, false)
	result := 0

	for _, l := range input {
		fields := strings.Fields(l)

		manual := Manual{}

		manual.readInput(fields)
		result += manual.findShortest()
	}

	fmt.Println(result)
}
