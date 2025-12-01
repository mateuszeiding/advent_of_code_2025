package utils

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func Read_to_matrix(day int, test bool) [][]string {
	file := open_file(day, test)

	defer file.Close()

	reader := bufio.NewReader(file)
	matrix := [][]string{}

	for {
		line, err := reader.ReadString('\n')

		if err != nil && err != io.EOF {
			log.Fatalf("error reading file: %v", err)
		}

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

func Read_lines(day int, test bool) []string {
	file := open_file(day, test)

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

func open_file(day int, test bool) *os.File {
	path := fmt.Sprintf("%s%02d", "./inputs/day_", day)

	if test {
		path = path + ".test"
	}

	path = path + ".txt"

	file, err := os.Open(path)

	if err != nil {
		log.Fatalf("Cannot read file: %v", err)
	}

	return file
}
