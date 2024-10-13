## 🦁 Aplikasi REST API AnekaZoo🦒

Selamat datang di aplikasi REST API AnekaZoo yang ditulis dalam bahasa Go! Aplikasi ini dirancang untuk mengelola data hewan dengan menggunakan RESTful API.


## 📄 Postman Documentation
Anda dapat melihat dokumentasi API melalui Postman pada link berikut:

[Postman API Documentation](https://documenter.getpostman.com/view/37642908/2sAXxS9Bor)


## 📄 Fitur Utama
- Mengambil Daftar Semua Hewan
- Menambahkan Data Hewan Baru
- Mengambil Detail Hewan Berdasarkan ID
- Memperbarui Data Hewan
- Menghapus Data Hewan dari Database


## 📚 Prasyarat
Sebelum menjalankan aplikasi ini, pastikan Anda telah menginstal:
- [Docker](https://www.docker.com/get-started)
- [Go](https://golang.org/dl/)


## 📚 Daftar Library

Pastikan sudah mengunduh semua library yang diperlukan. Jalankan perintah berikut untuk mengunduh semua library Go yang dibutuhkan:

```bash
go mod tidy
```

Berikut adalah beberapa library penting yang harus di-install:
- github.com/gorilla/mux
- github.com/jinzhu/gorm
- github.com/go-sql-driver/mysql
- log
- net/http


## 🚀 Cara Menjalankan Aplikasi
1. Menjalankan Sistem Penyimpanan
Untuk menjalankan sistem penyimpanan menggunakan Docker, gunakan perintah berikut:

```bash
docker-compose -f docker-storage.yaml up -d
```
Perintah ini akan memulai kontainer Docker yang diperlukan untuk menyimpan data aplikasi.

2. Setelah sistem penyimpanan berjalan, Anda dapat menjalankan aplikasi dengan perintah berikut:
```bash
go run main.go
```
Aplikasi ini akan berjalan di localhost:8080.  


## 📥 Alamat API
Berikut adalah alamat API yang dapat Anda akses:

- **GET /zoos** : Mengambil Daftar Semua Hewan
- **POST /zoos** :Menambahkan Data Hewan Baru
- **GET /zoos/{id}** :Mengambil Detail Hewan Berdasarkan ID
- **PUT /zoos/{id}** :Memperbarui Data Hewan
- **DELETE /zoos/{id}** :Menghapus Data Hewan dari Database


## 📝 Informasi Tambahan
- Pastikan semua dependensi telah terinstal sebelum menjalankan aplikasi.
- Aplikasi ini menggunakan GORM untuk ORM dan Gorilla Mux untuk routing.
- Anda dapat menambahkan, mengedit, dan menghapus data hewan melalui API.
- Untuk pengujian, Anda dapat menggunakan Postman atau alat serupa untuk menguji endpoint API.
- Jika Anda mengalami masalah, periksa log di konsol untuk informasi lebih lanjut.


## 📬 Kontak
Jika Anda memiliki pertanyaan, silakan hubungi saya di:
- **Email:** septianidrp@gmail.com

Terima kasih! 🦓


