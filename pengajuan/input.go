package pengajuan

import (
	"home-loans/user"
)

type GetPengajuanInput struct {
	UserID int `uri:"user_id" binding:"required"`
}

type CreatePengajuanInput struct {
	Nik                string `json:"nik" binding:"required"`
	UserID             int    `json:"user_id"`
	NamaLengkap        string `json:"nama_lengkap" binding:"required"`
	TempatLahir        string `json:"tempat_lahir" binding:"required"`
	TanggalLahir       string `json:"tanggal_lahir" binding:"required"`
	Alamat             string `json:"alamat" binding:"required"`
	Pekerjaan          string `json:"pekerjaan" binding:"required"`
	PendapatanPerbulan int    `json:"pendapatan_perbulan" binding:"required"`
	Status             string `json:"status"`
	User               user.User
}

type UpdatePengajuanStatusInput struct {
	Status string `json:"status" binding:"required"`
}
