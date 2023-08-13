package main

import "fmt"

func main() {
	var base1, base2, tinggi, luas int
	fmt.Println("Base 1 : ")
	fmt.Scan(&base1)
	fmt.Println("Base 2 : ")
	fmt.Scan(&base2)
	fmt.Println("Tinggi : ")
	fmt.Scan(&tinggi)
	luas = (base1 + base2) * tinggi / 2
	fmt.Println("Luas Trapesium : ", luas)
}
