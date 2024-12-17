# Laporin Backend

**Laporin Backend** adalah backend service yang dikembangkan menggunakan **Golang** untuk mendukung aplikasi **Laporin**, platform yang digunakan untuk membuat pengaduan terkait layanan publik.

---

## 📑 Fitur Utama

- **Autentikasi dan Autorisasi**:  
  Mendukung login, registrasi, dan manajemen peran pengguna.
- **Manajemen Data**:  
  CRUD untuk data layanan dan informasi yang disediakan.
- **API Terstruktur**:  
  Mendukung integrasi dengan frontend melalui REST API.
- **Keamanan**:  
  Menggunakan JWT untuk autentikasi, validasi input, dan sanitasi data.
- **Testing**:  
  Unit testing untuk memastikan stabilitas aplikasi.

---

## 🛠️ Teknologi yang Digunakan

- **Bahasa Pemrograman**: Golang  
- **Framework**: Echo (untuk HTTP server)  
- **Database**: MySQL  
- **ORM**: GORM  
- **Middleware**: JWT, CORS  
- **Deployment**: AWS  
- **Tooling**:  
  - Postman untuk dokumentasi API  

---

## 📂 Struktur Proyek

```plaintext
Laporin-Backend/
├── config/         # Konfigurasi aplikasi (database, JWT, dll.)
├── controllers/    # Logika bisnis dan handler untuk HTTP request
├── entities/       # Model database
├── middlewares/    # File unit testing
├── repositories/   # Repository layer untuk akses data
├── routes/         # Routing untuk endpoint API
├── services/       # Logika layanan yang terpisah dari controller
├── uploads/        # Fungsi penyimpan file csv 
├── utils/          # Fungsi pendukung
└── main.go         # Entry point aplikasi

---

## 🚀 Cara Menjalankan Proyek

### Prasyarat

Sebelum memulai, pastikan Anda sudah menginstal:  

- **Go**: Versi 1.21 atau lebih baru  
- **Database**: MySQL  
- **Git**: Untuk meng-clone repository  

---

### Langkah-Langkah

1. **Clone repository ini**  

   ```bash
   git clone https://github.com/e-complaint-kelompok-8/Backend.git
   cd Backend

---

2. **Konfigurasi file .env: Buat file .env di root project dan tambahkan konfigurasi berikut sebagai contoh:**  

   ```bash
   DATABASE_HOST="wishlistdb.c5c26iyuumlc.ap-southeast-2.rds.amazonaws.com"
   DATABASE_PORT="3306"
   DATABASE_USER="root"
   DATABASE_PASSWORD="Mamatsuramat1518"
   DATABASE_NAME="capstone"
   JWT_SECRET_KEY="ilhan321"
   APP_ENV="development"

   DATABASE_HOST="localhost"
   DATABASE_PORT="3306"
   DATABASE_USER="root"
   DATABASE_PASSWORD="Mamatsutiyem1518"
   DATABASE_NAME="capstone"

   SMTP_PASSWORD="lspq gjjw zuui pkpv"
   SMTP_EMAIL="filipi.ketaren@gmail.com"

---

3. **Jalankan perintah berikut untuk menginstal dependency:**  

   ```bash
   go mod tidy

---

4. **Migrasikan database:**  

   ```bash
   go run main.go migrate

---

5. **Jalankan server:**  

   ```bash
   go run main.go

**Aplikasi akan berjalan pada http://localhost:8080.**

---

## 📜 Dokumentasi API

**Gunakan Postman untuk melihat dokumentasi API. Setelah server berjalan, buka:**

```bash
https://laporin-capstone-project.postman.co/workspace/My-Workspace~f88d9198-7196-4c52-8b2b-80bc41759c23/collection/38993574-7dd5fbee-cbea-40a2-9ade-5362d37769af

---

## 🧪 Testing

**Untuk menjalankan unit test, gunakan perintah berikut:**

```bash
go test ./...

---

*Dikembangkan oleh Kelompok 8, Capstone Project Alterra Batch 7.*