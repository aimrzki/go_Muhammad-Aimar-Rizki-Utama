# Recursive, Number Theory, Searching, and Sorting

## Recursive
Rekursif adalah keadaan dimana suatu fungsi memecahkan masalah dengan memanggil dirinya berulang kali, jika masalahnya cukup kecil fungsi rekursi menghasilkan jawaban langsung. Jika masalahnya terlalu besar fungsi akan memanggil dirintya sendiri dengan yang lebih kecil ruang lingkup masalahnya

Beberapa alasan recursive sangat dibutuhkan :
* Banyak masalah lebih mudah untuk memecahkan dan mempersingkat kode saat menggunakan pendekatan Rekursif
* Pada dasarnya, strategi iteratif (misalnya dengan loop for), dan rekursif melakukan sesuatu yang berulang. Namun terkadang solusi iteratif untuk suatu masalah sangaty sulit untuk dipikirkan dan memerlukan teknik khusus
* Dengan solusi rekursif, mungkin lebih mudah untuk melihat dan merancang jalur penyelesaian

Beberapa strategi untuk melakukan recursive :
* Kasus dasar : apa kasus yang paling sederhana dari masalah ini
* Hubungan berulang : apa hubungan recursive dari masalah ini dengan masalah serupa yang lebih kecil

Contoh soal :

Dalam matematika, faktorial bilangan bulat posityif n, dilambangkan dengan n!, adalah hasil kali semua bilangan bulat positif yang kurang dari atau sama dengan n :

5! = 5 x 4 x 3 x 2 x 1 = 120

```
func factorial (n int){
    if n ==1 {
        return 1
    }else {
        return n*factorial(n-1)
    }
}
```

## Number Theory 
Number theory merupakan cabang matematika yang mempelajari bilangan bulat. ada banyak topik dalam teori bilangan yaitu bilangan prima, pembagi persekutuan terbesar, kelipatan persekutuan terkecil, faktorial, prima faktor, dan masih banyak lagi

Contoh program bilangan prima :

Versi paling naive adalah menguji menurut definisi yaitu menguji apakah N habis dibagi oleh pembagi E [2..N-1]. ini berfungsi, tetapi berjalan di O(N).

Peningkatan besar pertama adalah menguji apakah N habis dibagi oleh pembagi E [2...sqrt(N)]

```
func checkPrime(number int) bool{
    if number<2{
        return false
    }
    sqrtNumber := int(math.Sqrt(float(number)))
    for i := 2;i<=sqrtNumber;i++{
        if number%i == 0{
            return false
        }
    }
    return true
}
```

Contoh Faktor Persekutuan terbesar :

Pembagian persekutuan terbesar dari bilangan bulat a dan b, dilambangkan gcd(a,b), adalah bilangan bulat terbesar yang membagi a dan b. misalnya, gcd(30,12)=6. Algoritma eucid memberikan cara yang efisien untuk menghitung nilai gcd (a,b).

```
func gcd(a int, b int) int{
    if b==0{
        return a
    }
    return gcd (b, a%b)
}
```


### Searching
Merupakan proses untuk menemukan posisi nilai tertentu dalam daftar nilai

Contoh : 
* Linear Search - O(n)

```
func linearSearch(elements []int,x int) int{
    for i:=0;i<len(elements);i++{
        if elements[i]==x{
            return i
        }
    }
    return -1
}
```

* Builtins Search

```
elements := []int{12,15,15,19,24,31,53,59,60}
value := 31
findIndex := sort.SearchInts(elements, value)
if value == elemnts[findIndex]{
    fmt.println("Value found in elements")
}else{
    fmt.println("Value not found in elements")
}
```

## Sorting
Sorting adalah proses menyusun data dalam urutan tertentu, biasanya kami mengurutkan berdasarkan nilai elemen. Kita dapat mengurutkan angka, kata, pasangan, dan masih banyak lagi. Misalnya kita dapat mengurutkan siswa berdasarkan tinggi badannya, dan kita dapat mengurutkan kota menurut abjad atau menurut jumlahwarganya, urutan yang paling sering digunakan adalah urutan angka dan urutan abjad

Mari pertimbangkan himpunan paling sederhana, sebuah array yang terdiri dari bilangan bulat :

5, 2, 8, 14, 1, 16

misal kita ingin mengurutkannya menjadi urutan berikut :

1, 2, 5, 8, 14, 16

ada banyak algoritma pengurutan, dan sangat berbeda dalam hal kompelksitas t waktu dan poenggunaan memori

Contoh :
* Selection Sort -O (n^2)

```
package main

import "fmt"

func selectionSort(arr []int) {
    n := len(arr)

    for i := 0; i < n-1; i++ {
        minIndex := i
        for j := i + 1; j < n; j++ {
            if arr[j] < arr[minIndex] {
                minIndex = j
            }
        }
        arr[i], arr[minIndex] = arr[minIndex], arr[i] // Swap the minimum element with the current element
    }
}

func main() {
    arr := []int{64, 25, 12, 22, 11}
    fmt.Println("Unsorted array:", arr)

    selectionSort(arr)

    fmt.Println("Sorted array:", arr)
}

```

* Counting Sort -O(n+k)

```
package main

import (
	"fmt"
)

func countingSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	// Temukan nilai maksimum dalam array
	max := arr[0]
	for _, num := range arr {
		if num > max {
			max = num
		}
	}

	// Inisialisasi array hitungan dan array hasil
	countArray := make([]int, max+1)
	sortedArray := make([]int, len(arr))

	// Hitung frekuensi setiap elemen
	for _, num := range arr {
		countArray[num]++
	}

	// Mengakumulasikan frekuensi untuk mendapatkan posisi akhir setiap elemen
	for i := 1; i <= max; i++ {
		countArray[i] += countArray[i-1]
	}

	// Membangun array hasil
	for _, num := range arr {
		sortedArray[countArray[num]-1] = num
		countArray[num]--
	}

	return sortedArray
}

func main() {
	arr := []int{4, 2, 2, 8, 3, 3, 1}
	fmt.Println("Unsorted array:", arr)

	sortedArray := countingSort(arr)

	fmt.Println("Sorted array:", sortedArray)
}

```
