package main

import "fmt"

func Mapping(slice []string) map[string]int {
	counts := make(map[string]int)

	for _, str := range slice {
		counts[str]++
	}

	return counts
}

func main() {
	// Test cases
	fmt.Println(Mapping([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe"}))
	fmt.Println(Mapping([]string{"asd", "qwe", "asd"}))
	fmt.Println(Mapping([]string{}))
}
