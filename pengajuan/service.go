package pengajuan

import "errors"

type Service interface {
	GetPengajuans() ([]Pengajuan, error)
	GetPengajuan(userID int) ([]Pengajuan, error)
	CreatePengajuan(input CreatePengajuanInput) (Pengajuan, error)
	SaveBuktiKTP(userID int, fileLocation string) (Pengajuan, error)
	SaveBuktiSlipGaji(userID int, fileLocation string) (Pengajuan, error)
	SavePengajuanStatus(userID GetPengajuanInput, input UpdatePengajuanStatusInput) (Pengajuan, error)
	CheckRecommendation(userID int) (string, error)
	DeletePengajuan(userID int) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetPengajuans() ([]Pengajuan, error) {
	pengajuans, err := s.repository.FindAll()
	if err != nil {
		return pengajuans, err
	}

	return pengajuans, nil
}

func (s *service) GetPengajuan(userID int) ([]Pengajuan, error) {
	pengajuans, err := s.repository.FindByUserID(userID)
	if err != nil {
		return pengajuans, err
	}

	return pengajuans, nil
}

func (s *service) CreatePengajuan(input CreatePengajuanInput) (Pengajuan, error) {
	pengajuan := Pengajuan{}
	pengajuan.UserID = input.UserID
	pengajuan.Nik = input.Nik
	pengajuan.NamaLengkap = input.NamaLengkap
	pengajuan.TempatLahir = input.TempatLahir
	pengajuan.TanggalLahir = input.TanggalLahir
	pengajuan.Alamat = input.Alamat
	pengajuan.Pekerjaan = input.Pekerjaan
	pengajuan.PendapatanPerbulan = input.PendapatanPerbulan
	pengajuan.Status = input.Status

	isPengajuanExist, err := s.repository.FindByID(input.UserID)

	if err != nil {
		return pengajuan, err
	}

	if isPengajuanExist.ID > 0 {
		return Pengajuan{}, errors.New("anda sudah mengajukan KPR")
	}

	newPengajuan, err := s.repository.Save(pengajuan)
	if err != nil {
		return newPengajuan, err
	}

	return newPengajuan, nil
}

func (s *service) SaveBuktiKTP(userID int, fileLocation string) (Pengajuan, error) {
	pengajuan, err := s.repository.FindByID(userID)
	if err != nil {
		return Pengajuan{}, err
	}

	pengajuan.BuktiKtp = fileLocation

	updatedPengajuan, err := s.repository.Update(pengajuan)
	if err != nil {
		return updatedPengajuan, err
	}

	return updatedPengajuan, nil
}

func (s *service) SaveBuktiSlipGaji(userID int, fileLocation string) (Pengajuan, error) {
	pengajuan, err := s.repository.FindByID(userID)
	if err != nil {
		return Pengajuan{}, err
	}

	pengajuan.BuktiSlipGaji = fileLocation

	updatedPengajuan, err := s.repository.Update(pengajuan)
	if err != nil {
		return updatedPengajuan, err
	}

	return updatedPengajuan, nil
}

func (s *service) SavePengajuanStatus(userID GetPengajuanInput, input UpdatePengajuanStatusInput) (Pengajuan, error) {

	pengajuan, err := s.repository.FindByID(userID.UserID)
	if err != nil {
		return Pengajuan{}, err
	}

	pengajuan.Status = input.Status

	updatedPengajuan, err := s.repository.Update(pengajuan)
	if err != nil {
		return updatedPengajuan, err
	}

	return updatedPengajuan, nil
}

func (s *service) CheckRecommendation(userID int) (string, error) {

	pengajuans, err := s.repository.FindByUserID(userID)
	if err != nil {
		return "", err
	}

	if len(pengajuans) == 0 {
		return "", errors.New("data tidak ditemukan")
	}

	pengajuan := pengajuans[0]

	if pengajuan.Kelengkapan.ID == 0 {
		return "", errors.New("customer belum melengkapi data pengajuan")
	}

	kemampuan_cicilan_perbulan := pengajuan.PendapatanPerbulan / 3
	kenyataan_cicilan_perbulan := (pengajuan.Kelengkapan.HargaRumah / pengajuan.Kelengkapan.JangkaPembayaran) / 12

	if kemampuan_cicilan_perbulan > kenyataan_cicilan_perbulan {
		return "Diperbolehkan", nil
	} else {
		return "", errors.New("tidak diperbolehkan, karena kemampuan cicilan perbulan kurang dari kenyataan cicilan perbulan")
	}

}

func (s *service) DeletePengajuan(userID int) error {
	//find pengajuan id by user id
	pengajuan, err := s.repository.FindByID(userID)
	if err != nil {
		return err
	}

	//check if pengajuan kosong
	if pengajuan.ID == 0 {
		return errors.New("data pengajuan tidak ditemukan")
	}

	//find kelengkapan by pengajuan id
	err = s.repository.Delete(pengajuan.ID)
	if err != nil {
		return err
	}

	return nil
}
