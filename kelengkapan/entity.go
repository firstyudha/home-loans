package kelengkapan

import (
	"home-loans/user"

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

type Pengajuan struct {
	ID                 int
	UserID             int
	Nik                string
	NamaLengkap        string
	TempatLahir        string
	TanggalLahir       string
	Alamat             string
	Pekerjaan          string
	PendapatanPerbulan int
	BuktiKtp           string
	BuktiSlipGaji      string
	Status             string
	User               user.User
}
