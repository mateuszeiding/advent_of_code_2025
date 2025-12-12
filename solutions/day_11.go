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

type Debugger struct {
	fft      bool
	dac      bool
	curr     string
	quantity int
}

func (d Debugger) new(v string) Debugger {
	return Debugger{
		fft:      d.fft,
		dac:      d.dac,
		curr:     v,
		quantity: d.quantity,
	}
}

func Day11Part02() {
	input := utils.ReadLinesDiffPt2Test(11, false)
	devicesMap := map[string][]string{}
	whereAmI := map[string]Debugger{}
	whereAmI["svr"] = Debugger{
		fft:      false,
		dac:      false,
		curr:     "svr",
		quantity: 1,
	}

	for _, l := range input {
		split := strings.Split(l, ":")
		key := split[0]
		values := strings.Fields(split[1])

		devicesMap[key] = values
	}

	count := 0
	for len(whereAmI) != 0 {
		newPaths := map[string]Debugger{}
		for _, v := range whereAmI {
			outputs := devicesMap[v.curr]

			if slices.Contains(outputs, "fft") {
				v.fft = true
			}
			if slices.Contains(outputs, "dac") {
				v.dac = true
			}

			if slices.Contains(outputs, "out") {
				if v.fft && v.dac {
					count++
				}
				continue
			}

			for _, pth := range outputs {
				if entry, ok := newPaths[pth]; ok {
					entry.quantity++
					newPaths[pth] = entry
				} else {
					new := v.new(pth)
					newPaths[pth] = new
				}
			}
		}

		whereAmI = newPaths
	}

	fmt.Println(count)
}
