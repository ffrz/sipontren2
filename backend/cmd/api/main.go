package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"pesantren-monorepo/backend/internal/shared/config"
	"pesantren-monorepo/backend/internal/shared/db"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// 1. Muat Konfigurasi
	cfg := config.LoadConfig()
	log.Println("Konfigurasi dimuat. Port Server:", cfg.ServerPort)

	// 2. Inisialisasi Database (Untuk implementasi modular selanjutnya)
	// Kita akan mencoba inisialisasi DB, tapi Gagal di sini tidak akan menghentikan
	// server karena kita hanya ingin menguji koneksi HTTP.

	_, err := db.InitDB(cfg)
	if err != nil {
		// Log error, tetapi tidak Fatal, agar server bisa tetap berjalan
		log.Println("PERINGATAN: Gagal terhubung ke database. Koneksi API dasar tetap berjalan. Error:", err)
	} else {
		log.Println("Koneksi Database berhasil.")
	}

	// 3. Setup Router (Go-Chi)
	router := chi.NewRouter()

	// Middlewares
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))

	// 4. Daftarkan Rute Hello World (Endpoint Uji Coba)
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello World! API Pesantren Modular Monolith berjalan dengan sukses."))
	})

	// Placeholder untuk rute modular (kita akan mendaftarkannya di langkah selanjutnya)
	router.Route("/api", func(r chi.Router) {
		r.Get("/status", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Status: OK. API Ready."))
		})
	})

	// 5. Jalankan Server
	serverAddr := fmt.Sprintf(":%s", cfg.ServerPort)
	log.Printf("Server API berjalan di http://localhost%s", serverAddr)
	if err := http.ListenAndServe(serverAddr, router); err != nil {
		log.Fatal("Server berhenti:", err)
	}
}
