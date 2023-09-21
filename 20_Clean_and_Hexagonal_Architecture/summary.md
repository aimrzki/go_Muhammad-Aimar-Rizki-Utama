# Clean dan Hexagonal Architecture dalam Golang

Clean dan Hexagonal Architecture adalah pola arsitektur yang digunakan dalam pengembangan perangkat lunak untuk menciptakan aplikasi yang mudah dipelihara, dapat diuji, dan dapat diskalakan. Di Golang, prinsip-prinsip ini dapat diterapkan untuk membangun kode yang kuat dan terorganisir.

## Clean Architecture

Clean Architecture, juga dikenal sebagai Onion Architecture, menekankan pemisahan masalah dalam aplikasi Anda. Ini terdiri dari lingkaran konsentris, dengan lingkaran terdalam mewakili logika bisnis inti dan lingkaran terluar mewakili ketergantungan eksternal.
Contoh Structure Project :
```
myapp/
├── entities/
│   └── user.go
├── usecases/
│   └── user_usecase.go
├── interfaces/
│   ├── controllers/
│   │   └── http_controller.go
│   └── repositories/
│       └── user_repository.go
├── main.go

```

### Konsep Utama:

1. **Entities**: Ini adalah objek bisnis inti yang menggabungkan aturan bisnis yang paling penting.

2. **Use Cases**: Use cases atau interactors berisi logika bisnis aplikasi. Mereka berperan sebagai perantara antara dunia eksternal dan logika bisnis inti.

3. **Interfaces**: Interfaces mendefinisikan kontrak untuk komunikasi antara berbagai lapisan arsitektur, memungkinkan penggantian komponen dengan mudah.

4. **Controllers/Adapters**: Mereka bertanggung jawab untuk menangani masukan eksternal (permintaan HTTP, perintah CLI, dll.) dan mengonversinya menjadi tindakan pada use cases.

5. **Frameworks/Drivers**: Lapisan terluar berurusan dengan kerangka kerja eksternal, perpustakaan, dan alat seperti database, server web, atau API pihak ketiga.

## Hexagonal Architecture

Hexagonal Architecture, juga dikenal sebagai Ports and Adapters, memperluas ide Clean Architecture. Ini menekankan ide port (interfaces) dan adapter (implementasi). Port mendefinisikan batas komunikasi aplikasi Anda, sementara adapter mengimplementasikan batas tersebut menggunakan kerangka kerja atau teknologi eksternal.

### Konsep Utama:

1. **Ports**: Ini adalah antarmuka melalui mana aplikasi Anda berkomunikasi dengan dunia eksternal. Port mendefinisikan metode atau operasi yang diungkapkan oleh aplikasi.

2. **Adapters**: Adapter mengimplementasikan antarmuka yang didefinisikan oleh port. Mereka mengadaptasi teknologi eksternal (seperti database, API, atau kerangka kerja UI) agar sesuai dengan arsitektur aplikasi.

# Konteks di Golang

Di Golang, paket `context` digunakan untuk mengelola siklus hidup operasi dan menyebar deadlines, pembatalan, dan nilai-nilai yang berhubungan dengan permintaan melintasi batas API. Ini sangat berguna dalam pemrograman mikro dan pemrograman konkuren.

### Konsep Utama:

- **Konteks**: Tipe `context.Context` ada di pusat paket `context`. Ini membawa batas waktu, pembatalan, dan nilai-nilai yang berhubungan dengan permintaan melintasi panggilan fungsi.

- **Konteks Awal**: Konteks awal adalah konteks akar dari mana semua konteks lainnya berasal. Ini tidak memiliki nilai atau batas waktu yang ditetapkan.

- **WithCancel dan WithDeadline**: Fungsi-fungsi ini memungkinkan Anda membuat konteks anak dengan pembatalan atau batas waktu.

- **Nilai Konteks**: Anda dapat melampirkan pasangan kunci-nilai ke konteks. Nilai-nilai ini dapat digunakan untuk mengirimkan data, seperti ID permintaan atau informasi pengguna, melintasi panggilan fungsi.

# Dari MVC ke Clean Code

Pola arsitektur Model-View-Controller (MVC) adalah pola arsitektur umum yang digunakan dalam pengembangan web. Berpindah dari MVC ke Clean Code melibatkan menerapkan prinsip-prinsip Clean Architecture untuk membuat kode Anda lebih mudah dipelihara dan diuji.

Langkah-langkah:

1. **Pisahkan Masalah**: Bagi kode Anda ke dalam lapisan yang berbeda: Model (Entities), Kontroler (Use Cases), dan Tampilan (interaksi UI).

2. **Gunakan Dependency Injection**: Alih-alih mengkodekan dependensi, sisipkan mereka ke dalam komponen Anda. Ini sesuai dengan Prinsip Inversi Dependensi dari Clean Architecture.

3. **Pindahkan Logika Bisnis**: Pindahkan logika bisnis dari kontroler ke komponen use case yang terpisah. Ini membuat kode lebih modular dan dapat diuji.

4. **Tentukan Antarmuka**: Gunakan antarmuka untuk mendefinisikan kontrak antara komponen. Ini memungkinkan penggantian implementasi dengan mudah, sesuai dengan Prinsip Inversi Dependensi.

5. **Pengembangan Berbasis Pengujian (TDD)**: Tulis pengujian unit untuk use case Anda dan komponen lainnya. Pastikan kode Anda diuji secara menyeluruh untuk menjaga kualitas kode.

# Menulis Unit Test di Golang

Pengujian unit di Golang penting untuk memastikan keandalan dan kemudahan pemeliharaan kode Anda. Kerangka kerja pengujian bawaan Go membuatnya mudah untuk menulis dan menjalankan pengujian unit.

Langkah-langkah untuk menulis pengujian unit di Golang:

1. **Buat Berkas Pengujian**: Buat berkas pengujian terpisah untuk setiap paket yang ingin diuji. Berkas pengujian harus memiliki sufiks `_test.go`.

2. **Impor Paket Pengujian**: Impor paket `testing` dalam berkas pengujian Anda.

3. **Tulis Fungsi Pengujian**: Tulis fungsi yang dimulai dengan kata "Test" diikuti dengan nama fungsi dan menerima parameter `*testing.T`. Fungsi-fungsi ini akan berisi kasus uji Anda.

4. **Gunakan Fungsi Pengujian**: Dalam fungsi pengujian Anda, gunakan fungsi pengujian seperti `t.Errorf` atau `t.Fatalf` untuk melaporkan kesalahan atau kegagalan.

5. **Jalankan Pengujian**: Gunakan perintah `go test` untuk menjalankan pengujian Anda. Golang akan secara otomatis menemukan dan menjalankan fungsi pengujian Anda.

Contoh unit testing golang :

```
package main

import (
    "testing"
)

func Add(a, b int) int {
    return a + b
}

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    expected := 5
    if result != expected {
        t.Errorf("Expected %d, but got %d", expected, result)
    }
}

```
