package main

import (
	"fmt"
	"time"
)

func printMultiples(x int) {
	for i := 1; ; i++ {
		fmt.Printf("Kelipatan %d: %d\n", i, x*i)
		time.Sleep(3 * time.Second)
	}
}

func main() {
	var userInput int
	fmt.Print("Masukkan angka kelipatan yang diinginkan: ")
	fmt.Scan(&userInput)

	go printMultiples(userInput)
	time.Sleep(30 * time.Second)
}
