package pengajuan

type PengajuanFormatter struct {
	ID                 int                           `json:"id"`
	UserID             int                           `json:"user_id"`
	Nik                string                        `json:"nik"`
	NamaLengkap        string                        `json:"nama_lengkap"`
	TempatLahir        string                        `json:"tempat_lahir"`
	TanggalLahir       string                        `json:"tanggal_lahir"`
	Alamat             string                        `json:"alamat"`
	Pekerjaan          string                        `json:"pekerjaan"`
	PendapatanPerbulan int                           `json:"pendapatan_per_bulan"`
	BuktiKtp           string                        `json:"bukti_ktp"`
	BuktiSlipGaji      string                        `json:"bukti_slip_gaji"`
	Status             string                        `json:"status"`
	Kelengkapan        PengajuanKelengkapanFormatter `json:"kelengkapan"`
	User               PengajuanUserFormatter        `json:"user"`
}

type PengajuanUserFormatter struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	LoginAs  int    `json:"login_as"`
}

type PengajuanKelengkapanFormatter struct {
	ID               int    `json:"id"`
	PengajuanID      int    `json:"pengajuan_id"`
	AlamatRumah      string `json:"alamat_rumah"`
	LuasRumah        int    `json:"luas_rumah"`
	HargaRumah       int    `json:"harga_rumah"`
	JangkaPembayaran int    `json:"jangka_pembayaran"`
	DokumenPendukung string `json:"dokumen_pendukung"`
	Status           string `json:"status"`
}

func FormatPengajuan(pengajuan Pengajuan) PengajuanFormatter {
	pengajuanFormatter := PengajuanFormatter{}
	pengajuanFormatter.ID = pengajuan.ID
	pengajuanFormatter.UserID = pengajuan.UserID
	pengajuanFormatter.Nik = pengajuan.Nik
	pengajuanFormatter.NamaLengkap = pengajuan.NamaLengkap
	pengajuanFormatter.TempatLahir = pengajuan.TempatLahir
	pengajuanFormatter.TanggalLahir = pengajuan.TanggalLahir
	pengajuanFormatter.Alamat = pengajuan.Alamat
	pengajuanFormatter.Pekerjaan = pengajuan.Pekerjaan
	pengajuanFormatter.PendapatanPerbulan = pengajuan.PendapatanPerbulan
	pengajuanFormatter.BuktiKtp = pengajuan.BuktiKtp
	pengajuanFormatter.BuktiSlipGaji = pengajuan.BuktiSlipGaji
	pengajuanFormatter.Status = pengajuan.Status

	user := pengajuan.User

	pengajuanUserFormatter := PengajuanUserFormatter{}
	pengajuanUserFormatter.ID = user.ID
	pengajuanUserFormatter.Username = user.Username
	pengajuanUserFormatter.LoginAs = user.LoginAs

	pengajuanFormatter.User = pengajuanUserFormatter

	kelengkapan := pengajuan.Kelengkapan

	pengajuanKelengkapanFormatter := PengajuanKelengkapanFormatter{}
	pengajuanKelengkapanFormatter.ID = kelengkapan.ID
	pengajuanKelengkapanFormatter.PengajuanID = kelengkapan.PengajuanID
	pengajuanKelengkapanFormatter.AlamatRumah = kelengkapan.AlamatRumah
	pengajuanKelengkapanFormatter.LuasRumah = kelengkapan.LuasRumah
	pengajuanKelengkapanFormatter.HargaRumah = kelengkapan.HargaRumah
	pengajuanKelengkapanFormatter.HargaRumah = kelengkapan.HargaRumah
	pengajuanKelengkapanFormatter.JangkaPembayaran = kelengkapan.JangkaPembayaran
	pengajuanKelengkapanFormatter.DokumenPendukung = kelengkapan.DokumenPendukung
	pengajuanKelengkapanFormatter.Status = kelengkapan.Status

	pengajuanFormatter.Kelengkapan = pengajuanKelengkapanFormatter

	return pengajuanFormatter
}

func FormatPengajuans(pengajuans []Pengajuan) []PengajuanFormatter {
	pengajuansFormatter := []PengajuanFormatter{}

	for _, pengajuan := range pengajuans {
		pengajuanFormatter := FormatPengajuan(pengajuan)
		pengajuansFormatter = append(pengajuansFormatter, pengajuanFormatter)
	}

	return pengajuansFormatter
}
