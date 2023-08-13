package main

import "fmt"

func main() {
	var nilai int
	fmt.Println("Masukan Nilai : ")
	fmt.Scan(&nilai)
	if nilai%2 == 0 {
		fmt.Println("Bilangan Genap")
	} else {
		fmt.Println("Bilangan Ganjil")
	}
}
