package main

import "fmt"

func main() {
	var (
		nilai int
		grade string
	)
	fmt.Print("Masukkan nilai: ")
	fmt.Scan(&nilai)
	if nilai <= 100 && nilai >= 80 {
		grade = "A"
	} else if nilai <= 79 && nilai >= 65 {
		grade = "B"
	} else if nilai <= 64 && nilai >= 50 {
		grade = "C"
	} else if nilai <= 49 && nilai >= 35 {
		grade = "D"
	} else if nilai >= 0 && nilai <= 34 {
		grade = "E"
	} else {
		grade = "Nilai Invalid"
	}
	fmt.Println(grade)
}
