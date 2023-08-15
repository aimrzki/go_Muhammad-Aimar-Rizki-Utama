package main

import (
	"fmt"
	"math/big"
)

func isPrime(n *big.Int) bool {
	if n.Cmp(big.NewInt(2)) <= 0 {
		return n.Cmp(big.NewInt(1)) == 0
	}

	if n.Bit(0) == 0 {
		return false
	}

	sqrtN := new(big.Int).Sqrt(n)

	for i := big.NewInt(3); i.Cmp(sqrtN) <= 0; i.Add(i, big.NewInt(2)) {
		if new(big.Int).Mod(n, i).Cmp(big.NewInt(0)) == 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isPrime(big.NewInt(1000000007))) // Menghasilkan true (bilangan prima)
	fmt.Println(isPrime(big.NewInt(13)))         // Menghasilkan true (bilangan prima)
	fmt.Println(isPrime(big.NewInt(17)))         // Menghasilkan true (bilangan prima)
	fmt.Println(isPrime(big.NewInt(20)))         // Menghasilkan false (bukan bilangan prima)
	fmt.Println(isPrime(big.NewInt(35)))         // Menghasilkan false (bukan bilangan prima)
}
