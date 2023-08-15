package main

import (
	"fmt"
	"math"
)

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{9, 8, 9},
	}

	difference := 0

	size := len(matrix)

	for i, row := range matrix {
		difference += row[i] - row[size-i-1]
	}

	absoluteDifference := int(math.Abs(float64(difference)))
	fmt.Println("Absolute Difference:", absoluteDifference)
}
