package utils

import "fmt"

func PrintMatrix(matrix [][]string) {
	for _, i := range matrix {
		for _, j := range i {
			fmt.Print(j)
		}
		fmt.Print("\n")
	}
}

func PrintLines(arr []string) {
	for _, i := range arr {
		fmt.Println(i)
	}
}
