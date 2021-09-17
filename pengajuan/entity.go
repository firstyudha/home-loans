package pengajuan

import (
	"home-loans/kelengkapan"
	"home-loans/user"
	"time"

	"gorm.io/gorm"
)

type Pengajuan struct {
	gorm.Model
	ID                 int
	UserID             int
	Nik                string
	NamaLengkap        string
	TempatLahir        string
	TanggalLahir       time.Time
	Alamat             string
	Pekerjaan          string
	PendapatanPerbulan int
	BuktiKtp           string
	BuktiSlipGaji      string
	Status             string
	Kelengkapan        kelengkapan.Kelengkapan
	User               user.User
}
