package kelengkapan

type KelengkapanFormatter struct {
	ID               int    `json:"id"`
	PengajuanID      int    `json:"pengajuan_id"`
	AlamatRumah      string `json:"alamat_rumah"`
	LuasRumah        int    `json:"luas_rumah"`
	HargaRumah       int    `json:"harga_rumah"`
	JangkaPembayaran int    `json:"jangka_pembayaran"`
	Status           string `json:"status"`
}

func FormatKelengkapan(kelengkapan Kelengkapan) KelengkapanFormatter {
	kelengkapanFormatter := KelengkapanFormatter{}
	kelengkapanFormatter.ID = kelengkapan.ID
	kelengkapanFormatter.PengajuanID = kelengkapan.PengajuanID
	kelengkapanFormatter.AlamatRumah = kelengkapan.AlamatRumah
	kelengkapanFormatter.LuasRumah = kelengkapan.LuasRumah
	kelengkapanFormatter.HargaRumah = kelengkapan.HargaRumah
	kelengkapanFormatter.JangkaPembayaran = kelengkapan.JangkaPembayaran
	kelengkapanFormatter.Status = kelengkapan.Status

	return kelengkapanFormatter
}

func FormatKelengkapans(kelengkapans []Kelengkapan) []KelengkapanFormatter {
	kelengkapansFormatter := []KelengkapanFormatter{}

	for _, kelengkapan := range kelengkapans {
		kelengkapanFormatter := FormatKelengkapan(kelengkapan)
		kelengkapansFormatter = append(kelengkapansFormatter, kelengkapanFormatter)
	}

	return kelengkapansFormatter
}
