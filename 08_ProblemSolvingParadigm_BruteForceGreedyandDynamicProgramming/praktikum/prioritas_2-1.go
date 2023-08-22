package main

import (
	"fmt"
	"math"
)

func Frog(jumps []int) int {
	n := len(jumps)
	if n <= 2 {
		return 0
	}

	// Membuat Slice untuk menyimpan biaya minimum untuk mencapai setiap batu
	dp := make([]int, n)
	dp[0] = 0
	dp[1] = int(math.Abs(float64(jumps[1] - jumps[0])))

	for i := 2; i < n; i++ {
		// Hitung biaya melompat dari batu sebelumnya
		// dan biaya lompat dari batu sebelum yang sebelumnya
		dp[i] = min(dp[i-1]+int(math.Abs(float64(jumps[i]-jumps[i-1]))), dp[i-2]+int(math.Abs(float64(jumps[i]-jumps[i-2]))))
	}

	return dp[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20}))         // 30
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50})) // 40
}
