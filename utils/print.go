package utils

import "fmt"

func Print_matrix(matrix [][]string) {
	for _, i := range matrix {
		for _, j := range i {
			fmt.Print(j)
		}
	}
}

func Print_lines(arr []string) {
	for _, i := range arr {
		fmt.Print(i)
	}
}
