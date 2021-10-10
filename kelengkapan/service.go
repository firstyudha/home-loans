package kelengkapan

import "errors"

type Service interface {
	GetKelengkapans() ([]Kelengkapan, error)
	GetKelengkapan(userID int) ([]Kelengkapan, error)
	CreateKelengkapan(userID int, input CreateKelengkapanInput) (Kelengkapan, error)
	SaveDokumenPendukung(UserID int, fileLocation string) (Kelengkapan, error)
	SaveKelengkapanStatus(userID GetKelengkapanInput, input UpdateKelengkapanStatusInput) (Kelengkapan, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetKelengkapans() ([]Kelengkapan, error) {

	//user doesnt user params
	kelengkapans, err := s.repository.FindAll()
	if err != nil {
		return kelengkapans, err
	}

	return kelengkapans, nil
}

func (s *service) GetKelengkapan(userID int) ([]Kelengkapan, error) {
	pengajuan_id, err := s.repository.FindPengajuanIDByUserID(userID)

	if err != nil {
		return []Kelengkapan{}, err
	}

	kelengkapan, err := s.repository.FindByPengajuanID(pengajuan_id)
	if err != nil {
		return kelengkapan, err
	}

	return kelengkapan, nil

}

func (s *service) CreateKelengkapan(userID int, input CreateKelengkapanInput) (Kelengkapan, error) {
	kelengkapan := Kelengkapan{}
	pengajuan_id, err := s.repository.FindPengajuanIDByUserID(userID)

	if err != nil {
		return kelengkapan, err
	}

	kelengkapan.PengajuanID = pengajuan_id
	kelengkapan.AlamatRumah = input.AlamatRumah
	kelengkapan.LuasRumah = input.LuasRumah
	kelengkapan.HargaRumah = input.HargaRumah
	kelengkapan.JangkaPembayaran = input.JangkaPembayaran
	kelengkapan.Status = input.Status

	isKelengkapanExist, err := s.repository.FindByPengajuanID(pengajuan_id)

	if err != nil {
		return kelengkapan, err
	}

	if len(isKelengkapanExist) > 0 {
		return Kelengkapan{}, errors.New("anda sudah mengajukan Kelengkapan Dokumen")
	}

	newKelengkapan, err := s.repository.Save(kelengkapan)
	if err != nil {
		return newKelengkapan, err
	}

	return newKelengkapan, nil
}

func (s *service) SaveDokumenPendukung(userID int, fileLocation string) (Kelengkapan, error) {
	//find pengajuan id by user id
	pengajuan_id, err := s.repository.FindPengajuanIDByUserID(userID)
	if err != nil {
		return Kelengkapan{}, err
	}

	//find kelengkapan by pengajuan id
	kelengkapan, err := s.repository.FindByID(pengajuan_id)
	if err != nil {
		return Kelengkapan{}, err
	}

	kelengkapan.DokumenPendukung = fileLocation

	updatedKelengkapan, err := s.repository.Update(kelengkapan)
	if err != nil {
		return updatedKelengkapan, err
	}

	return updatedKelengkapan, nil
}

func (s *service) SaveKelengkapanStatus(userID GetKelengkapanInput, input UpdateKelengkapanStatusInput) (Kelengkapan, error) {

	pengajuan_id, err := s.repository.FindPengajuanIDByUserID(userID.UserID)
	if err != nil {
		return Kelengkapan{}, err
	}

	//check if pengajuan kosong
	if pengajuan_id == 0 {
		return Kelengkapan{}, errors.New("data kelengkapan tidak ditemukan")
	}

	kelengkapan, err := s.repository.FindByID(pengajuan_id)
	if err != nil {
		return Kelengkapan{}, err
	}

	kelengkapan.Status = input.Status

	updatedkelengkapan, err := s.repository.Update(kelengkapan)
	if err != nil {
		return updatedkelengkapan, err
	}

	return updatedkelengkapan, nil
}
