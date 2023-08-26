package main

// Definisikan struktur kendaraan
type kendaraan struct {
	totalRoda       int
	kecepatanPerJam int
}

// Definisikan struktur mobil dengan mewarisi struktur kendaraan
type mobil struct {
	kendaraan
}

// Metode untuk membuat mobil berjalan dengan menambahkan kecepatan
func (m *mobil) berjalan() {
	m.tambahKecepatan(10)
}

// Metode untuk menambahkan kecepatan pada mobil
func (m *mobil) tambahKecepatan(kecepatanBaru int) {
	m.kecepatanPerJam += kecepatanBaru
}

func main() {
	// Inisialisasi mobil cepat
	mobilCepat := mobil{}

	// Mobil cepat berjalan tiga kali
	mobilCepat.berjalan()
	mobilCepat.berjalan()
	mobilCepat.berjalan()

	// Inisialisasi mobil lamban
	mobilLamban := mobil{}

	// Mobil lamban berjalan satu kali
	mobilLamban.berjalan()
}
