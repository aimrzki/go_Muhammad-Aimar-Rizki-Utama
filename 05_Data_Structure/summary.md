# Data Structure
Struktur data adalah cara untuk menyimpan, mengelola, dan mengatur data dalam suatu program. Ini seperti wadah yang berisi tipe data yang berbeda, memungkinkan kita untuk melakukan operasi kompleks pada data dengan efisien.

ada beberapa contoh data structure yang ada dalam bahasa golang seperti :
* Array
* Slice
* Map
* Function

Mari kita bahas satu persatu 

## Array
Array adalah salah satu struktur data dasar dalam bahasa pemrograman, termasuk dalam Golang. Secara sederhana, array adalah kumpulan elemen dengan tipe data yang sama, yang diindeks oleh angka untuk memudahkan aksesnya.

Contoh Pendeklarasian array

```
var namaArray [ukuran]tipeData

```

Lalu untuk mengakses indeks yang ada dalam suatu array kita dapat menggunakan 

```
namaArray [ukuran]
```

## Slice
Slice adalah struktur data dinamis dalam Golang yang memungkinkan Anda membuat potongan fleksibel dari array. Ini adalah alternatif yang lebih kuat daripada array tradisional karena memungkinkan perluasan dan penyusutan ukuran dengan mudah.

Karakteristik Utama Slice :
* Dinamis : Ukuran slice tidak tetap dan dapat berubah seiring berjalannya program
* Referensi : Slice adalah referensi ke bagian dari array yang ada, bukan data yang sebenarnya.
* Sifat Sama dengan Array : Slice dapat diakses menggunakan indeks dan juga mendukung operasi seperti iterasi dan manipulasi elemen.
* Deklarasi Ringkas: Anda dapat mendeklarasikan slice tanpa harus menentukan ukuran awal.
* Lebih Fleksibel: Slice lebih cocok untuk kasus di mana ukuran koleksi data mungkin berubah seiring waktu.

Contoh Deklarasi dan Penggunaan Slice

```
var potongan []int

```

```
angka := [5]int{10, 20, 30, 40, 50}
potongan := angka[1:4] // Ini menghasilkan slice [20, 30, 40]

```

Membuat slice Baru

```
potonganBaru := make([]int, 3, 5) // Slice dengan kapasitas 5, tetapi panjang awal 3
```

## Map
Map adalah struktur data kunci-nilai dalam Golang yang memungkinkan Anda menghubungkan nilai-nilai dengan kunci-kunci unik. Ini serupa dengan kamus di dunia nyata, di mana Anda dapat mencari nilai berdasarkan kata kunci yang Anda miliki.

Karakteristik Utama Map :
* Kunci Unik: Setiap kunci dalam map adalah unik, tidak boleh ada kunci yang sama.
* Tidak Tetap: Ukuran map dapat berubah seiring berjalannya program.
* Operasi Efisien: Pencarian dan penambahan elemen dalam map dilakukan dengan kecepatan tinggi.
* Dinamis: Anda dapat menambahkan, mengubah, atau menghapus pasangan kunci-nilai dalam map.

Contoh deklarasi dan Penggunaan Map

```
var mahasiswa map[string]int // Kunci: string, Nilai: int

```

```
mahasiswa = map[string]int{
    "John": 25,
    "Alice": 22,
    "Bob": 24,
}
```

```
mahasiswa["Eva"] = 23 // Menambahkan pasangan kunci-nilai baru atau mengubah nilai Eva
```

```
delete(mahasiswa, "Alice") // Menghapus pasangan kunci-nilai dengan kunci "Alice"
```

```
umur, ada := mahasiswa["John"] // Mengecek apakah "John" ada dalam map
```

## Function
Function adalah blok kode yang dapat dipanggil untuk menjalankan tugas tertentu. Dalam Golang, function adalah unit dasar dalam pemrograman yang memungkinkan Anda memecah kode menjadi bagian-bagian yang terorganisir, memfasilitasi penggunaan kembali kode, dan menjaga struktur yang teratur.

Karakteristik Utama Function :
* Pemanggilan: Function dipanggil dengan nama fungsinya diikuti oleh tanda kurung, seperti namaFungsi().
* Argumen: Anda dapat memberikan argumen ke function untuk mengirim data ke dalamnya.
* Return Value: Function dapat mengembalikan nilai setelah selesai berjalan.
* Deklarasi: Function dideklarasikan dengan menyebutkan nama, parameter, tipe return value, dan tubuh fungsinya.
* Paket dan Sifat Publik: Function yang dimaksudkan untuk digunakan di luar paket harus diawali dengan huruf kapital agar dapat diakses dari luar.

Contoh Deklarasi dan Penggunaan Function :

```
func sapa(nama string) {
    fmt.Println("Halo, " + nama)
}
```

```
sapa("Alice") // Output: "Halo, Alice"

```

```
func tambah(a, b int) int {
    return a + b
}

```

```
func tambah(a, b int) int {
    return a + b
}
```

```
hasil := tambah(5, 3) // hasil = 8
```
