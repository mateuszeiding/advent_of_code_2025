package utils

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func ReadToMatrix(day int, test bool) [][]string {
	file := openFile(day, test, false)

	defer file.Close()

	reader := bufio.NewReader(file)
	matrix := [][]string{}

	for {
		line, err := reader.ReadString('\n')

		if err != nil && err != io.EOF {
			log.Fatalf("error reading file: %v", err)
		}

		line = strings.TrimSuffix(line, "\r\n")
		row := []string{}

		for _, ch := range line {
			row = append(row, string(ch))
		}

		matrix = append(matrix, row)

		if err == io.EOF {
			break
		}
	}

	return matrix
}

func ReadLines(day int, test bool) []string {
	file := openFile(day, test, false)

	return readLines(file)
}

func ReadLinesDiffPt2Test(day int, test bool) []string {
	file := openFile(day, test, true)

	return readLines(file)
}

func readLines(file *os.File) []string {
	defer file.Close()

	reader := bufio.NewReader(file)
	arr := []string{}

	for {
		line, err := reader.ReadString('\n')

		if err != nil && err != io.EOF {
			log.Fatalf("error reading file: %v", err)
		}

		line = strings.TrimSuffix(line, "\r\n")
		arr = append(arr, line)

		if err == io.EOF {
			break
		}
	}

	return arr
}

func openFile(day int, test bool, pt2 bool) *os.File {
	path := fmt.Sprintf("%s%02d", "./inputs/day_", day)

	if test {
		if pt2 {
			path = path + "_02"
		}

		path = path + ".test"
	}

	path = path + ".txt"

	file, err := os.Open(path)

	if err != nil {
		log.Fatalf("Cannot read file: %v", err)
	}

	return file
}
