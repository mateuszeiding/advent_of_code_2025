package utils

import "fmt"

func PrintMatrix(matrix [][]string) {
	for i, arr := range matrix {
		fmt.Printf("%03d: ", i)
		for _, v := range arr {
			fmt.Print(v)
		}
		fmt.Print("\n")
	}
}

func PrintLines[T comparable](arr []T) {
	for _, i := range arr {
		fmt.Println(i)
	}
}
