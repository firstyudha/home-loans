package kelengkapan

import (
	"gorm.io/gorm"
)

type Kelengkapan struct {
	gorm.Model
	ID               int
	PengajuanID      int
	AlamatRumah      string
	LuasRumah        int
	HargaRumah       int
	JangkaPembayaran int
	DokumenPendukung string
	Status           string
}
