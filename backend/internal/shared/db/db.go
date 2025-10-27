package db

import (
	"fmt"
	"log"
	"os"   // Tambahkan import os
	"time" // Tambahkan import time

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"pesantren-monorepo/backend/internal/shared/config" // Asumsikan config.GlobalConfig ada
)

// InitDB menginisialisasi koneksi database menggunakan konfigurasi
func InitDB(cfg *config.Config) *gorm.DB {
	// PENTING: Menggunakan fmt.Sprintf secara eksplisit untuk menjamin GORM menggunakan DB_NAME
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

	// Log DSN untuk Debugging
	log.Printf("Attempting DB connection with DSN: %s", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		// PENTING: Gunakan log.Printf dan os.Exit(1) untuk lingkungan dev (air restart)
		log.Printf("Gagal terhubung ke database Postgres. Error: %v", err)
		// Keluarkan program untuk memicu restart (Hot Reload)
		os.Exit(1)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Printf("Gagal mendapatkan *sql.DB. Error: %v", err)
		os.Exit(1)
	}

	// Set konfigurasi koneksi
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("Koneksi Database berhasil.")

	// db.AutoMigrate(&model.User{}) // Jika Anda punya model awal

	return db
}
