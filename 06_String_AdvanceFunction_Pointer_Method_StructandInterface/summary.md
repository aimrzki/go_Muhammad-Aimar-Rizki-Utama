# String, Advance Function, Pointer, Method, Struct and Interface

## Outline
* String
* Advance Function
* Pointer
* Package
* Error Handling

### String
String adalah tipe data yang digunakan untuk merepresentasikan urutan karakter. String dapat berisi karakter tunggal atau sejumlah karakter yang membentuk teks. Dalam Go, string direpresentasikan sebagai serangkaian karakter Unicode.

1. Cara pendeklarasian string adalah sebagai berikut :

```
var str1 string = "Hello, Golang!"
str2 := "Welcome to Go programming"

```

2. Panjang String
```
length := len(str1) // Mengembalikan nilai 14
```

3. Mengakses karakter

```
firstChar := str1[0] // Mendapatkan karakter pertama ("H")
```

4. Operasi Pada String :

```
concatenated := str1 + " " + str2 // Menghasilkan "Hello, Golang! Welcome to Go programming"
```

5. Iterasi Pada String
```
for i := 0; i < len(str1); i++ {
    // Lakukan sesuatu dengan str1[i]
}
```

6. String Bersifat Immutable 
Dalam Go, string bersifat immutable, artinya setelah string dibuat, Anda tidak dapat mengubah karakter-karakter individunya. Jika Anda perlu mengubah string, Anda sebenarnya membuat string baru.

### Variadic Option
variadic function adalah fungsi yang dapat menerima jumlah argumen yang bervariasi. Ini memungkinkan Anda untuk memanggil fungsi dengan jumlah argumen yang tidak tetap. Opsi ini sangat bermanfaat saat Anda ingin menangani sejumlah nilai tanpa harus membatasi jumlah argumen secara eksplisit. Dalam Go, variadic function diimplementasikan dengan menggunakan tanda elipsis (...) sebelum tipe parameter yang ingin digunakan untuk menerima argumen variadic.

1. Pendeklarasian Varadic Function

```
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

```

2. Pemanggilan Varadic Function
```
result := sum(1, 2, 3, 4, 5) // Memanggil fungsi sum dengan argumen variadic

```

3. Menggunakan Nilai slice
Dalam tubuh fungsi, parameter variadic dianggap sebagai slice, yang artinya Anda dapat memperlakukannya seperti slice pada umumnya, seperti melakukan iterasi menggunakan perulangan.

4. Penggunaan Kombinasi argumen

```
result1 := sum(1, 2, 3)       // Menghasilkan 6
result2 := sum(10, 20, 30, 40) // Menghasilkan 100

```

### Anonymous Function

Di dalam bahasa pemrograman Go (Golang), "anonymous function" mengacu pada fungsi yang tidak memiliki nama dan biasanya didefinisikan dalam lokasi yang sama di mana mereka digunakan. Anonymous function juga dikenal sebagai "lambda function" atau "closure" dalam beberapa bahasa pemrograman lain. Mereka memungkinkan Anda untuk membuat fungsi kecil atau sementara tanpa harus memberikan nama resmi ke fungsi tersebut.

Berikut adalah penjelasan lebih lanjut mengenai anonymous function dalam bahasa pemrograman Go:

1. Deklarasi Anonymus Function 

```
func main() {
    // Mendefinisikan dan memanggil anonymous function
    result := func(a, b int) int {
        return a + b
    }(3, 5)
    fmt.Println(result) // Output: 8
}

```

2. Penggunaan anonymous Function 

Anonymous function sering digunakan dalam beberapa situasi, seperti:
* Sebagai argumen untuk fungsi lain (callback functions).
* Ketika Anda ingin melakukan operasi sederhana tanpa perlu membuat fungsi dengan nama.
* Saat Anda ingin membuat "closure" untuk menyimpan state bersama dengan fungsi.

3. Closure 
Anonymous function sering digunakan untuk membuat "closure". Closure adalah sebuah fungsi yang tidak hanya mengandung kode fungsinya, tetapi juga menyimpan lingkup (scope) di mana fungsi tersebut didefinisikan. Ini memungkinkan closure untuk mengakses variabel-variabel yang didefinisikan di luar fungsi, bahkan setelah fungsi tersebut selesai dieksekusi.

