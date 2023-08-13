package main

import "fmt"

func main() {
	rows := 5

	for i := 1; i <= rows; i++ {
		//Menambah spasi
		for j := 1; j <= rows-i; j++ {
			fmt.Print(" ")
		}
		//Mencetak bintang
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
