# Concurrency Programing in Golang

Outline :
* Big Search Website
* Different Sequential, Parallel, and Concurrent
* Goroutines
* Channels and Select
* Race Condition
* Fixing Data Races

## Big Search Website
"Big Search Website" adalah istilah yang dapat merujuk pada situs web besar yang menyediakan mesin pencarian dalam skala besar. Ini merujuk pada situs web yang memiliki infrastruktur yang kuat dan canggih untuk melakukan pencarian dan indeksasi data dari berbagai sumber di seluruh internet. Konsep ini didasarkan pada ide untuk mengumpulkan, mengindeks, dan menyajikan informasi kepada pengguna dalam waktu yang cepat dan efisien.

## Sequental VS Parallel VS Concurrent

1. Sequental Program
Pada Sequental Program, sebelum task baru dimulai task yang sebelumnya harus telah selesai lebih dahulu

2. Parallel Program 
Pada Parallel Programs, banyak task dapat dijalankan secara simultaneously (Mengharuskan mesin multi-Core)

3. Concurrent Program
Pada Concurrent Program, banyak tas yang dapat dijalankan secara indepent dan mungkin simultaneous


## Goroutine 
Goroutine adalah salah satu fitur kunci dalam bahasa pemrograman Go (Golang). Ini adalah cara yang efisien dan ringan untuk melakukan pemrograman konkuren atau paralel di dalam program kita. Goroutine memungkinkan untuk menjalankan fungsi atau blok kode secara independen, tanpa harus menunggu fungsi tersebut selesai sebelum melanjutkan eksekusi program utama.

Karakteristik Utama Goroutine :
* Ringan dan Efisien: Goroutine sangat ringan dibandingkan dengan thread pada umumnya. Golang mengelola goroutine menggunakan sedikit memori dan overhead, memungkinkan Anda membuat banyak goroutine dalam program tanpa khawatir akan beban yang besar.

* Independen: Setiap goroutine memiliki alur eksekusi sendiri dan berbagi memori dengan goroutine lain. Namun, Go menyediakan mekanisme sinkronisasi bawaan untuk menghindari masalah akses bersama yang tidak aman.

* Pembuatan Goroutine: Anda dapat membuat goroutine dengan menempatkan kata kunci go sebelum pemanggilan fungsi. Contohnya: go fungsiAnda().

* Komunikasi dan Sinkronisasi: Goroutine dapat berkomunikasi dan berkoordinasi melalui saluran (channels) yang disediakan oleh Golang. Ini adalah cara yang aman untuk bertukar data antara goroutine.

Example Goroutine

```
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Mulai program")

	go fungsiGoroutine("Goroutine 1")
	go fungsiGoroutine("Goroutine 2")

	time.Sleep(time.Second * 2)
	fmt.Println("Selesai program")
}

func fungsiGoroutine(nama string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(nama, ":", i)
		time.Sleep(time.Millisecond * 500)
	}
}

```

## Channel & Select

### Channel

Channel adalah struktur data yang digunakan untuk mengizinkan komunikasi dan koordinasi antara goroutine dalam program Go. Ini adalah mekanisme yang kuat untuk mengirim dan menerima data antar goroutine dengan aman, menghindari race condition dan masalah sinkronisasi.

Karakteristik Utama:

* Membuat Channel: Anda dapat membuat channel dengan menggunakan fungsi make dan menentukan tipe data yang akan dikirim melalui channel. Contohnya: ch := make(chan int).

* Mengirim dan Menerima Data: Goroutine dapat mengirim dan menerima data melalui channel dengan operator <-. Misalnya, ch <- data untuk mengirim data ke channel dan data <- ch untuk menerima data dari channel.

* Blocking: Operasi mengirim dan menerima pada channel bersifat blocking. Jika goroutine mencoba mengirim data ke channel yang penuh, ia akan terblokir hingga ada ruang kosong. Jika goroutine mencoba menerima dari channel yang kosong, ia akan terblokir hingga ada data yang tersedia.

* Sifat Synchronizing: Selain berfungsi sebagai saluran komunikasi, channel juga dapat digunakan untuk mensinkronkan goroutine. Misalnya, penggunaan channel dapat memastikan bahwa satu goroutine menunggu hingga tugas tertentu selesai oleh goroutine lain.


### Select
Select adalah konstruksi kontrol yang digunakan untuk memilih salah satu dari beberapa operasi channel yang siap dieksekusi. Ini memungkinkan Anda untuk melakukan operasi I/O non-blok, yang dapat membantu dalam menghindari blok goroutine saat menunggu operasi channel tertentu.

Karakteristik Utama:

* Pemilihan Operasi Channel: Dalam blok select, Anda dapat menentukan beberapa kasus (cases) yang merupakan operasi pada channel. Kasus ini dapat berupa mengirim data, menerima data, atau bahkan default (untuk menangani situasi ketika tidak ada operasi channel yang siap dieksekusi).

* Eksekusi Non-Blok: Select memeriksa semua kasus secara bergantian dan memilih satu kasus pertama yang siap dieksekusi. Jika tidak ada kasus yang siap, select akan memasuki mode non-blok dan tidak akan terblokir.

* Kasus Default: Anda dapat menyertakan kasus default dalam select. Jika tidak ada kasus lain yang siap dieksekusi, kasus default akan dieksekusi.

Example :

