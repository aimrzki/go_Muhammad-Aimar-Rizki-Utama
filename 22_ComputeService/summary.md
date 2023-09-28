# Introduction Deployment

## System & Software Deployment
Deployment merupakan kegiatan yang bertujuan untuk menyebarkan aplikasi/produk yang telah dikerjakan oleh para pengembang dan seringkali untuk mengubah dari status sementara ke permanen. Penyebarannya dapat melalui beragam cara tergantung dari jenis aplikasinya, aplikasi web & api ke server sedangkan aplikasi mobile ke Playstore/Appstore.

## Strategi Deployment
* Bing-Bang Deployment Strategy atau sering disebut Replace Deployment Strategy
* Rollout Deployment Strategy
* Blue/Green Deployment Strategy
* Canary Deployment Strategy

### Big-Bang / Replace Deployment Strategy
Kelebihan 
* Mudah di-implementasikan, cara klasik tinggal Replace
* Perubahan kepada sistem langsung 100% secara implementasikan

Kekurangan
* Terlalu beresiko, rata rata downtime cukup lama

### Rollout Deployment Strategy
Kelebihan
* Lebih aman dan less downtime dari versi sebelumnya

Kekurangan 
* Akan ada 2 versi aplikasi berjalan secara barengan sampai semua server terdeploy, dan bisa membuat bingung
* Karena sifatnya perlahan satu persatu, untuk deployment dan rollback lebih lama dari yang Bigbang, karena prosesnya perlahan lahan sampai semua server terkena efeknya
* Tidak ada kontrol request, Server yang baru ter-deploy dengan aplikasi versi bary, langsung mendapatkan request yang sama banyaknya dengan server lain.

### Blue/Green Deoloyment Strategy
Kelebihan :
* Perubahan sangat cepat, sekali switch service langsung berubah 100%
* Tidak ada issue beda versi pada service seperti yang terjadi pada Rollout Deployment

Kekurangan :
* Resource yang dibutuhkan lebih banyak, karena untuk setiap deployment kita harus menyediakan service yang berupa environtmenta dengan yang sedang berjalan di production
* Testing harus benar-benar sangat diprioritaskan sebelum di switch, aplikasi harus kita pastikan aman dari request yang tiba-tiba banyak.

### Canary Deployment Strategy
Kelebihan :
* Cukup aman
* Mudah untuk rollback jika terjadi error/bug, tanpa berimbas kesemua user

Kekurangan :
* Untuk mencapai 100% cukup lama dibanding dengan Blue/Green deployment. Dengan Blue/Green deployment, aplikasi langsung 100% terdeploy keseluruh user.

Pada Deployemnt ini saya menggunakan 2 komponen pada GCP (Google Cloud Platform) yaitu App Engine dan CloudSQL. Mari Kita bahasa satu satu lebih dalam.

## GCP (Google Cloud Platform)
Google Cloud Platform (GCP) adalah salah satu penyedia layanan cloud computing terkemuka di dunia yang disediakan oleh Google. GCP memberikan infrastruktur dan berbagai layanan cloud untuk membantu perusahaan dan individu dalam mengelola aplikasi, data, dan sumber daya komputasi mereka secara efisien. Untuk memberikan penjelasan yang rinci, detail, jelas, dan menarik tentang GCP, mari kita bahas beberapa aspek pentingnya:

1. Infrastruktur Global:
GCP memiliki pusat data di seluruh dunia, yang memungkinkan pelanggan untuk menjalankan aplikasi dan menyimpan data mereka di berbagai lokasi. Ini memungkinkan akses cepat dan andal ke layanan GCP dari hampir mana saja di dunia.

2. Komputasi:
GCP menyediakan berbagai jenis mesin virtual yang dapat diatur sesuai kebutuhan. Anda dapat menjalankan mesin virtual dengan spesifikasi yang berbeda, mulai dari mesin kecil hingga mesin dengan daya komputasi tinggi, sesuai dengan tuntutan aplikasi Anda. Google juga menawarkan layanan seperti Google Kubernetes Engine (GKE) untuk mengelola wadah (container) dan memfasilitasi pengembangan berbasis wadah.

3. Penyimpanan dan Basis Data:
GCP menyediakan layanan penyimpanan termasuk Cloud Storage untuk objek, Cloud SQL untuk basis data SQL, dan Cloud Firestore untuk basis data NoSQL. Ini memungkinkan Anda untuk menyimpan dan mengelola data dengan aman dan skalabilitas yang tinggi.

4. Jaringan dan Keamanan:
GCP memiliki jaringan global yang kuat dan menyediakan alat-alat keamanan canggih untuk melindungi data dan aplikasi Anda. Layanan seperti Virtual Private Cloud (VPC) memungkinkan Anda untuk mengisolasi sumber daya Anda dalam jaringan yang aman.

