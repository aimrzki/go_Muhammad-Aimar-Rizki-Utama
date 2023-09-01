package main

import (
	"fmt"
)

func MaxSequence(arr []int) int {
	maxSum := 0
	currentSum := 0

	for _, num := range arr {
		currentSum += num

		// Jika nilai currentSum negatif, kita reset menjadi 0
		if currentSum < 0 {
			currentSum = 0
		}

		// Membandingkan dengan nilai maksimum yang ditemukan sejauh ini
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}

	return maxSum
}

func main() {
	fmt.Println(MaxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 6
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6}))   // 7
	fmt.Println(MaxSequence([]int{-2, -3, 4, -1, -2, 1, 5, -3}))   // 7
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 6, -6}))   // 8
	fmt.Println(MaxSequence([]int{-2, -5, 6, 2, -3, 1, 6, -6}))    // 12
}