Contoh : 

```
func makeCounter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

func main() {
    counter := makeCounter()
    fmt.Println(counter()) // Output: 1
    fmt.Println(counter()) // Output: 2
}

```

### Pointer 
Pointer adalah konsep penting dalam pemrograman yang juga ditemukan dalam bahasa pemrograman Go (Golang). Pointer memungkinkan Anda untuk mengakses dan mengubah nilai dari suatu variabel melalui referensi atau alamat memori variabel tersebut. Konsep ini memungkinkan Anda untuk melakukan manipulasi data secara efisien dan menghindari overhead dari menyalin data yang besar.

Di Golang, seperti bahasa pemrograman lainnya, setiap variabel memiliki lokasi memori yang terkait dengannya. Pointer adalah tipe data khusus yang dapat menunjuk ke lokasi memori suatu variabel. Saat Anda bekerja dengan pointer, Anda sebenarnya mengakses data asli yang disimpan di lokasi memori tersebut.

Berikut beberapa konsep penting mengenai pointer dalam Golang:

1. Deklarasi Pointer 
```
var x int = 42
var ptr *int // Deklarasi pointer ke tipe int
ptr = &x     // Menginisialisasi pointer dengan alamat x

```

2. Operasi Referensi Pointer
* Operator &: Digunakan untuk mendapatkan alamat memori dari suatu variabel.
* Operator *: Digunakan untuk mendapatkan nilai dari variabel yang ditunjuk oleh pointer (dereferensi).

3. Contoh Penggunaan

```
var x int = 42
var ptr *int
ptr = &x

fmt.Println("Nilai x:", x)       // Output: Nilai x: 42
fmt.Println("Alamat x:", &x)     // Menampilkan alamat memori x
fmt.Println("Nilai ptr:", *ptr)  // Output: Nilai ptr: 42 (dereferensi)

```

4.Penggunaan dalam fungsi :
Pointer sering digunakan dalam fungsi untuk mengubah nilai variabel di luar fungsi tersebut.

```
func modifyValue(ptr *int) {
    *ptr = 99
}

func main() {
    var x int = 42
    fmt.Println("Sebelum: ", x) // Output: Sebelum: 42
    modifyValue(&x)
    fmt.Println("Sesudah: ", x) // Output: Sesudah: 99
}

```

5. Keamanan dan Null Pointer 
Golang memiliki fitur keamanan pointer yang tinggi. Saat dideklarasikan, pointer akan diinisialisasi dengan nilai nil (kosong). Saat Anda berusaha mengakses nilai pada pointer nil, program tidak akan menjalankan operasi tersebut dan akan memunculkan kesalahan.
Pointer pada Golang memungkinkan Anda untuk mengelola dan berinteraksi dengan data secara lebih efisien, terutama saat Anda perlu berurusan dengan data yang besar dan kompleks. Namun, karena pointer juga dapat menimbulkan kesalahan jika tidak digunakan dengan benar, penting untuk memahami konsep ini dengan baik sebelum menggunakannya dalam program.

### Package 
Package (paket) adalah konsep fundamental dalam bahasa pemrograman Go (Golang) yang memungkinkan pemisahan dan pengelompokan kode program. Golang menggunakan sistem package untuk mengatur kode program ke dalam unit yang lebih terstruktur dan mudah dikelola. Dalam penjelasan ini, kita akan menggali lebih dalam tentang apa itu package dalam Golang.

1. Definisi Package : Package adalah cara untuk mengorganisir kode program di Golang. Setiap file Go harus dimasukkan ke dalam package tertentu. Package berisi satu atau lebih file yang terkait dan bekerja bersama untuk mencapai tujuan tertentu. Package juga memungkinkan untuk mengatur tingkat akses terhadap kode, artinya Anda dapat menentukan apakah kode dalam package tersebut dapat diakses dari package lain atau tidak.