```
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(time.Second)
		ch1 <- 42
	}()

	go func() {
		time.Sleep(time.Second * 2)
		ch2 <- 100
	}()

	select {
	case val := <-ch1:
		fmt.Println("Data dari ch1:", val)
	case val := <-ch2:
		fmt.Println("Data dari ch2:", val)
	case <-time.After(time.Second * 3):
		fmt.Println("Timeout! Tidak ada data yang diterima.")
	}
}

```


Kesimpulan :

Channel dan Select adalah fitur penting dalam bahasa pemrograman Go yang memungkinkan komunikasi dan koordinasi yang aman antara goroutine. Channel digunakan untuk mengirim dan menerima data antar goroutine, sementara Select memungkinkan Anda untuk memilih operasi channel yang siap dieksekusi, menghindari blok dan memungkinkan pemrograman konkuren yang lebih efisien.

## Race condition
Race condition adalah situasi di mana hasil dari program bergantung pada urutan eksekusi dari operasi- operasi konkuren (concurrent) atau paralel. Terjadi ketika dua atau lebih goroutine mengakses dan memodifikasi data bersamaan tanpa sinkronisasi yang memadai, dan hasil akhir dari program menjadi tidak dapat diprediksi dan tidak konsisten.

Dalam bahasa pemrograman Go (Golang), yang menempatkan penekanan kuat pada konkurensi dan paralelisme, race condition adalah masalah yang perlu diwaspadai.

Contoh Penggunaan : 

```
package main

import (
	"fmt"
	"time"
)

var count = 0

func increment() {
	count = count + 1
}

func main() {
	for i := 0; i < 1000; i++ {
		go increment()
	}

	time.Sleep(time.Second)
	fmt.Println("Count:", count)
}

```
Dalam contoh ini, kita memiliki satu variabel count yang diakses dan dimodifikasi oleh banyak goroutine yang berjalan secara konkuren. Setiap goroutine mencoba meningkatkan count sebanyak satu. Namun, karena tidak ada mekanisme sinkronisasi, hasil akhir dari count tidak dapat diprediksi dengan benar. Anda mungkin akan melihat nilai yang berbeda setiap kali menjalankan program.

Mengatasi Race Condition:

* Untuk mengatasi race condition dalam Golang, Anda perlu menggunakan mekanisme sinkronisasi yang disediakan oleh bahasa ini. Beberapa pendekatan yang umum digunakan adalah:

* Mutex (Mutual Exclusion): Menggunakan sync.Mutex untuk mengunci (lock) bagian kode yang mengakses data bersamaan sehingga hanya satu goroutine yang dapat mengaksesnya pada suatu waktu.

* Channel: Menggunakan channel untuk mengirim data antar goroutine. Channel memastikan bahwa data hanya dapat diakses oleh satu goroutine pada suatu waktu.

* Atomic Operations: Menggunakan operasi atomik seperti atomic.AddInt64 untuk memodifikasi data secara atomik tanpa memerlukan mekanisme penguncian manual.

* Wait Groups: Menggunakan sync.WaitGroup untuk menunggu selesainya sejumlah goroutine sebelum melanjutkan eksekusi.

## Fixing Data Race 
Data race adalah situasi di mana dua atau lebih goroutine mengakses dan memodifikasi data bersamaan tanpa sinkronisasi yang memadai. Hal ini dapat mengakibatkan hasil yang tidak konsisten dan tidak dapat diprediksi dalam program. Memperbaiki data race sangat penting untuk memastikan program berjalan dengan benar dalam lingkungan konkuren dan paralel.

Berikut adalah langkah-langkah untuk memperbaiki data race dalam bahasa pemrograman Go:

* Pahami Lokasi Data Race:
Identifikasi bagian kode mana yang mengalami data race. Ini dapat ditemukan melalui laporan yang dihasilkan oleh alat-alat seperti go run -race atau go build -race. Jika ada bagian kode yang mengakses dan memodifikasi data bersamaan tanpa sinkronisasi, itu adalah tempat potensial untuk data race.

* Gunakan Sinkronisasi:
Gunakan mekanisme sinkronisasi seperti mutex, channel, atau operasi atomik untuk melindungi data bersamaan dari akses bersama yang tidak aman.

* Mutex (Mutual Exclusion):
Gunakan sync.Mutex untuk mengunci (lock) data saat diakses atau dimodifikasi. Pastikan hanya satu goroutine yang dapat mengakses data pada suatu waktu.

* Channel:
Menggunakan channel untuk mengirim dan menerima data antar goroutine. Channel secara inheren aman karena hanya satu goroutine yang dapat mengirim atau menerima data pada satu waktu.

* Operasi Atomik:
Dalam beberapa kasus, operasi atomik seperti atomic.AddInt64 dapat digunakan untuk memodifikasi data secara aman tanpa perlu mekanisme penguncian manual.

* Wait Groups:
Jika Anda memiliki sejumlah goroutine yang perlu menyelesaikan pekerjaan sebelum program berakhir, gunakan sync.WaitGroup untuk menunggu mereka selesai.\

* Testing:
Setelah melakukan perbaikan, gunakan alat-alat seperti go test atau go run -race untuk memastikan bahwa data race telah diperbaiki.

* Review Kode:
Selalu lakukan review kode dengan seksama dan pastikan bahwa mekanisme sinkronisasi digunakan secara konsisten di seluruh program.


Kesimpulan:

Memperbaiki data race sangat penting untuk memastikan program Go berjalan dengan benar dalam lingkungan konkuren dan paralel. Dengan menggunakan mekanisme sinkronisasi seperti mutex, channel, atau operasi atomik, Anda dapat melindungi data bersamaan dan menghindari data race yang dapat merusak konsistensi dan keandalan program Anda.

