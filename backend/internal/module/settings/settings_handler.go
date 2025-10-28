package settings

import (
	"encoding/json"
	"net/http"

	"gorm.io/gorm"

	"pesantren-monorepo/backend/pkg/model"
)

// SettingsHandler berisi dependensi (seperti DB) untuk modul settings.
type SettingsHandler struct {
	DB *gorm.DB
}

// NewSettingsHandler membuat instance baru dari SettingsHandler.
func NewSettingsHandler(db *gorm.DB) *SettingsHandler {
	return &SettingsHandler{DB: db}
}

// GetInstitutionName menangani permintaan untuk mendapatkan nama institusi dari pengaturan global.
func (h *SettingsHandler) GetInstitutionName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var setting model.GlobalSetting
	// Cari setting dengan Key yang spesifik
	result := h.DB.Where("key = ?", "INSTITUTION_NAME").First(&setting)

	if result.Error != nil {
		// Jika tidak ditemukan (walaupun harusnya sudah diinisialisasi)
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Pengaturan nama institusi tidak ditemukan."})
		return
	}

	// Kirim respons
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"institution_name": setting.Value})
}
