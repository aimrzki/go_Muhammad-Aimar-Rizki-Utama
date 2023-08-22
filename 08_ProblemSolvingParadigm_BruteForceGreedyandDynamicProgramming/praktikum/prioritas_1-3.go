package main

import "fmt"

var memo map[int]int

func fibonacci(number int) int {
	if val, ok := memo[number]; ok {
		return val
	}

	if number <= 1 {
		memo[number] = number
		return number
	}

	result := fibonacci(number-1) + fibonacci(number-2)
	memo[number] = result
	return result
}

func main() {
	memo = make(map[int]int)

	fmt.Println(fibonacci(0))  // 0
	fmt.Println(fibonacci(2))  // 1
	fmt.Println(fibonacci(9))  // 34
	fmt.Println(fibonacci(10)) // 55
	fmt.Println(fibonacci(12)) // 144
}
