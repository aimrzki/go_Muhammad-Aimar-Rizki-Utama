package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	for i, j := 0, len(s)-1; i < j; {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	fmt.Print("Masukkan kata: ")

	// Membaca baris input secara lengkap
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	if isPalindrome(input) {
		fmt.Printf("captured : %s \nPalindrome", input)
	} else {
		fmt.Printf("captured : %s \nBukan Palindrome", input)
	}
}
