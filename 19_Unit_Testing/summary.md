# Unit Testing

## Apa itu software testing ?
Suatu proses menganalisis item perangkat lunak untuk mendeteksi perbedaan antara kondisi yang ada dan yang diperlukan serta untuk mengevaluasi fitur item perangkat lunak

## Tujuan Testing
* Mencegah regresi
* Tingkat kepercayaan dalam pemfaktoran ulang
* Tingkatkan desain kode
* Dokumentasi Kode
* Meningkatkan skala tim

## Level dari testing
1. Unit -> menguji bagian terkecil yang dapat diuji (operasi logis tunggal) dari suatu aplikasi melalui metode
2. Integrasi -> menguji modul atau sub sistem tertentu melalui aplikasi
3. UI -> (End to End) menguji interaksi antara keseluruhan melalui antarmuka pengguna

## Framework
Framework menyediakan alat dan struktur yang diperlukan untuk menjalankan pengujian secara efisien

## Structure 
Dua pola yang biasa digunakan:
* Sentralisasi -> file pengujian Anda di dalam folder pengujian
* Simpan file pengujian bersama dengan file produksi

## Runner
* Alat yang menjalankan file tes
* Gunakan mode tontonan -> (jika ada perubahan pada basis kode, otomatis jalankan tes lagi, biar lebih efisien)
* Pilihlah pelari yang dapat berlari paling cepat

## Mocking
Kebutuhan Anda untuk membuat objek tiruan (dan) objek palsu yang mensimulasikan perilaku objek nyata

## Coverage
Pembuat kode perlu memastikan apakah mereka telah membuat pengujian yang cukup. Alat cakupan menganalisis kode aplikasi saat pengujian sedang berjalan.

### Coverange Report 
* Function
* Statment
* Beranch
* Lines

### Coverage Format 
* CLI
* XML
* HTML
* LCOV

## Cara untuk melakukan test
* Buat file pengujian baru di go -> nama berakhiran library_test.go
* Tulis fungsi pengujian -> Nama : Tes & Kata dengan huruf kapital

Pahami apa yang baru saja Anda buat:
* Lebih dari satu fungsi
* Ikuti aturan pengujian saat bepergian

Run Testing
```
go test ./lib/calculate -cover
```

Run Testing with report Coverage
```
go test ./lib/calculate -coverprofil=cover.out && go tool cover -html=cover.out
```
