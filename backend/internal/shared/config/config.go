package config

import (
	"log"
	"os"
)

// Config struct menampung semua variabel lingkungan aplikasi.
type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	ServerPort string
	JWTSecret  string
}

// GlobalConfig adalah instance konfigurasi yang akan digunakan di seluruh aplikasi.
var GlobalConfig Config

// LoadConfig memuat konfigurasi dari variabel lingkungan.
// Di lingkungan Docker, ini akan mengambil variabel yang disuntikkan dari docker-compose.yml.
func LoadConfig() {
	// Fungsi ini tidak lagi memanggil godotenv.Load() karena di lingkungan Docker,
	// variabel sudah disuntikkan langsung, dan godotenv dapat menyebabkan konflik.

	GlobalConfig = Config{
		DBHost:     getEnv("DB_HOST", "db"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "user_pesantren"),
		DBPassword: getEnv("DB_PASSWORD", "supersecret"),
		DBName:     getEnv("DB_NAME", "pesantren_db"), // PENTING: Memastikan DB_NAME terbaca dengan benar.
		ServerPort: getEnv("SERVER_PORT", "8080"),
		JWTSecret:  getEnv("JWT_SECRET", "super_secure_jwt_key_please_change"),
	}

	log.Println("Konfigurasi berhasil dimuat.")
	log.Printf("DB_HOST: %s, DB_PORT: %s, DB_NAME: %s", GlobalConfig.DBHost, GlobalConfig.DBPort, GlobalConfig.DBName)
}

// getEnv adalah helper untuk membaca variabel lingkungan dengan nilai default.
func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// Init menjalankan LoadConfig
func Init() {
	LoadConfig()
}
