package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var lock sync.Mutex

	numWorkers := 4
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	var sum int

	for i := 0; i < numWorkers; i++ {
		go func(index int) {
			defer wg.Done()

			delay := time.Duration(index * 150)
			time.Sleep(time.Millisecond * delay)

			lock.Lock()
			sum += index * index // Menghitung kuadrat dari index
			lock.Unlock()
		}(i)
	}

	wg.Wait()

	fmt.Printf("Result: %d\n", sum)
}
