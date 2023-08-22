package main

import (
	"fmt"
	"strconv"
)

func generateBinaryNumbers(n int) []string {
	result := make([]string, n+1)
	for i := 0; i <= n; i++ {
		result[i] = strconv.FormatInt(int64(i), 2)
	}
	return result
}

func main() {
	var n int
	fmt.Print("Masukkan nilai n: ")
	fmt.Scan(&n)

	output := generateBinaryNumbers(n)
	fmt.Println(output)
}
