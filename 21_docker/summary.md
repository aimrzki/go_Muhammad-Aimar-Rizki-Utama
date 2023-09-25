# Docker

Docker adalah sebuah platform yang digunakan untuk mengembangkan, menguji, dan menjalankan aplikasi dalam lingkungan yang disebut sebagai "kontainer." Kontainer adalah unit pemaketan yang berisi semua yang dibutuhkan untuk menjalankan sebuah aplikasi, termasuk kode, runtime, pustaka, dan dependensi lainnya. Docker memberikan cara yang konsisten dan portabel untuk mengemas, mendistribusikan, dan menjalankan aplikasi di berbagai lingkungan komputasi, seperti server fisik, mesin virtual, atau bahkan di lingkungan komputasi awan.

## Mengapa Docker ?
* Kontainerisasi: Docker memungkinkan Anda untuk mengemas aplikasi Anda beserta semua dependensinya ke dalam kontainer yang terisolasi. Ini memungkinkan aplikasi untuk berjalan dengan konsisten di berbagai lingkungan, menghindari konflik dan masalah yang sering terjadi dengan pemasangan manual.
* Ringan: Kontainer Docker jauh lebih ringan dibandingkan mesin virtual (VM). Mereka berbagi kernel sistem operasi tuan rumah, yang membuatnya hemat sumber daya dan memungkinkan Anda menjalankan banyak kontainer pada satu host.
* Portabilitas: Docker membuat aplikasi dapat berjalan di mana saja yang mendukung Docker, baik itu di laptop pengembang, server lokal, atau di lingkungan produksi. Ini menghilangkan masalah yang sering muncul saat menggeser aplikasi antara lingkungan yang berbeda.
* Skalabilitas: Anda dapat dengan mudah menambahkan atau mengurangi instance kontainer untuk menangani lonjakan lalu lintas atau beban kerja yang berubah. Docker juga mendukung orkestrasi kontainer, seperti Kubernetes, untuk mengelola kontainer dalam skala besar.
* Mudah Digunakan: Docker menyediakan perintah sederhana dan antarmuka pengguna yang mudah digunakan untuk mengelola kontainer. Ini membuatnya sangat berguna bagi pengembang, administrator sistem, dan tim operasi.
* Ekosistem Besar: Docker memiliki ekosistem yang luas dengan ribuan gambar kontainer yang tersedia di Docker Hub. Ini memungkinkan Anda untuk dengan mudah menemukan dan menggunakan aplikasi dan layanan yang telah di-dockerize.
* Open Source: Docker adalah perangkat lunak open source, yang berarti Anda dapat menggunakannya secara gratis dan mengkustomisasinya sesuai kebutuhan Anda.
* Kemudahan Integrasi: Docker dapat diintegrasikan dengan berbagai alat dan layanan seperti Jenkins, Git, dan banyak lagi, membuatnya sangat cocok untuk alur kerja pengembangan berkelanjutan (CI/CD).

Dalam ringkasan, Docker adalah alat yang sangat berguna untuk mengemas, mendistribusikan, dan menjalankan aplikasi dengan cepat dan konsisten di berbagai lingkungan. Ini menghemat waktu dan sumber daya, dan merupakan elemen kunci dalam pengembangan perangkat lunak modern dan praktik DevOps.

## Perbedaan Container dan Virtual Machines 
Container :
* Abstraksi pada lapisan aplikasi
* Kontainer memakan lebih sedikit ruang dibandingkan VM
* Menangani lebih banyak aplikasi dan memerlukan lebih sedikit VM dan Sistem Operasi

Virtual Machines :
* Abstraksi perangkat keras fisik
* Setiap VM menyertakan salinan lengkap Sistem operasi
* Juga lambat untuk boot

## Syntax Docker
* FROM -> Mendapatkan gambar dari registri buruh pelabuhan
* RUN -> Jalankan perintah bash saat membuat Kontainer
* ENV -> Tetapkan variabel di dalam Kontainer
* ADD -> Salin file dengan beberapa proses lainnya
* COPY -> Salin file
* WORKDIR -> Atur direktori file kerja
* ENTRYPOINT -> Jalankan perintah setelah selesai membangun Kontainer
* CMD -> Jalankan perintah tetapi dapat ditimpa

## Contoh pembuatan Container
```
FROM golang:1.21.0-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o main.app .

EXPOSE 8000

CMD ["/app/main.app"]
```

Membuat container :
```
docker build -t flask-tutorial:latest .
```

Login ke container registry kita :
```
docker login --username=yourusername
```

Tag dan push image container kita :
```
docker tag 557bec4698b6
```

```
docker push aimar/flask-tutorial
```

Pull dan run image container kita :

```
docker pull aimar/flask-tutorial:1.0
```

```
docker run -d -p 5000:5000 --name flaskdemo
```
