# Clean Code dalam Golang: Ringkasan

## Pendahuluan

Repository ini bertujuan untuk memperkenalkan konsep Clean Code (kode bersih) dan bagaimana konsep ini dapat diterapkan dengan efektif dalam pengembangan perangkat lunak menggunakan bahasa pemrograman Go (Golang). Clean Code bukan hanya tentang membuat kode yang berfungsi, tetapi juga tentang membuat kode yang mudah dimengerti, mudah dirawat, dan dapat diubah dengan lancar. Dalam ringkasan ini, kita akan menjelaskan prinsip-prinsip Clean Code dan memberikan contoh konkret dalam konteks bahasa Go.

## Prinsip-prinsip Clean Code

### 1. **Keterbacaan Kode Tinggi**

Kode harus mudah dibaca dan dimengerti oleh pengembang lain. Gunakan nama variabel, fungsi, dan metode yang deskriptif. Hindari singkatan yang ambigu.

### 2. **Fungsi yang Pendek dan Fokus**

Fungsi sebaiknya pendek dan fokus pada satu tugas. Jika fungsi melebihi sekitar 20 baris, pertimbangkan untuk memecahnya menjadi fungsi yang lebih kecil.

### 3. **Pengelompokan Logis**

Kelompokkan potongan kode yang berkaitan dalam blok atau fungsi. Jangan biarkan kode tercecer di berbagai tempat.

### 4. **Menghindari Duplikasi**

Duplikasi kode dapat memunculkan masalah perawatan di masa depan. Buatlah utilitas atau fungsi bantu untuk menghindari duplikasi.

### 5. **Komentar yang Bermakna**

Komentar sebaiknya menjelaskan "mengapa" dibandingkan "apa". Kode seharusnya bisa menjelaskan dirinya sendiri. Hindari komentar yang hanya mengulangi kode.

## Contoh Penerapan dalam Golang

### 1. **Penggunaan Nama Variabel yang Deskriptif**

```go
// Tidak disarankan
x := 42

// Disarankan
totalItems := 42

// Tidak disarankan
func HitungTotalHarga(items []Item) float64 {
    // kode yang panjang dan kompleks
}

// Disarankan
func HitungTotalHarga(items []Item) float64 {
    // kode yang fokus hanya pada perhitungan total harga
}

// Tidak disarankan
func HitungLuasPersegiPanjang(panjang, lebar float64) float64 {
    return panjang * lebar
}

func HitungKelilingPersegiPanjang(panjang, lebar float64) float64 {
    return 2 * (panjang + lebar)
}

// Disarankan
func HitungLuasPersegiPanjang(panjang, lebar float64) float64 {
    return panjang * lebar
}

func HitungKelilingPersegiPanjang(panjang, lebar float64) float64 {
    return 2 * (panjang + lebar)
}

```

Menerapkan prinsip-prinsip Clean Code dalam bahasa Go dapat menghasilkan kode yang lebih mudah dimengerti, lebih mudah dirawat, dan lebih fleksibel untuk perubahan di masa depan. Dengan mengedepankan keterbacaan, fokus, dan pengelompokan logis, kita dapat meningkatkan kualitas perangkat lunak yang kita kembangkan.
