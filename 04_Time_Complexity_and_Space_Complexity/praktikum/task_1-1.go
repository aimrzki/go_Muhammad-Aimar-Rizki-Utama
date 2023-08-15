package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	if n <= 1 {
		return n == 1
	}

	if n <= 3 {
		return true
	}

	if n%2 == 0 || n%3 == 0 {
		return false
	}

	sqrtN := int(math.Sqrt(float64(n)))

	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isPrime(1000000007)) // Menghasilkan true (bilangan prima)
	fmt.Println(isPrime(13))         // Menghasilkan true (bilangan prima)
	fmt.Println(isPrime(17))         // Menghasilkan true (bilangan prima)
	fmt.Println(isPrime(20))         // Menghasilkan false (bukan bilangan prima)
	fmt.Println(isPrime(35))         // Menghasilkan false (bukan bilangan prima)
}
