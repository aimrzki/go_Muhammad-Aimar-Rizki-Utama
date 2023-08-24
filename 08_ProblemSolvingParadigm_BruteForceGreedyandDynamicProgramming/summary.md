# Problem Solving Paradigm - Brute Force Greedy and Dynamic Programming

Outline : 
* What is Problem Solving Paradigm
* Brute Force
* Divide & Conquer
* Greedy 
* Dynamic Programming

## Problem Solving Paradigm
Problem solving paradigms adalah pendekatan yang biasa digunakan untuk memecahkan masalah dalam:
* Complete Search (a.k.a Brute Force)
* Divide and Conquer
* The Greedy approach
* Dynamic Programming

Setiap Masalah perlu kita selesaikan dengan pendekatan yang sesuai

## Complete Search
Complete Search (juga dikenal sebagai) Bruteforce adalah metode untuk memecahkan masalah dengan melintasi seluruh ruang pencarian untuk mendapatkan solusi yang diperlukan. Bruteforce terjadi ketika tidak ada algoritma lain yang tersedia. Biasanya mudah ditulis karena lugas. Secara teoritis semua masalah dapat diselesaikan dengan menggunakan pendekatan BruteForce terutama ketika kita memiliki waktu tidak terbatas

Contoh : 

Find Max and Min

Problem statment :
Kita diberikan array A yang mengandung n <= 10.000

```
findMaxMin ([10,7,3,5,8,2,9]) // 10 2
```

## Divide and Conquer
Divide & Conquer (D&C) adalah paradigma pemecahan masalah di mana suatu masalah dibuat lebih sederhana dengan 'membagi' menjadi bagian-bagian yang lebih kecil dan kemudian menaklukkan setiap bagian. Berikut langkahnya :
* Devide : membagi masalah yang besar menjadi masalah yang lebih kecil
* Conquer : ketika masalah sudah cukup kecil untuk diselesaikan langsung selesaikan
* Combine : jika dibutuhkan maka perlu menggabungkan solusi dari masalah-masalah yang lebih kecil menjadi solusi untuk masalah yang besar

Binary Search

Problem Steatment 
Diberikan sorted array A, temukan apakah ada bilangan bulat D pada array itu dan kembalikan indeks dari nilai itu

Example :

```
binarySearch ([1,1,3,5,5,6,7],3) // 2
binarySearch ([12,15,15,19,24,31,53,59,60],53) // 6
binarySearch ([12,15,15,19,24,31,53,59,60],100) // -1
```

Algortima Binary Search :
```
package main

import "fmt"

func binarySearch(arr []int, target int) int {
    left, right := 0, len(arr)-1

    for left <= right {
        mid := left + (right-left)/2

        if arr[mid] == target {
            return mid
        } else if arr[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }

    return -1 // Return -1 if the target is not found
}

func main() {
    sortedArray := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
    target := 12

    result := binarySearch(sortedArray, target)

    if result != -1 {
        fmt.Printf("Target %d found at index %d\n", target, result)
    } else {
        fmt.Println("Target not found in the array")
    }
}

```

## Greedy
sebuah algoritma dikatakan serakah jika membuat pilihan optimal lokal pada setiap langkah dengan harapan pada akhirnya mencapai solusi optimal global. Dalam beberapa kasus, Greedy berhasil - solusinya singkat dan berjalan efisien

Coin Change 

Problem Statment 
Mengingat jumlah target V sen dan daftar denominasi n koin, yaitu kita memiliki coinValue[i] (dalam sen) untuk jenis koin i E [0..n-1]. berapa jumlah minimum koin yang harus kita gunakan untuk mewakili jumlah V ? Asumsikan bahwa kita memiliki persediaan koin jenis apa pun yang tidak terbatas. Nilai koin = [10,25,5,1]

Exampe 
```
coinChanges(42) // 25 10 5 1 1
```

## Dynamic Programming
Dynamic Programming (DP) adalah sebuah metode dalam ilmu komputer dan matematika yang digunakan untuk memecahkan masalah optimasi dengan cara membagi masalah yang kompleks menjadi submasalah yang lebih kecil dan menyimpan solusi dari submasalah-submasalah tersebut untuk menghindari perhitungan berulang yang tidak efisien. DP sangat berguna ketika ada tumpang tindih dalam submasalah yang perlu dipecahkan.

Dalam bahasa pemrograman Go (Golang), Anda dapat menerapkan Dynamic Programming untuk menyelesaikan berbagai masalah yang melibatkan optimasi, seperti mencari nilai maksimal, minimal, atau menghitung berbagai kombinasi.

Berikut ini adalah panduan langkah-demi-langkah dalam menerapkan Dynamic Programming dalam Go:

* Pahami Masalah: Pertama, Anda harus memahami masalah yang ingin Anda selesaikan dan tentukan apa yang ingin Anda optimalkan. Misalnya, apakah Anda mencari solusi terbaik dengan biaya minimum, atau mencari kombinasi maksimal, atau hal lainnya.

* Identifikasi Submasalah: Pisahkan masalah utama menjadi serangkaian submasalah yang lebih kecil. Submasalah-submasalah ini harus memiliki karakteristik tumpang tindih sehingga solusi dari submasalah yang sama tidak perlu dihitung berkali-kali.

* Definisikan Fungsi DP: Tentukan fungsi DP yang akan menghitung solusi optimal dari setiap submasalah. Fungsi ini harus memanfaatkan hasil dari submasalah yang lebih kecil yang telah dipecahkan sebelumnya. Pada tahap ini, Anda perlu mendefinisikan struktur data yang akan digunakan untuk menyimpan hasil DP, seperti array atau map.

* Buat Solusi Berbasis Iterasi: Dalam implementasi DP, Anda akan sering menggunakan pendekatan berbasis iterasi (bottom-up) daripada rekursi (top-down) karena lebih efisien. Anda akan mulai menghitung solusi dari submasalah yang paling kecil dan secara bertahap membangun solusi untuk masalah yang lebih besar.

* Inisialisasi Nilai Awal: Tentukan nilai awal untuk submasalah yang paling sederhana. Ini sering kali merupakan bagian penting dari implementasi DP.

* Iterasi untuk Menghitung Solusi: Gunakan loop untuk mengisi nilai solusi untuk setiap submasalah. Dalam setiap iterasi, Anda akan menggunakan hasil submasalah yang lebih kecil yang telah dihitung sebelumnya.

* Kembalikan Solusi Akhir: Setelah Anda menghitung solusi untuk masalah utama, kembalikan nilai solusi akhir yang diinginkan.]

Contoh penerapan kode untuk mencari bilangan fibonanci :

```
package main

import "fmt"

func fibonacci(n int) int {
    if n <= 1 {
        return n
    }

    dp := make([]int, n+1)
    dp[0], dp[1] = 0, 1

    for i := 2; i <= n; i++ {
        dp[i] = dp[i-1] + dp[i-2]
    }

    return dp[n]
}

func main() {
    n := 6
    result := fibonacci(n)
    fmt.Printf("Fibonacci(%d) = %d\n", n, result)
}

```

Ingatlah bahwa konsep DP bisa lebih kompleks tergantung pada masalah yang kita hadapi. Jika kita ingin menerapkan DP pada masalah yang lebih rumit, perlu waktu untuk memahami karakteristik submasalah dan menemukan pola tumpang tindih yang dapat membantu kita menghemat waktu perhitungan.

```
hasil := tambah(5, 3) // hasil = 8
```
