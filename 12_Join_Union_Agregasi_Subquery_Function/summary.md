# Ringkasan Materi Database: Schema & Data Definition Language, dan Perintah SQL

## Schema & Data Definition Language (DDL)
* CREATE TABLE: Digunakan untuk membuat tabel baru dalam basis data. Dalam perintah ini, Anda mendefinisikan kolom-kolom yang akan ada di dalam tabel berserta tipe data dan batasan lainnya.
```
  CREATE TABLE Mahasiswa (
    id INT PRIMARY KEY,
    nama VARCHAR(50),
    umur INT
);

```
* ALTER TABLE: Digunakan untuk mengubah struktur tabel setelah tabel dibuat, seperti menambah atau menghapus kolom.
```
ALTER TABLE Mahasiswa
ADD COLUMN alamat VARCHAR(100);

```
## Manipulasi Data (DML)
* INSERT INTO: Digunakan untuk memasukkan data baru ke dalam tabel.
```
INSERT INTO Mahasiswa (id, nama, umur)
VALUES (1, 'John Doe', 20);

```
* SELECT: Digunakan untuk mengambil data dari tabel. Dapat disesuaikan dengan klausa seperti WHERE, LIKE, BETWEEN, dan AND/OR untuk mengatur kondisi pencarian.
```
SELECT nama, umur FROM Mahasiswa
WHERE umur > 18;

```
* UPDATE: Digunakan untuk memperbarui data yang sudah ada dalam tabel.
```
UPDATE Mahasiswa
SET umur = 21
WHERE nama = 'John Doe';

```
* DELETE: Digunakan untuk menghapus data dari tabel.
```
DELETE FROM Mahasiswa
WHERE id = 1;

```

## Klausa Tambahan
* LIKE / BETWEEN: Digunakan dalam klausa WHERE untuk memfilter data berdasarkan pola atau rentang nilai tertentu.
```
SELECT nama FROM Mahasiswa
WHERE nama LIKE 'J%';
```
* AND / OR: Digunakan untuk menggabungkan kondisi dalam klausa WHERE untuk pencarian yang lebih kompleks.
```
SELECT nama FROM Mahasiswa
WHERE umur > 18 AND (nama LIKE 'J%' OR nama LIKE 'M%');
```
* ORDER BY: Digunakan untuk mengurutkan hasil query berdasarkan kolom tertentu, baik secara naik atau turun.
```
SELECT nama, umur FROM Mahasiswa
ORDER BY umur DESC;
```

* LIMIT: Digunakan untuk membatasi jumlah hasil yang ditampilkan dalam query.
```
SELECT nama FROM Mahasiswa
LIMIT 5;
```

## Penggabungan Tabel (JOIN)
* JOIN: Menggabungkan data dari dua atau lebih tabel berdasarkan kolom yang memiliki nilai sama.
```
SELECT Mahasiswa.nama, MataKuliah.nama AS mata_kuliah
FROM Mahasiswa
JOIN MataKuliah ON Mahasiswa.id = MataKuliah.mahasiswa_id;
```
* INNER JOIN: Mengambil hanya baris yang memiliki nilai yang cocok di kedua tabel yang digabungkan.
```
SELECT Mahasiswa.nama, Nilai.nilai
FROM Mahasiswa
INNER JOIN Nilai ON Mahasiswa.id = Nilai.mahasiswa_id;
```
* LEFT JOIN: Mengambil semua baris dari tabel kiri dan baris yang cocok dari tabel kanan.
```
SELECT Mahasiswa.nama, Nilai.nilai
FROM Mahasiswa
LEFT JOIN Nilai ON Mahasiswa.id = Nilai.mahasiswa_id;
```
* RIGHT JOIN: Kebalikan dari LEFT JOIN, mengambil semua baris dari tabel kanan dan baris yang cocok dari tabel kiri.
```
SELECT Mahasiswa.nama, Nilai.nilai
FROM Mahasiswa
RIGHT JOIN Nilai ON Mahasiswa.id = Nilai.mahasiswa_id;
```
* UNION: Menggabungkan hasil dari dua atau lebih query yang menghasilkan struktur kolom yang sama.
```
SELECT nama FROM Mahasiswa WHERE umur > 20
UNION
SELECT nama FROM Mahasiswa WHERE umur < 18;
```

## Fungsi Agregasi
* Fungsi agregasi seperti COUNT, SUM, AVG, MAX, dan MIN digunakan untuk menghitung atau merangkum data dalam suatu kolom.
```
SELECT COUNT(*) AS total_mahasiswa 
FROM Mahasiswa;
```

## Subquery 
* Subquery adalah query yang ada di dalam query utama. Digunakan untuk mengambil data yang akan digunakan dalam kondisi atau perhitungan lain dalam query utama.
```
SELECT nama FROM Mahasiswa
WHERE umur > (SELECT AVG(umur) FROM Mahasiswa);
```

## Function
* Function: Dalam SQL, ada banyak fungsi bawaan seperti CONCAT, SUBSTRING, DATE_FORMAT, dll., yang dapat digunakan untuk memanipulasi data sebelum ditampilkan atau disimpan.
```
SELECT CONCAT(nama, ' - ', umur) AS info_mahasiswa FROM Mahasiswa;

```
