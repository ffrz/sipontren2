package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"pesantren-monorepo/backend/internal/shared/config"
	"pesantren-monorepo/backend/pkg/model"
)

// InitDB menginisialisasi koneksi database menggunakan konfigurasi
func InitDB(cfg *config.Config) *gorm.DB {
	// PENTING: Menggunakan fmt.Sprintf secara eksplisit untuk menjamin GORM menggunakan DB_NAME
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

	log.Printf("Attempting DB connection with DSN: %s", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Gagal terhubung ke database Postgres. Error: %v", err)
		os.Exit(1)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Printf("Gagal mendapatkan *sql.DB. Error: %v", err)
		os.Exit(1)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("Koneksi Database berhasil.")

	// --- Migrasi & Seed Data ---
	runMigrations(db)

	return db
}

func runMigrations(db *gorm.DB) {
	// Lakukan AutoMigrate untuk semua model
	err := db.AutoMigrate(&model.GlobalSetting{})
	if err != nil {
		log.Fatalf("Gagal melakukan migrasi database: %v", err)
	}

	// Inisialisasi data pengaturan default (Seed Data)
	seedSettings(db)
	log.Println("Migrasi dan Seed Data selesai.")
}

// seedSettings memastikan pengaturan dasar ada di database
func seedSettings(db *gorm.DB) {
	defaultSettings := []model.GlobalSetting{
		{Key: "INSTITUTION_NAME", Value: "Yayasan Pesantren Digital Indonesia"},
	}

	for _, setting := range defaultSettings {
		// Cek apakah key sudah ada
		var existingSetting model.GlobalSetting
		result := db.Where("key = ?", setting.Key).First(&existingSetting)

		if result.Error != nil && result.Error == gorm.ErrRecordNotFound {
			// Jika belum ada, buat record baru
			db.Create(&setting)
		}
	}
}
