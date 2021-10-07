package kelengkapan

import (
	"home-loans/user"
)

type GetKelengkapanInput struct {
	UserID int `uri:"user_id" binding:"required"`
}

type CreateKelengkapanInput struct {
	PengajuanID      int    `json:"pengajuan_id"`
	AlamatRumah      string `json:"alamat_rumah" binding:"required"`
	LuasRumah        int    `json:"luas_rumah" binding:"required"`
	HargaRumah       int    `json:"harga_rumah" binding:"required"`
	JangkaPembayaran int    `json:"jangka_pembayaran" binding:"required"`
	Status           string `json:"status"`
	User             user.User
}

type UpdateKelengkapanStatusInput struct {
	Status string `json:"status" binding:"required"`
}
