package main

import (
	"fmt"
)

var romanNumerals = []struct {
	value  int
	symbol string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func intToRoman(num int) string {
	roman := ""
	for _, pair := range romanNumerals {
		for num >= pair.value {
			roman += pair.symbol
			num -= pair.value
		}
	}
	return roman
}

func main() {
	inputs := []int{4, 9, 23, 2021, 1646}

	for _, num := range inputs {
		roman := intToRoman(num)
		fmt.Printf("Input: %d\nOutput: %s\n\n", num, roman)
	}
}
