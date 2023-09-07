package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// URL endpoint API untuk mengambil data postingan
	url := "https://jsonplaceholder.typicode.com/posts"

	// Mengirim permintaan GET ke server
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Gagal mengirim permintaan GET:", err)
		return
	}
	defer response.Body.Close()

	// Mengecek apakah permintaan berhasil atau tidak
	if response.StatusCode == http.StatusOK { // Kode status 200 berarti berhasil (OK)
		// Membaca data respons dari server
		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Gagal membaca data respons:", err)
			return
		}

		// Menampilkan data postingan yang telah disimpan
		fmt.Println("Data Postingan:")
		fmt.Println(string(responseData))
	} else {
		fmt.Println("Gagal mengambil data postingan. Kode status:", response.Status)
	}
}
