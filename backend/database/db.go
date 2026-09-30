package database

import (
	"fmt"
	"log"
	"os"

	"backend/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// Muat file .env jika ada (opsional jika dijalankan via docker/os env)
	_ = godotenv.Load()

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Susun DSN secara dinamis dari file .env
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", 
		dbUser, dbPass, dbHost, dbPort, dbName,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal koneksi ke database MySQL Laragon:", err)
	}

	// Migrasi otomatis tabel ke MySQL
	DB.AutoMigrate(&models.User{})
	log.Println("Berhasil terhubung ke database MySQL Laragon melalui .env!")
}