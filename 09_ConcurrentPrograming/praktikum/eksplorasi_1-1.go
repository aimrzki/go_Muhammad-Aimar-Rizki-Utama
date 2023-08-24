package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type Product struct {
	Title    string  `json:"title"`
	Price    float64 `json:"price"`
	Category string  `json:"category"`
}

func fetchProducts(url string, wg *sync.WaitGroup, ch chan<- Product) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	var products []Product
	err = json.NewDecoder(resp.Body).Decode(&products)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, p := range products {
		ch <- p
	}
}

func main() {
	url := "https://fakestoreapi.com/products"
	numWorkers := 5

	var wg sync.WaitGroup
	wg.Add(numWorkers)

	ch := make(chan Product)

	for i := 0; i < numWorkers; i++ {
		go fetchProducts(url, &wg, ch)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	fmt.Println("Product data")

	for p := range ch {
		fmt.Printf("title: %s\nprice: %.2f\ncategory: %s\n===\n", p.Title, p.Price, p.Category)
	}
}
