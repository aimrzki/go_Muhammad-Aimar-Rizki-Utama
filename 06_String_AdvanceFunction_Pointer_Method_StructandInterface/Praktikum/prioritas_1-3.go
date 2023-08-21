package main

import (
	"fmt"
	"strings"
)

func Compare(a, b string) string {
	if strings.HasPrefix(b, a) {
		return a
	}

	commonSubstring := ""
	for i := 0; i < len(a); i++ {
		if i < len(b) && a[i] == b[i] {
			commonSubstring += string(a[i])
		} else {
			break
		}
	}

	return commonSubstring
}

func main() {
	fmt.Println(Compare("AKA", "AKASHI"))     // AKA
	fmt.Println(Compare("KANGOORO", "KANG"))  // KANG
	fmt.Println(Compare("KI", "KIJANG"))      // KI
	fmt.Println(Compare("KUPU-KUPU", "KUPU")) // KUPU
	fmt.Println(Compare("ILALANG", "ILA"))    // ILA
}
