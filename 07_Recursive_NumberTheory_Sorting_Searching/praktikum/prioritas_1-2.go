package main

import (
	"fmt"
	"sort"
)

type pair struct {
	name  string // Nama item.
	count int    // Jumlah kemunculan item.
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
	sort.SliceStable(pairs, func(i, j int) bool {
		return pairs[i].count < pairs[j].count
	})
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
