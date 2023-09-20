package main

import (
	"testing"
)

func TestCalculator(t *testing.T) {
	k := Kalkulator{}

	// Pengujian penjumlahan
	result := k.Addition(5, 3)
	if result != 8 {
		t.Errorf("Seharusnya : 8, Hasil Run :  %f", result)
	}

	// Pengujian pengurangan
	result = k.Subtraction(5, 3)
	if result != 2 {
		t.Errorf("Seharusnya : 2, Hasil Run :  %f", result)
	}

	// Pengujian pembagian
	result, err := k.Division(6, 2)
	if err != nil || result != 3 {
		t.Errorf("Seharusnya :  3 tanpa error, tetapi hasilnya %f dengan error: %v", result, err)
	}

	// Pengujian pembagian oleh 0 (error case)
	_, err = k.Division(6, 0)
	if err == nil {
		t.Errorf("Seharusnya ada error saat membagi oleh 0, tetapi tidak ada error yang ditemukan")
	}

	// Pengujian perkalian
	result = k.Multiplication(5, 3)
	if result != 15 {
		t.Errorf("Seharusnya:  15, Hasil Run :  %f", result)
	}

	// Tambahkan pengujian untuk mencakup jalur alternatif dalam Division
	result, err = k.Division(3, 0)
	if err == nil {
		t.Errorf("Seharusnya ada error saat membagi oleh 0, tetapi tidak ada error yang ditemukan")
	}
}
