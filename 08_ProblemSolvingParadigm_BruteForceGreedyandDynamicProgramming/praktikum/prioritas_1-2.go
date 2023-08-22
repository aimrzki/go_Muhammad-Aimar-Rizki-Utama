package main

import (
	"fmt"
)

func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}

	triangle := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		triangle[i] = make([]int, i+1)
		triangle[i][0] = 1
		triangle[i][i] = 1

		for j := 1; j < i; j++ {
			triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
		}
	}

	return triangle
}

func printPascalTriangle(triangle [][]int) {
	numRows := len(triangle)
	maxDigits := len(fmt.Sprint(triangle[numRows-1][numRows-1]))

	for i := 0; i < numRows; i++ {
		padding := maxDigits * (numRows - i - 1)
		fmt.Printf("%*s", padding/2, "")

		for j := 0; j <= i; j++ {
			fmt.Printf("%-*d", maxDigits, triangle[i][j])
			if j < i {
				fmt.Print(" ")
			}
		}

		fmt.Println()
	}
}

func main() {
	numRows := 5 // Ganti dengan jumlah baris yang diinginkan
	triangle := generate(numRows)
	printPascalTriangle(triangle)
}
