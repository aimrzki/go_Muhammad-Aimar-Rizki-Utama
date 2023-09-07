
# System Design

## Outline
* Digram Design
* Characteristics of Distributed Systems
* Horizontal Scaling vs Vertical Scaling
* Job/Work Queue
* Load Balancing
* Monolithic vs Microservices
* SQL vs NoSQL
* Caching
* Database Indexing
* Database Replication

## Diagram
Diagram adalah representasi simbolis dari informasi dengan menggunakan teknik visualisasi

#### Diagram Design Tools
* Smartdraw
* Lucidchart
* Whimisical
* draw.io 
* Microsoft Visio 

#### A Common used software design 
* Flowchart
* Use Case
* ERD
* HLA 


## Flowchart
* The theory - represent 1 process
* Process, Decision, Terminator

## Use Case Diagram 
* Diagram use case merangkum rincian pengguna sistem Anda (juga dikenal sebagai aktor) dan interaksi mereka dengan sistem

## Entity Relationship Diagram
Teori : Diagram Entity Relationship (ER) adalah jenis diagram alur yang menggambarkan bagaimana entitas seperti orang, objek, atau konsep berhubungan satu sama lain dalam suatu sistem.

## System Design Basic 
Setiap kali kita merancang sistem yang besar, kita perlu mempertimbangkan beberapa hal:
* Apa sajakah potongan arsitektur berbeda yang dapat digunakan?
* Bagaimana bagian-bagian ini bekerja satu sama lain?
* Bagaimana kita dapat memanfaatkan bagian-bagian ini dengan sebaik-baiknya, apa saja trade-off yang tepat?
Membiasakan konsep-konsep ini akan sangat bermanfaat dalam memahami konsep sistem terdistribusi

## Key Characteritic of Distributed Systems
* Scalability
* Reliablity
* Availability
* Efficiency 
* Serviceability or Manageability

### Scalability
Skalabilitas adalah kemampuan suatu sistem, proses, atau jaringan untuk tumbuh dan mengelola peningkatan permintaan. Sistem terdistribusi apa pun yang dapat terus berkembang untuk mendukung peningkatan jumlah pekerjaan dianggap dapat diskalakan. Suatu sistem mungkin harus melakukan penskalaan karena berbagai alasan seperti peningkatan volume data atau peningkatan jumlah pekerjaan, misalnya jumlah transaksi. Sistem yang dapat diskalakan ingin mencapai penskalaan ini tanpa kehilangan performa

### reliability
Menurut definisi, keandalan adalah probabilitas suatu sistem akan gagal dalam periode tertentu. Secara sederhana, sistem terdistribusi dianggap dapat diandalkan jika sistem tersebut tetap memberikan layanannya bahkan ketika satu atau beberapa komponen perangkat lunak atau perangkat kerasnya gagal. Keandalan mewakili salah satu karakteristik utama dari setiap sistem terdistribusi, karena dalam sistem seperti itu, mesin apa pun yang rusak selalu dapat digantikan oleh mesin lain yang sehat, sehingga memastikan penyelesaian tugas yang diminta.

### Availability
Menurut definisi, ketersediaan adalah waktu suatu sistem tetap beroperasi untuk menjalankan fungsi yang diperlukan dalam periode tertentu. Ini adalah ukuran sederhana dari persentase waktu suatu sistem, layanan, atau mesin tetap beroperasi dalam kondisi normal.

### Efficiency
Untuk memahami cara mengukur Efisiensi sistem terdistribusi, mari kita asumsikan kita memiliki operasi yang berjalan secara terdistribusi dan menghasilkan sejumlah item sebagai hasilnya. Dua ukuran standar efisiensinya adalah waktu respons (atau latensi) yang menunjukkan penundaan untuk mendapatkan item pertama dan throughput (atau bandwidth) yang menunjukkan jumlah item yang dikirimkan dalam satuan waktu tertentu (misalnya, satu detik)

