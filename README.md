# Cudocomm API

Proyek ini adalah backend API yang dibangun dengan [Golang](https://golang.org/) dan menggunakan framework [Echo](https://echo.labstack.com/). Proyek ini terhubung dengan PostgreSQL dan menggunakan JWT untuk autentikasi.

## 🛠️ Fitur

- Echo Web Framework
- PostgreSQL database
- JWT Authentication
- Modular project structure

---

## 🚀 Setup Environment

Buat file `.env` di root direktori dan isi dengan konfigurasi berikut:

```env
DB_HOST=127.0.0.1
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgresqls
DB_NAME=database_api

APP_PORT=3000
SECRET_KEY=secret-key
JWT_EXPIRED=30 # in minute
