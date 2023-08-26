package main

// User adalah struktur yang merepresentasikan informasi pengguna.
type User struct {
	ID       int
	Username string
	Password string
}

// UserService adalah layanan untuk mengelola pengguna.
type UserService struct {
	Users []User
}

// GetAllUsers mengembalikan semua pengguna dalam layanan.
func (us *UserService) GetAllUsers() []User {
	return us.Users
}

// GetUserByID mengembalikan pengguna dengan ID yang cocok.
func (us *UserService) GetUserByID(id int) User {
	for _, user := range us.Users {
		if id == user.ID {
			return user
		}
	}
	return User{} // Mengembalikan struct kosong jika tidak ditemukan.
}

/*
Jawaban

1. Setidaknya ada total 6 kekurangan yang ada dalam code program tersebut

2. Berikut beberapa kekurangan yang saya temui pada code program tersebut
- Deklarasi pada struktur yang ada pada variabel `user` tidak mengikuti gaya penamaan konversi dalam go
- Nama-nama field yang ada dalam struktur `user` tidak menggambarkan informasi yang benar
- Metode getallusers() dan getuserbyid() menggunakan nama yang tidak konsisten dan tidak sesuai dengan konvensi Go.
- Variabel t dalam struktur userservice tidak memiliki nama yang jelas dan tidak menggambarkan isi variabel tersebut.
- Metode-metode di dalam userservice seharusnya menggunakan receiver dengan tipe pointer (*UserService) agar perubahan dapat dilakukan di dalam metode tersebut.
- Variabel r pada perulangan tidak memiliki nama yang bermakna.

3. Berikut alasan dari setiap kekurangn tersebut]
- Struktur dan field harus menggunakan gaya penamaan dalam huruf kapital agar bisa diakses dari luar paket. Nama field juga sebaiknya menggambarkan informasi yang jelas.
- Nama metode harus diawali dengan huruf kapital agar dapat diakses dari luar paket. Penggunaan nama yang konsisten membantu memahami fungsionalitas metode.
- Variabel harus memiliki nama yang menjelaskan isinya, sehingga mudah dipahami.
- Receiver dengan tipe pointer diperlukan agar metode dapat memodifikasi data pada receiver.
- Memberikan nama yang bermakna pada variabel membantu membaca kode dengan lebih baik.
*/
