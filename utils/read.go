package utils

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func Read_to_matrix(day int, test bool) [][]string {
	path := fmt.Sprintf("%s%02d", "./inputs/day_", day)

	if test {
		path = path + ".test"
	}

	path = path + ".txt"

	file, err := os.Open(path)

	if err != nil {
		log.Fatalf("Cannot read file: %v", err)
	}

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
