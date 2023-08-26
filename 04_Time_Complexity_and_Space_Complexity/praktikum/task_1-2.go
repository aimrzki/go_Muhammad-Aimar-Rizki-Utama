package main

import (
	"fmt"
)

func pow(x, n int) int {
	// kondisi pertama: jika n adalah 0, maka langsung mengembalikan nilai 1.
	if n == 0 {
		return 1
	}
	// kondisi kedua: jika n genap, maka menghitung nilai pangkat
	if n%2 == 0 {
		// Menghitung nilai pangkat dengan memanggil fungsi pow
		temp := pow(x, n/2)
		// Mengembalikan hasil perhitungan pangkat
		return temp * temp
	}
	// kondisi ketiga: jika n ganjil, maka menghitung nilai pangkat
	temp := pow(x, (n-1)/2)
	// Mengembalikan hasil perhitungan pangkat .
	return x * temp * temp
}

func main() {
	fmt.Println(pow(2, 3)) // 8

	fmt.Println(pow(5, 3)) // 125

	fmt.Println(pow(10, 2)) // 100

	fmt.Println(pow(2, 5)) // 32

	fmt.Println(pow(7, 3)) // 343
}