### Service or Manageability
Serviceability atau Manageability adalah kesederhanaan dan kecepatan dimana suatu sistem dapat diperbaiki atau dipelihara. jika waktu untuk memperbaiki sistem yang gagal bertambah, maka ketersediaan akan berkurang. Hal-hal yang perlu dipertimbangkan untuk pengelolaan adalah kemudahan mendiagnosis dan memahami masalah ketika terjadi, kemudahan melakukan pembaruan atau modifikasi, dan seberapa sederhana sistem dioperasikan (yaitu, apakah sistem beroperasi secara rutin tanpa kegagalan atau pengecualian)

### Load Balancing
Load Balancer (LB) adalah komponen penting lainnya dari sistem terdistribusi. ini membantu menyebarkan lalu lintas ke sekelompok server untuk meningkatkan daya tanggap dan ketersediaan aplikasi, situs web, atau database. LB juga melacak status semua sumber daya saat mendistribusikan permintaan. Jika server tidak tersedia untuk menerima permintaan baru atau tidak merespons atau memiliki tingkat kesalahan yang tinggi, LB akan berhenti mengirimkan lalu lintas ke server tersebut. Untuk memanfaatkan skalabilitas dan redundansi penuh, kita dapat mencoba menyeimbangkan beban di setiap lapisan sistem. Kita dapat menambahkan LB di tiga tempat:
* Antara pengguna dan server web
* Antara server web dan lapisan platform internal, seperti server aplikasi atau server cache
* Antara lapisan platform internal dan database

## Monolithic and Microservices

### Monolithic
Aplikasi monolitik memiliki basis kode tunggal dengan banyak modul. modul dibagi menjadi fitur bisnis atau fitur teknis. ia memiliki sistem build tunggal yang membangun seluruh aplikasi dan/atau ketergantungan. ia juga memiliki biner tunggal yang dapat dieksekusi atau diterapkan.

### Microservices
Layanan mikro adalah layanan yang dapat diterapkan secara independen dan dimodelkan pada domain bisnis. Mereka berkomunikasi satu sama lain melalui jaringan, dan sebagai pilihan arsitektur menawarkan banyak pilihan untuk memecahkan masalah yang mungkin Anda hadapi. maka arsitektur layanan mikro didasarkan pada beberapa layanan mikro kolaborasi


## SQL vs NoSQL
SQL dan NoSQL (atau database relasional dan database non-relasional)
* Basis data relasional terstruktur dan memiliki skema yang telah ditentukan sebelumnya seperti buku telepon yang menyimpan nomor telepon dan alamat
* Basis data non-relasional tidak terstruktur, dan memiliki skema dinamis seperti folder file yang menyimpan segala sesuatu mulai dari alamat seseorang dan nomor telepon hingga 'suka' di facebook dan preferensi belanja online.


## Caching
cache yang digunakan pada data yang baru-baru ini diminta kemungkinan besar akan diminta lagi. Mereka digunakan di hampir setiap lapisan komputasi. perangkat keras, sistem operasi, browser web, aplikasi web, dan banyak lagi. Cache seperti memori jangka pendek : memiliki jumlah ruang terbatas, namun biasanya lebih cepat dari sumber data aslinya dan berisi item yang paling terakhir diakses. Cache bisa ada di semua level dalam arsitektur, namun sering ditemukan pada level yang mendekati ujung daun di mana cache diterapkan untuk mengembalikan data dengan cepat tanpa membebani level hilir.

## Database Replication

### Redudancy and Replication
Redudansi adalah duplikasi komponen-komponen penting fungsi suatu sistem dengan tujuan meningkatkan keandalan sistem, biasanya dalam bentuk cadangan fail-safe, atau untuk meningkatkan kinerja sistem sebenarnya. Misalnya, jika hanya ada satu salinan file yang disimpan di satu server, maka kehilangan server tersebut berarti kehilangan file tersebut. Karena kehilangan data jarang merupakan hal yang baik, kita dapat membuat duplikat atau salinan berlebihan dari file tersebut untuk mengatasi masalah ini.
## Database Indexing

### Indexing
Pengindeksan adalah cara untuk mengoptimalkan kinerja database dengan meminimalkan jumlah akses disk yang diperlukan saat kueri diproses. Ini adalah teknik struktur data yang digunakan untuk menemukan dan mengakses data dalam database dengan cepat
