package main

import "fmt"

func caesar(offset int, input string) string {
	result := ""

	for _, char := range input {
		// Jika karakter adalah huruf kecil
		if char >= 'a' && char <= 'z' {
			// Melakukan pergeseran sesuai offset
			shifted := 'a' + (char-'a'+rune(offset))%26
			result += string(shifted)
		} else {
			// Jika bukan huruf, langsung tambahkan ke hasil
			result += string(char)
		}
	}

	return result
}

func main() {
	fmt.Println(caesar(3, "abc"))                           // Output: def
	fmt.Println(caesar(2, "alta"))                          // Output: cnvc
	fmt.Println(caesar(10, "alterraacademy"))               // Output: kvdobbkkmknowi
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))    // Output: bcdefghijklmnopqrstuvwxyza
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz")) // Output: mnopqrstuvwxyzabcdefghijkl
}
