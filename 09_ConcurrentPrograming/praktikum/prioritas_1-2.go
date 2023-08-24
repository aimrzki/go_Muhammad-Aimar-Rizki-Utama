package main

import (
	"fmt"
	"sync"
)

func printMultiplesOfThree(ch chan<- int, wg *sync.WaitGroup, n int) {
	defer wg.Done()

	for i := 1; i <= n; i++ {
		if i%3 == 0 {
			ch <- i
		}
	}

	close(ch)
}

func main() {
	n := 30 // Jumlah bilangan kelipatan 3 yang ingin dicetak
	ch := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go printMultiplesOfThree(ch, &wg, n)

	go func() {
		wg.Wait()
		close(ch)
	}()

	for num := range ch {
		fmt.Println(num)
	}
}
