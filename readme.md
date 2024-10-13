🦁 Aplikasi REST API AnekaZoo🦒

Selamat datang di aplikasi REST API AnekaZoo yang ditulis dalam bahasa Go! Aplikasi ini dirancang untuk mengelola data hewan dengan menggunakan RESTful API.


📄 Fitur Utama
- Tambah data hewan baru
- Melihat daftar semua hewan
- Melihat detail hewan berdasarkan ID
- Memperbarui data hewan
- Hapus data hewan dari database


📚 Prasyarat
Sebelum menjalankan aplikasi ini, pastikan Anda telah menginstal:
- [Docker](https://www.docker.com/get-started)
- [Go](https://golang.org/dl/)


🚀 Cara Menjalankan Aplikasi
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


📥 Alamat API
Berikut adalah alamat API yang dapat Anda akses:

GET /zoos: Mendapatkan daftar semua hewan
POST /zoos: Menambahkan hewan baru
GET /zoos/{id}: Mendapatkan detail hewan berdasarkan ID
PUT /zoos/{id}: Memperbarui data hewan
DELETE /zoos/{id}: Menghapus hewan dari database


📝 Informasi Tambahan
- Pastikan semua dependensi telah terinstal sebelum menjalankan aplikasi.
- Aplikasi ini menggunakan GORM untuk ORM dan Gorilla Mux untuk routing.
- Anda dapat menambahkan, mengedit, dan menghapus data hewan melalui API.
- Untuk pengujian, Anda dapat menggunakan Postman atau alat serupa untuk menguji endpoint API.
- Jika Anda mengalami masalah, periksa log di konsol untuk informasi lebih lanjut.


📬 Kontak
Jika Anda memiliki pertanyaan, silakan hubungi saya di:
- **Email:** septianidrp@gmail.com

Terima kasih! 🦓