5. Kecerdasan Buatan (AI) dan Pembelajaran Mesin (ML):
Google Cloud menyediakan alat dan infrastruktur untuk pengembangan dan implementasi model pembelajaran mesin dan kecerdasan buatan. Anda dapat menggunakan TensorFlow (perpustakaan pembelajaran mesin open-source) atau AutoML untuk membuat dan melatih model AI dengan mudah.

6. Layanan Khusus Industri:
GCP juga menawarkan solusi industri khusus, seperti Healthcare API untuk perusahaan kesehatan, Retail Solutions untuk peritel, dan banyak lagi. Ini memungkinkan perusahaan untuk mengintegrasikan solusi cloud yang disesuaikan dengan kebutuhan mereka.

7. Manajemen Layanan:
Google Cloud Console menyediakan antarmuka pengguna yang mudah digunakan untuk mengelola sumber daya dan layanan GCP Anda. Selain itu, GCP juga menawarkan alat berbasis baris perintah (Command Line Interface) untuk otomatisasi dan skrip.

8. Harga yang Kompetitif:
GCP bersaing secara agresif dalam hal harga dengan penyedia cloud lainnya. Mereka menawarkan model penagihan yang fleksibel, termasuk penagihan berdasarkan penggunaan (pay-as-you-go), sehingga Anda hanya membayar untuk sumber daya yang benar-benar Anda gunakan.

9. Ekosistem Partner:
GCP memiliki ekosistem mitra yang besar, termasuk berbagai penyedia layanan manajemen cloud, penyedia perangkat lunak independen, dan konsultan. Ini membantu pelanggan untuk mengintegrasikan dan mengelola solusi cloud mereka dengan lebih baik.

Dengan Google Cloud Platform, perusahaan dan pengembang memiliki akses ke teknologi canggih yang dapat membantu mereka meningkatkan skalabilitas, keamanan, dan kinerja aplikasi mereka, sambil mengurangi biaya infrastruktur dan perawatan fisik. GCP terus berkembang dan menawarkan inovasi terbaru dalam dunia cloud computing, menjadikannya salah satu pemimpin di industri ini.

## App Engine GCP
Google App Engine adalah salah satu layanan komputasi yang disediakan oleh Google Cloud Platform (GCP). Ini adalah platform layanan cloud yang dirancang khusus untuk memungkinkan pengembang untuk dengan mudah membangun, mengelola, dan meng-host aplikasi web dan layanan tanpa harus merisaukan aspek infrastruktur dan administrasi server yang rumit. Untuk memberikan penjelasan yang rinci, detail, jelas, dan menarik tentang Google App Engine (App Engine) di GCP, mari kita bahas beberapa aspek pentingnya:

1. Lingkungan Pemrograman yang Fleksibel:
App Engine mendukung berbagai bahasa pemrograman populer seperti Python, Java, Node.js, Go, dan lainnya. Ini memberikan fleksibilitas bagi pengembang untuk membangun aplikasi menggunakan bahasa yang mereka sukai.

2. Model Penyimpanan Data:
App Engine menyediakan layanan penyimpanan data yang terintegrasi, termasuk Cloud Datastore (basis data NoSQL) dan Cloud SQL (basis data SQL). Ini memungkinkan pengembang untuk menyimpan dan mengelola data aplikasi mereka dengan mudah.

3. Otomatisasi Infrastruktur:
Salah satu keuntungan besar dari App Engine adalah bahwa Google mengelola semua infrastruktur secara otomatis. Anda tidak perlu khawatir tentang mengelola server fisik, pembaruan perangkat lunak, atau skalabilitas. Google akan menangani semuanya sehingga Anda dapat fokus pada pengembangan aplikasi.

4. Skalabilitas Otomatis:
App Engine secara otomatis menangani skalabilitas horizontal. Ini berarti bahwa jika aplikasi Anda mendapatkan lonjakan lalu lintas, App Engine akan secara otomatis menambahkan sumber daya yang diperlukan untuk menangani beban tersebut, dan mengurangi sumber daya saat lalu lintas menurun. Ini membuat aplikasi Anda tetap responsif dan efisien.

5. Manajemen Versi Aplikasi:
Anda dapat dengan mudah mengelola versi yang berbeda dari aplikasi Anda di App Engine. Ini memungkinkan Anda untuk menguji perubahan aplikasi secara aman sebelum menerapkannya ke versi produksi.

6. Pemantauan dan Log:
App Engine menyediakan alat pemantauan dan pelaporan yang kuat, termasuk integrasi dengan Google Cloud Monitoring dan Google Cloud Logging. Anda dapat melacak kinerja aplikasi Anda dan menganalisis log untuk pemecahan masalah.

7. Keamanan:
App Engine memprioritaskan keamanan. Google menyediakan firewall dan mekanisme otentikasi yang kuat, serta pembaruan keamanan otomatis untuk menjaga aplikasi Anda tetap aman.

