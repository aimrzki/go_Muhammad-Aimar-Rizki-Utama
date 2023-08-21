package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {
	merged := append(arrayA, arrayB...)

	uniqueNames := make(map[string]bool)
	result := []string{}

	for _, name := range merged {
		if !uniqueNames[name] {
			uniqueNames[name] = true
			result = append(result, name)
		}
	}

	return result
}

func main() {
	// Test cases
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	fmt.Println(ArrayMerge([]string{"devil jin"}, []string{"devil jin", "sergei"}))
	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	fmt.Println(ArrayMerge([]string{}, []string{}))
}
