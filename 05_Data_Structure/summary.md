# Time Complexity & Space Complexity
Time Complexity : merupakan cara untuk kita dapat memperkirakan waktu yang berjalan terhadap suatu program. Kompleksitas dapat dilihat sebagai jumlah maksimum operasi primitif yang dapat dijalankan oleh suatu program. Operasi reguler adalah penjumlahan tunggal, perkalian, penugasan, dan masih banyak lagi. Lalu hal yang paling sering muncul dalam operasi tersebut biasa disebut sebagai dominan 

Contoh Operasi Dominan

func dominant (n int) int {
* var result int = 0
* for i:=0 ; i<n; i++{
* result +=1
* }
* result = result +10
* return result
*}

Pada kode diatas operasi pada baris ke-4 adalah dominan dan akan dieksekusi sebanyak n kali. Kompleksitas dijelaskan dalam notasi Big-O dalam hal ini O(n) - Kompleksitas Linear


Berikut contoh beberapa Time Compexities yang berbeda :
* Constant Time - O(1)
* Linear Time - O(n)
* Linear Time - O(n+m)
* Logarithmic Time - O(log n)
* Quadratic Time - O(n^2)

## Batas Waktu 
Saat ini, rata rata komputer dapat melakukan 10^8 operasi dalam waktu kurang dari satu detik. Batas waktu yang ditetapkan untuk tes online biasanya dari 1 hingga 10 detik
* n<=1.000.000, Kompleksitas waktu yang diharapkan adalah O(n) atau 0(n log n)
* n<=10.000, kompleksitas waktu yang diharapkan adalah 0(n^2)
* n<=500, kompleksitas waktu yang diharapkan adalah 0(n^3)

Batasan ini belum tentu tepat, ini hanya perkiraan dan akan bervariasi tergantung pada tugas tertentu

## Space Complexity
Memori limits memberikan informasi tentang kompleksitas ruang yang diharapkan. Kita dapat memperkirakan jumlah variabel yang dapat kita deklarasikan dalam program. Jika kita memiliki jumlah variabel yang konstan, kita juga memiliki kompleksitas ruang yang konstan : dalam notasi Big-O  ini adalah O(1). Jika kita perlu mendeklarasikan array dengan n elemen, maka kita memiliki kompleksitas ruang linear -O(n)
