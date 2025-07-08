# Backend Project Setup

## 1. Clone Repository

Clone project ini ke dalam environment lokal Anda:

```bash
git clone <repository-url>
cd <project-folder>
```

## 2. Konfigurasi Environment

Salin file contoh konfigurasi environment:

```bash
cp .env.example .env
```

Lalu, perbarui isi file `.env` sesuai dengan konfigurasi database yang Anda gunakan (host, port, user, password, database name, dll).

## 3. PostgreSQL Setup

Untuk menghindari error berikut saat menjalankan aplikasi:

```
ERROR: function uuid_generate_v4() does not exist (SQLSTATE 42883)
```

Jalankan perintah SQL berikut di database PostgreSQL Anda:

```sql
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
```

Pastikan Anda menjalankan perintah ini **dalam konteks database yang akan digunakan oleh aplikasi**.

## 4. Menjalankan Aplikasi

Untuk menjalankan aplikasi secara langsung:

```bash
go run .
```

Atau, jika Anda menggunakan [Air](https://github.com/cosmtrek/air) untuk hot-reloading saat development:

```bash
air
```
