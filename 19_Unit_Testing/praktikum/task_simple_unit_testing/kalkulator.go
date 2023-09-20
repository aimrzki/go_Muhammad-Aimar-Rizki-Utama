package main

import (
	"errors"
)

// Struktur untuk menyimpan fungsi-fungsi kalkulasi
type Kalkulator struct{}

// Penjumlahan
func (k Kalkulator) Addition(a, b float64) float64 {
	return a + b
}

// Pengurangan
func (k Kalkulator) Subtraction(a, b float64) float64 {
	return a - b
}

// Pembagian
func (k Kalkulator) Division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Pembagian oleh nol tidak diizinkan")
	}
	return a / b, nil
}

// Pengkalian
func (k Kalkulator) Multiplication(a, b float64) float64 {
	return a * b
}
