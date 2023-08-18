package main

import (
	"fmt"
)

type pair struct {
	name  string
	count int
}

func MostAppearItem(items []string) []pair {
	countMap := make(map[string]int)

	for _, item := range items {
		countMap[item]++
	}

	var pairs []pair
	for name, count := range countMap {
		pairs = append(pairs, pair{name: name, count: count})
	}

	return pairs
}

func main() {
	pairs := MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"})
	for _, list := range pairs {
		fmt.Print(list.name, " -> ", list.count, " ")
	}
	fmt.Println()

	pairs = MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"})
	for _, list := range pairs {
		fmt.Print(list.name, " -> ", list.count, " ")
	}
	fmt.Println()

	pairs = MostAppearItem([]string{"football", "basketball", "tenis"})
	for _, list := range pairs {
		fmt.Print(list.name, " -> ", list.count, " ")
	}
}
