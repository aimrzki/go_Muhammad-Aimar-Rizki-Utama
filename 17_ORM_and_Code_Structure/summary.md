# ORM & Code Structure

## ORM
Object-relational mapping (ORM,O/RM,and O/R mapping tool) in computer science is a programming technique for converting data between incompatible type systems using object-oriented programming languages

## Gorm
GORM adalah sebuah ORM (Object-Relational Mapping) library untuk bahasa pemrograman Go (Golang). ORM adalah alat yang digunakan untuk mempermudah interaksi antara aplikasi Go dengan database relasional, seperti MySQL, PostgreSQL, SQLite, dan lainnya. GORM membantu pengembang Go dalam melakukan operasi seperti membuat, membaca, memperbarui, dan menghapus data dari database tanpa perlu menulis kueri SQL secara manual.

Berikut beberapa fitur utama dan konsep yang terkait dengan GORM:

* Mapping Struktur Data: GORM memungkinkan Anda untuk mendefinisikan struktur data dalam Go sebagai struktur (struct), dan GORM akan secara otomatis membuat tabel dalam database yang sesuai dengan struktur tersebut. Ini disebut sebagai "Auto-Migration."

* CRUD Operations: Anda dapat menggunakan GORM untuk melakukan operasi dasar seperti membuat (Create), membaca (Read), memperbarui (Update), dan menghapus (Delete) data dalam database dengan mudah.

* Query Builder: GORM menyediakan query builder yang kuat untuk melakukan kueri ke database. Anda dapat membangun kueri menggunakan metode-metode yang disediakan oleh GORM tanpa harus menulis kueri SQL secara manual.

* Hubungan antar Tabel: GORM mendukung hubungan antar tabel dalam database, seperti hubungan satu-ke-banyak (one-to-many), banyak-ke-satu (many-to-one), dan banyak-ke-banyak (many-to-many). Anda dapat mendefinisikan hubungan ini dalam struktur data Anda.

* Validasi Data: GORM memungkinkan Anda untuk mendefinisikan validasi data pada struktur Anda sehingga Anda dapat memastikan bahwa data yang masuk ke database sesuai dengan aturan bisnis yang diinginkan.

* Hooks: GORM juga menyediakan hooks yang memungkinkan Anda menambahkan logika tambahan sebelum atau setelah operasi database tertentu dilakukan.

* Transaksi: Anda dapat menggunakan GORM untuk mengelola transaksi dalam database, memastikan bahwa operasi-operasi tertentu dalam transaksi dilakukan secara atomik.

* Dukungan untuk Banyak Jenis Database: GORM mendukung berbagai jenis database relasional populer, sehingga Anda dapat menggunakan GORM dengan berbagai macam database.

Contoh penggunaan GORM:

```
package main

import (
    "fmt"
    "gorm.io/gorm"
    "gorm.io/driver/sqlite"
)

type User struct {
    ID   uint
    Name string
    Age  int
}

func main() {
    // Inisialisasi koneksi ke database SQLite menggunakan GORM
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to database")
    }
    
    // Auto-Migrate untuk membuat tabel sesuai dengan struktur data User
    db.AutoMigrate(&User{})
    
    // Membuat user baru
    newUser := User{Name: "John Doe", Age: 30}
    db.Create(&newUser)
    
    // Membaca data user dari database
    var retrievedUser User
    db.First(&retrievedUser, 1)
    
    fmt.Println("User:", retrievedUser)
}

```

Ini adalah contoh sederhana penggunaan GORM untuk membuat, membaca, dan memperbarui data dalam database SQLite. GORM sangat berguna untuk mempermudah pengelolaan data dalam aplikasi Go yang berinteraksi dengan database relasional.


## ORM Advantages
* Lebih sedikit query berulang
* Secara otomatis mengambil data ke objek siap pakai
* Cara sederhana jika Anda ingin menyaring data
* Sebelum menyimpannya di database
* Beberapa memiliki fitur permintaan cache

## ORM Disadvantages
* Tambahkan lapisan dalam kode dan biaya proses overhead
* Muat data hubungan yang tidak diperlukan
* Kueri mentah yang kompleks bisa memakan waktu lama untuk ditulis dengan ORM (>10 tabel join)
* Fungsi SQL tertentu yang terkait dengan satu vendor mungkin tidak didukung atau tidak ada fungsi tertentu (MySQL : Force Index)

## Database Migration
Cara memperbarui versi database agar sesuai dengan perubahan versi aplikasi. Perubahan dapat berupa upgrade ke versi terbaru atau rollback ke versi sebelumnya.

## Why DB Migration 
* Kemudahan Update Database
* Kemudahan Rollback Database
* Lacak perubahan pada Struktur database
* Riwayat struktur basis data ditulis pada sebuah kode
* Selalu kompatibel dengan perubahan versi aplikasi

## DB Relation in GORM
* Belongs to
* Has one
* Has many
* Many to many

## Install Gorm Command & MySQL Driver

```
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
```