8. Integrasi dengan Layanan Google Lainnya:
Anda dapat dengan mudah mengintegrasikan App Engine dengan layanan lain yang disediakan oleh GCP, seperti BigQuery, Cloud Storage, Pub/Sub, dan lainnya, untuk memperluas fungsionalitas aplikasi Anda.

9. Model Penagihan yang Elastis:
App Engine menawarkan model penagihan pay-as-you-go, di mana Anda hanya membayar untuk sumber daya yang digunakan. Ini memungkinkan Anda mengontrol biaya dan menghemat uang jika aplikasi Anda memiliki beban lalu lintas yang fluktuatif.

Google App Engine merupakan pilihan yang menarik bagi pengembang yang ingin fokus pada pengembangan aplikasi tanpa harus terlibat dalam perawatan infrastruktur. Dengan fitur-fitur yang kuat dan skalabilitas yang otomatis, App Engine memungkinkan pengembang untuk menciptakan aplikasi web yang dapat berkembang seiring pertumbuhan bisnis mereka.

## Cloud SQL GCP
Google Cloud SQL adalah salah satu layanan yang disediakan oleh Google Cloud Platform (GCP) yang dirancang khusus untuk mengelola basis data relasional. Ini merupakan layanan manajemen basis data yang dikelola sepenuhnya, yang berarti Anda dapat meng-host basis data SQL tanpa harus merisaukan administrasi infrastruktur atau pemeliharaan server. Mari kita bahas Cloud SQL secara rinci, detail, jelas, dan menarik:

1. Dukungan Basis Data SQL Populer:
Cloud SQL mendukung beberapa sistem basis data SQL paling populer, termasuk MySQL, PostgreSQL, dan SQL Server. Ini memungkinkan Anda memilih basis data yang sesuai dengan kebutuhan aplikasi Anda.

2. Otomatisasi Infrastruktur:
Salah satu keunggulan utama dari Cloud SQL adalah bahwa Google mengelola semua aspek infrastruktur secara otomatis. Ini mencakup manajemen server, pembaruan perangkat lunak, pemeliharaan harian, dan pencadangan data. Anda tidak perlu lagi khawatir tentang tugas administrasi ini.

3. Skalabilitas:
Cloud SQL memungkinkan Anda untuk mengubah ukuran basis data Anda dengan mudah untuk mengatasi pertumbuhan bisnis Anda. Anda dapat meningkatkan atau mengurangi kapasitas CPU dan memori basis data dengan beberapa klik.

4. Keamanan:
Google sangat memperhatikan keamanan data di Cloud SQL. Ini mencakup enkripsi data saat istirahat dan penyimpanan, firewall, pengaturan otentikasi yang kuat, serta pemantauan dan pelaporan keamanan.

5. Pemulihan Bencana:
Cloud SQL secara otomatis melakukan pencadangan data secara teratur. Anda dapat dengan mudah mengembalikan basis data ke titik waktu tertentu jika terjadi kegagalan atau bencana.

6. Penyediaan Data Global:
Anda dapat meng-host basis data Cloud SQL di berbagai zona dan wilayah GCP, sehingga Anda dapat memberikan pengalaman aplikasi yang lebih cepat kepada pengguna Anda di seluruh dunia.

7. Integrasi dengan Alat Google Cloud Lainnya:
Cloud SQL dapat dengan mudah diintegrasikan dengan layanan lain yang disediakan oleh GCP, seperti Google Kubernetes Engine (GKE), Google App Engine, dan BigQuery, sehingga Anda dapat membangun solusi end-to-end yang kuat.

8. Manajemen Melalui Antarmuka Web dan Baris Perintah:
Anda dapat mengelola basis data Cloud SQL melalui antarmuka web yang mudah digunakan atau melalui baris perintah dengan menggunakan Google Cloud SDK. Ini memberi Anda fleksibilitas dalam cara Anda berinteraksi dengan basis data Anda.

9. Penagihan Berbasis Penggunaan:
Cloud SQL mengikuti model penagihan pay-as-you-go, yang berarti Anda hanya membayar untuk sumber daya yang Anda gunakan. Ini membuatnya cocok untuk bisnis dengan anggaran yang berfluktuasi.

Google Cloud SQL adalah solusi yang kuat untuk mengelola basis data SQL di lingkungan cloud. Ini menghilangkan banyak beban administrasi yang terkait dengan basis data, memungkinkan pengembang dan perusahaan untuk fokus pada pengembangan aplikasi mereka tanpa harus khawatir tentang infrastruktur basis data yang rumit. Dengan fitur-fitur keamanan dan skalabilitas yang kuat, Cloud SQL membantu menjaga basis data Anda aman dan responsif.

## Link Deployment Restful API untuk tugas ini 
[Alta-Aimar](https://alta-aimar.uc.r.appspot.com/)

