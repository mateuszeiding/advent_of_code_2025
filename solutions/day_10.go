package solutions

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

type SettingsManual struct {
	expected int16
	buttons  []int16
}

func (m *SettingsManual) readInputForSettings(input []string) {
	battery := input[0]
	input = input[1 : len(input)-1]

	for i, part := range battery[1:] {
		if part == '#' {
			m.expected = m.expected | 1<<i
		}
	}

	for _, val := range input {
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

func (m *SettingsManual) findShortestSetting() int {
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

		manual := SettingsManual{}

		manual.readInputForSettings(fields)
		result += manual.findShortestSetting()
	}

	fmt.Println(result)
}

type JoltageManual struct {
	buttons [][]int
	joltage [12]int
}

func (m *JoltageManual) readInputForJoltage(input []string) {
	for _, val := range input[1 : len(input)-1] {
		button := []int{}
		splt := strings.Split(val[1:len(val)-1], ",")

		for _, v := range splt {
			conv, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}

			button = append(button, conv)
		}

		m.buttons = append(m.buttons, button)
	}

	joltage := input[len(input)-1:][0]
	for i, v := range strings.Split(joltage[1:len(joltage)-1], ",") {
		conv, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}

		m.joltage[i] = conv
	}
}

func (m *JoltageManual) findShortestJoltage() int {
	iter := 0

	currentSettings := map[[12]int][12]int{}
	currentSettings[[12]int{}] = [12]int{}

	for {
		iter++
		newSettings := map[[12]int][12]int{}
		for _, sett := range currentSettings {

			for _, button := range m.buttons {
				settClone := sett
				increaseOnIndexes(&settClone, &button)

				if isSameAsRequired(&m.joltage, &settClone) {
					return iter
				}

				if validForReq(&m.joltage, &settClone) {
					if _, ok := newSettings[settClone]; !ok {
						newSettings[settClone] = settClone
					}
				}
			}
		}

		currentSettings = newSettings

		if len(currentSettings) == 0 {
			panic("AAAAAA")
		}
	}
}

func increaseOnIndexes(dest *[12]int, indexes *[]int) {
	for _, v := range *indexes {
		dest[v]++
	}
}

func isSameAsRequired(requ, toValidate *[12]int) bool {
	for i, v := range requ {
		if toValidate[i] != v {
			return false
		}
	}

	return true
}

func validForReq(requ, toValidate *[12]int) bool {
	for i, v := range requ {
		if toValidate[i] > v {
			return false
		}
	}

	return true
}

func Day10Part02() {
	input := utils.ReadLines(10, false)
	result := 0

	for _, l := range input {
		fields := strings.Fields(l)

		manual := JoltageManual{}

		manual.readInputForJoltage(fields)
		result += manual.findShortestJoltage()
	}

	fmt.Println(result)
}
