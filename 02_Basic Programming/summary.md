# Basic Programming

Golang
Merupakan Bahasa open source yang dikembangkan oleh google yang bertujuan untuk memudahkan proses development name tetap simple, handle, dan efisiensi software

Golang sangat bogus untuk menulis lower level programing. Seperti membuat program aplikasi dan system program

contoh program aplikasi : 
E-commerce Music Player
Social Media Apps

Contoh System Program : 
APIs
Game Engine
CLI apps

## Kenapa Harus Golang ? 

Ada beberapa alasan mengapa kita harus mempelajari golang seperti : 

- Karena golang merupakan bahasa yang simple, dan membuat mengoding menjadi lebih menyenangkan. 
- Golang memiliki sytaks yang ringan dan memiliki build in concurrency didalamnya.
- Golang merupakan opensource yang banyak digunakan oleh perusahaan besar sebagai bahasa programing mereka

## Variabel dan Tipe Data
Variabel adalah tempat untuk menyimpan data. data biasa memiliki nama dan nama digunakan saat kita ingin memanggil data yang ada di dalamnya.

Tipe data yang ada di bahasa golang :  Boolean : True or False
Numeric : integer, float, dan complex
String

contoh Variabel : var age int

Ada dua cara deklarasi dalam bahasa golang 
* Long : var age int
* short : age:=

Kedua cara di atas valid/ benar tergantung kita lebih nyaman yang mana


Nilai Default

Lalu dalam golang juga ada nilai defuald value untuk variabel yaitu :  booleans -> false
floats -> 0.0
integers-> 0
strings -> “”


const -> merupakan variabel konstan yang tidak dapat diubah, sehingga nilai hanya dimasukan diawal dan tidak bisa diubah

const bisa diinisialisasikan dengan menggunakan 2 cara yaitu :

* single constant : cons pi float = 3.14
* multiple constant : const( 
pi float64 = 3.14 
pi2  
age=10
)

## Operator 
Sama seperti bahasa pemrograman lainnya, ada banyak operator yang dapat kita terapkan seperti :

Aritmatika : 
(+) (Penjumlahan)
(-) (Pengurangan)
(/) (Pembagian)
(*) (Perkalian)
(%) (Sisa hasil bagi)
(++) (increment)
(—) (decrement)

Comparision :
* (==)
* (!=)
* (>)
* (<)
* (>=)
* (<=)

Logical :
* (&&)
* (||)
* (!)

Bitwise :
* (&)
* (|)
* (^)
* (<<)
* (>>)

Assignment :
* (=)
* (+=)
* (-=)
* (*=)
* (/=)
* (%=)
* (<<=)
* (>>=)
* (&=)
* (^=)
* (!=)

Miscellaneous :
* (&)
* (* (Pointer))
 
#Control Structure 
merupakan  konsep dalam pemrograman yang digunakan untuk mengontrol aliran eksekusi program. Mereka memungkinkan Anda untuk membuat keputusan, mengulang kode, dan mengorganisir aliran program sesuai dengan kebutuhan.

ada beberapa metode yang dapat kita gunakan untuk melakukan brancing bisa menggunakan logika (if else) dan logika(switch)

Perulangan :

Perulangan digunakan untuk mengizinkan code untuk diulang sampai dengan kondisinya terpenuhi. Perulangan dapat kita terapkan dengan menggunakan logika (for)

misal logika untuk mencetak angka yang kurang dari 5 sebagai berikut :   for i:=0;i<5;i++{ fmt.println(i) }
