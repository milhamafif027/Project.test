package database

import (
	"log"
	"backend/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	// Format DSN MySQL Laragon: "user:password@tcp(127.0.0.1:3306)/nama_database?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:@tcp(127.0.0.1:3306)/db_project_test?charset=utf8mb4&parseTime=True&loc=Local"
	
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal koneksi ke database MySQL Laragon:", err)
	}

	// Migrasi otomatis tabel ke MySQL
	DB.AutoMigrate(&models.User{})
	log.Println("Berhasil terhubung ke database MySQL Laragon!")
}