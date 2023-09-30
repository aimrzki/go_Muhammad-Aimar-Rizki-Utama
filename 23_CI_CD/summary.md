# Continuous Inetgration Delivery Deployment (CI/CD)

## Continuous Inetgration
is an automated process. it is done in order to integrate various codes from different potential sources in order to build it or test it


## Cycle dalam CI/CD

1. Continuous Integration (CI):
Development: Tim pengembangan bekerja pada fitur atau perbaikan bug dalam kode.
Version Control: Kode dikelola dalam sistem kontrol versi seperti Git.
Automated Build: Setiap perubahan kode yang di-commit ke repositori Git akan memicu build otomatis untuk menghasilkan binary atau paket perangkat lunak.
Automated Testing: Setelah build selesai, berbagai tes otomatis dijalankan untuk memastikan bahwa perangkat lunak berfungsi seperti yang diharapkan.
Reporting: Hasil dari build dan tes diunggah ke sistem pelaporan untuk diperiksa oleh pengembang.

2. Continuous Delivery (CD):
Staging Environment: Setelah melewati CI, perangkat lunak dapat ditempatkan di lingkungan staging (mirip dengan produksi) untuk pengujian lebih lanjut.
User Acceptance Testing (UAT): Tim QA atau pemangku kepentingan melaksanakan pengujian UAT untuk memastikan perangkat lunak siap untuk produksi.
Manual Approval: Jika semua pengujian berhasil, tim pengembangan atau pemangku kepentingan memberikan persetujuan manual untuk melanjutkan ke produksi.
Deployment to Production: Perangkat lunak diterapkan di lingkungan produksi.

3. Continuous Deployment (CD):
Automated Deployment: Seiring dengan CI/CD, perangkat lunak secara otomatis diterapkan ke lingkungan produksi setelah lulus semua pengujian, tanpa persetujuan manual.
Monitoring and Feedback: Setelah di produksi, perangkat lunak terus dimonitor. Jika terdeteksi masalah, dapat dilakukan rollback atau tindakan perbaikan segera.

4. Continuous Monitoring:
Monitoring: Lingkungan produksi terus dimonitor untuk kinerja, keamanan, dan keandalan.
Feedback Loop: Jika terjadi masalah, informasi kembali ke tim pengembangan untuk perbaikan lebih lanjut.

5. Continuous Feedback and Improvement:
Feedback Gathering: Pemangku kepentingan dan pengguna memberikan umpan balik yang digunakan untuk perbaikan berkelanjutan.
Back to Development: Berdasarkan umpan balik, tim pengembangan membuat perubahan dan memulai siklus CI/CD lagi untuk memperbarui perangkat lunak.
Siklus CI/CD adalah pendekatan berkelanjutan yang memungkinkan pengembang untuk mengirimkan perubahan perangkat lunak secara cepat, aman, dan efisien. Ini menggabungkan otomatisasi, pengujian, pemantauan, dan umpan balik dari pengguna untuk memastikan perangkat lunak tetap berkualitas tinggi dan siap digunakan dalam waktu singkat.

## Tools CI/CD
1. Code & commit
* Git
* IBM Relational
* Github
* Stash
* Visual Studio

2. Build & Config 
* Maven
* Docker
* Gradle
* CHEF

3. Scan & Test
* Gerrit
* Sonar
* Soasta
* Fitnesse
* JUnit

4. Release 
* uDeploy
* Serena
* Release

5. Deploy
* Docker
* VMware
* Google App Engine
* Microsoft .net
* Openstack
* Kubernetes

## Kesimpulan
CI/CD adalah metodologi pengembangan perangkat lunak yang memfokuskan pada otomatisasi, pengujian, dan pengiriman perubahan perangkat lunak secara terus-menerus. Dalam CI, perubahan kode secara rutin diintegrasikan ke dalam repositori bersama, diikuti oleh proses otomatisasi yang mencakup kompilasi, pengujian, dan pelaporan. CD memperluas konsep ini dengan memungkinkan pengiriman perangkat lunak ke lingkungan produksi secara otomatis setelah melalui tahap pengujian, meminimalkan risiko kesalahan manusia.

Manfaat utama dari CI/CD adalah peningkatan kecepatan pengembangan, peningkatan kualitas perangkat lunak, dan kemampuan untuk merespons perubahan pasar lebih cepat. Dengan memungkinkan pengiriman perangkat lunak yang lebih sering dan stabil, CI/CD membantu organisasi untuk menjadi lebih responsif, inovatif, dan efisien dalam pengembangan dan pengiriman perangkat lunak mereka.
