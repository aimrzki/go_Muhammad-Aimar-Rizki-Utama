package main

import (
	"fmt"
)

func SimpleEquations(a, b, c int) {
	for x := 1; x <= 10000; x++ {
		for y := 1; y <= 10000; y++ {
			// Hitung nilai z berdasarkan a, x, dan y.
			z := a - x - y
			if z <= 0 {
				continue
			}

			// Jika x*y*z sama dengan b dan x^2 + y^2 + z^2 sama dengan c,
			// maka ini adalah solusi yang benar.
			if x*y*z == b && x*x+y*y+z*z == c {
				fmt.Printf("%d %d %d\n", x, y, z)
				return
			}
		}
	}

	fmt.Println("no solution")
}

func main() {
	SimpleEquations(1, 2, 3)  // Output: no solution
	SimpleEquations(6, 6, 14) // Output: 1 2 3
}
