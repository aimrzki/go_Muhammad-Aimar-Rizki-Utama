package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := make(chan int, 10) // Buat channel dengan buffer size 10
	var wg sync.WaitGroup

	// Goroutine untuk mengirim bilangan kelipatan 3 ke channel
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 100; i++ {
			if i%3 == 0 {
				numbers <- i
			}
		}
		close(numbers)
	}()

	// Goroutine untuk mencetak bilangan dari channel
	wg.Add(1)
	go func() {
		defer wg.Done()
		for num := range numbers {
			fmt.Println(num)
		}
	}()

	wg.Wait() // Menunggu kedua goroutine selesai
}
