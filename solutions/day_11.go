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
	fft  bool
	dac  bool
	curr string
	key  string
}

func (d Debugger) new(v string) Debugger {
	// newPath := d.path
	// newPath = append(newPath, v)
	new := Debugger{
		fft:  d.fft,
		dac:  d.dac,
		curr: v,
		key:  d.key + v,
		// path: newPath,
	}

	if !new.fft && new.curr == "fft" {
		new.fft = true
	}

	if !new.dac && new.curr == "dac" {
		new.dac = true
	}

	return new
}

func Day11Part02() {
	input := utils.ReadLinesDiffPt2Test(11, false)
	devicesMap := map[string][]string{}
	whereAmI := map[string]Debugger{}
	start := Debugger{
		fft:  false,
		dac:  false,
		curr: "svr",
		key:  "svr",
	}
	whereAmI[start.key] = start

	for _, l := range input {
		split := strings.Split(l, ":")
		key := split[0]
		values := strings.Fields(split[1])

		devicesMap[key] = values
	}

	done := []Debugger{}
	count := 0
	for len(whereAmI) != 0 {
		newPaths := map[string]Debugger{}
		for _, v := range whereAmI {
			if v.curr == "out" {
				done = append(done, v)
				continue
			}
			// fmt.Println(v)
			// fmt.Println(outputs)
			for _, pth := range devicesMap[v.curr] {
				new := v.new(pth)
				newPaths[new.key] = new
			}
		}
		// fmt.Println("---")

		fmt.Println(len(newPaths))
		whereAmI = newPaths
	}

	for _, v := range done {
		if v.fft && v.dac {
			count++
		}
	}

	fmt.Println(count)
}
