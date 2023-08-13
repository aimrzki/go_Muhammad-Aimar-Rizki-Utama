package main

import "fmt"

func main() {
	var number int
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&number)

	fmt.Printf("Faktor dari angka %d  adalah \n", number)
	for i := 1; i <= number; i++ {
		if number%i == 0 {
			fmt.Println(i)
		}
	}
}
