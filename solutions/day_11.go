package solutions

import (
	"aoc/utils"
	"fmt"
	"slices"
	"strings"
)

func Day11Part01() {
	input := utils.ReadLines(11, false)

	devicesMap := map[string][]string{}
	whereAmI := []string{"you"}

	for _, l := range input {
		split := strings.Split(l, ":")
		key := split[0]
		values := strings.Fields(split[1])

		devicesMap[key] = values
	}

	outCount := 0
	for len(whereAmI) != 0 {
		newKeys := []string{}
		for _, v := range whereAmI {
			outputs := devicesMap[v]

			if slices.Contains(outputs, "out") {
				outCount++
				continue
			}

			newKeys = append(newKeys, outputs...)
		}

		whereAmI = newKeys
	}

	fmt.Println(outCount)
}