2. Tujuan Package : 
* Organisasi Kode: Package memungkinkan untuk mengelompokkan dan mengatur kode program ke dalam unit yang lebih terstruktur. Ini membuat kode lebih mudah diorganisir, dipelihara, dan ditemukan.
* Penggunaan Ulang (Reusability): Anda dapat mengimpor package dari kode lain untuk menggunakan fungsionalitas yang telah ditentukan tanpa perlu menulis ulang kode yang sama.
* Penting untuk Modularitas: Package mendukung modularitas, yang memungkinkan untuk mengembangkan bagian-bagian kecil dari program secara independen dan kemudian menggabungkannya menjadi program yang lebih besar.

3. Structure Package 

Setiap package memiliki struktur yang mirip dengan direktori dalam sistem berkas. Dalam sebuah direktori package, Anda mungkin akan menemukan beberapa file Go dan mungkin satu file dengan nama main.go yang berfungsi sebagai entry point jika package tersebut dieksekusi sebagai program utama.

Contoh struktur package:

```
myapp/
|-- main.go
|-- utils/
|   |-- math.go
|   |-- string.go
|-- data/
|   |-- data.go

```

4. Penulisan Dalam Package 
* Penulisan Package: Pada awal file Go, Anda deklarasikan package dengan menyertakan pernyataan package nama_package. Ini menunjukkan bahwa kode dalam file ini adalah bagian dari package dengan nama tersebut.
* Pemanggilan Package: Untuk menggunakan kode dari package lain, Anda perlu mengimpornya menggunakan pernyataan import "nama_package". Setelah diimpor, Anda dapat menggunakan fungsi, tipe data, atau variabel yang didefinisikan dalam package tersebut.


### Error Handling

Error Handling dalam Bahasa Pemrograman Go (Golang)

Error handling adalah konsep penting dalam pemrograman yang memungkinkan program untuk mengatasi situasi yang tidak terduga atau kesalahan yang terjadi saat program dieksekusi. Bahasa pemrograman Go (Golang) memiliki pendekatan yang unik terhadap error handling yang didasarkan pada pengembalian nilai error sebagai bagian dari hasil dari suatu fungsi. Berikut ini penjelasan lebih rinci tentang error handling dalam Golang:

1. Error sebagai Tipe Data:
Di Golang, error diperlakukan sebagai tipe data. Tipe data error adalah interface yang didefinisikan sebagai berikut:
```
go
Copy code
type error interface {
    Error() string
}
```
Artinya, setiap tipe yang mengimplementasikan metode Error() sesuai dengan definisi di atas dianggap sebagai tipe error.

2. Fungsi yang Mengembalikan Error:
Banyak fungsi di Golang yang dapat mengembalikan nilai error sebagai bagian dari hasilnya. Ini memungkinkan program untuk mengevaluasi apakah operasi berhasil atau tidak dan melakukan tindakan yang sesuai.

Contoh fungsi yang mengembalikan error:
```
go
Copy code
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}
```
3. Menangani Error:
Untuk menangani error, Anda perlu memeriksa nilai error yang dikembalikan oleh fungsi. Anda dapat menggunakan pernyataan if atau switch untuk mengevaluasi error dan mengambil tindakan yang sesuai.

Contoh penanganan error:
```
go
Copy code
result, err := divide(10, 0)
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Println("Result:", result)
}
```

4. Error Customization:
Anda dapat membuat error yang lebih deskriptif dengan mengimplementasikan tipe custom yang memenuhi antarmuka error. Ini berguna untuk memberikan informasi yang lebih spesifik tentang jenis kesalahan yang terjadi.

Contoh error custom:

```
go
Copy code
type MyError struct {
    message string
}

func (e MyError) Error() string {
    return e.message
}

func doSomething() error {
    return MyError{message: "Something went wrong"}
}
```

5. Penggunaan Panic dan Recover:
Selain error handling dengan nilai error yang dihasilkan oleh fungsi, Golang juga memiliki mekanisme panic dan recover untuk menangani kondisi yang sangat tidak biasa atau kritis. Namun, penggunaan panic dan recover sebaiknya dibatasi karena lebih cocok untuk kasus penanganan kegagalan yang serius.

Error handling adalah bagian penting dari pengembangan perangkat lunak yang membantu Anda memahami dan mengelola situasi yang tidak diharapkan saat program berjalan. Dengan menggunakan tipe error dan pendekatan error handling yang khas Golang, Anda dapat membuat kode yang lebih aman dan dapat diandalkan.
