package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func main() {
	fmt.Print("Masukkan teks: ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	frequencies := make(map[rune]int)
	var wg sync.WaitGroup
	var mu sync.Mutex

	// memisahkan teks menjadi potongan-potongan kecil untuk pemrosesan paralel
	chunkSize := len(text) / 4
	chunks := make([]string, 0)
	for i := 0; i < len(text); i += chunkSize {
		end := i + chunkSize
		if end > len(text) {
			end = len(text)
		}
		chunks = append(chunks, text[i:end])
	}

	for _, chunk := range chunks {
		wg.Add(1)
		go func(chunk string) {
			defer wg.Done()
			processChunk(chunk, &mu, frequencies)
		}(chunk)
	}

	wg.Wait()

	for char, count := range frequencies {
		fmt.Printf("%c : %d\n", char, count)
	}
}

func processChunk(chunk string, mu *sync.Mutex, frequencies map[rune]int) {
	for _, char := range chunk {
		if isLetter(char) {
			mu.Lock()
			frequencies[char]++
			mu.Unlock()
		}
	}
}

func isLetter(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}
