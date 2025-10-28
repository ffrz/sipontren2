package model

import "gorm.io/gorm"

// GlobalSetting merepresentasikan tabel untuk pengaturan konfigurasi global key-value.
type GlobalSetting struct {
	gorm.Model
	Key   string `gorm:"uniqueIndex;not null"`
	Value string `gorm:"type:text"`
}
